package appdef

import (
	"errors"
	"io/ioutil"

	"path/filepath"

	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

var (
	CONFIG_FILE_NAME     = "fleetcmd.yaml"
	ERROR_READ_FILE      = errors.New("Error reading file")
	ERROR_FILE_STRUCTURE = errors.New("File has wrong structure")
	ERROR_FILE_NOT_FOUND = fmt.Errorf("Cannot find %s within current and parent directories", CONFIG_FILE_NAME)
)

type ClusterConfigSource interface {
	Config() *ClusterConfig
	Refresh() (bool, error)
}

type clusterConfigSourceFile struct {
	configPath    string
	baseDirectory string
	config        *ClusterConfig
}

// FindConfigSource looks current directories and directories above for fleetcmd.yaml file
func FindConfigSource(basePath string) (ClusterConfigSource, error) {
	path := basePath
	_, err := os.Stat(filepath.Join(path, CONFIG_FILE_NAME))
	for nil != err && len(path) > 1 {
		path = filepath.Dir(path)
		_, err = os.Stat(filepath.Join(path, CONFIG_FILE_NAME))
	}

	if err != nil {
		return nil, err
	}

	if len(path) > 1 {
		return NewClusterConfigSourceFile(path), nil
	}

	return nil, ERROR_FILE_NOT_FOUND
}

func NewClusterConfigSourceFile(path string) ClusterConfigSource {
	return &clusterConfigSourceFile{
		configPath:    path,
		baseDirectory: filepath.Dir(path),
		config:        &ClusterConfig{},
	}
}

func (cf *clusterConfigSourceFile) Refresh() (bool, error) {
	contents, err := ioutil.ReadFile(cf.configPath)
	if nil != err {
		return false, ERROR_READ_FILE
	}

	config := &ClusterConfig{}
	if err := yaml.Unmarshal(contents, config); nil != err {
		return false, ERROR_FILE_STRUCTURE
	}

	if err := config.refresh(cf.baseDirectory); nil != err {
		return false, err
	}

	cf.config = config

	//todo check if config was changed since last refresh and return false if it wasn't
	return true, nil
}

func (cf *clusterConfigSourceFile) Config() *ClusterConfig {
	return cf.config
}
