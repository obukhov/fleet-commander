package appdef

import (
	"errors"
	"io/ioutil"

	"path/filepath"

	"gopkg.in/yaml.v2"
)

var (
	ERROR_READ_FILE      = errors.New("Error reading file")
	ERROR_FILE_STRUCTURE = errors.New("File has wrong structure")
)

type ClusterConfigSource interface {
	Config() *ClusterConfig
	Refresh() error
}

type clusterConfigSourceFile struct {
	configPath    string
	baseDirectory string
	config        *ClusterConfig
}

func NewClusterConfigSourceFile(path string) ClusterConfigSource {
	return &clusterConfigSourceFile{
		configPath:    path,
		baseDirectory: filepath.Dir(path),
		config:        &ClusterConfig{},
	}
}

func (cf *clusterConfigSourceFile) Refresh() error {
	contents, err := ioutil.ReadFile(cf.configPath)
	if nil != err {
		return ERROR_READ_FILE
	}

	config := &ClusterConfig{}
	if err := yaml.Unmarshal(contents, config); nil != err {
		return ERROR_FILE_STRUCTURE
	}

	if err := config.refresh(cf.baseDirectory); nil != err {
		return err
	}

	cf.config = config

	return nil
}

func (cf *clusterConfigSourceFile) Config() *ClusterConfig {
	return cf.config
}
