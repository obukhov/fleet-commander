package appdef

import (
	"errors"
	"io/ioutil"
	"path/filepath"
)

type UnitFile struct {
	Name    string
	Source  string
	Content string
}

type Application struct {
	Name    string
	Cluster string
	Units   []*UnitFile
}

type Cluster struct {
	Name        string
	Endpoint    string
	Tunnel      string
	SshUsername string `yaml:"ssh-username"`
}

type ClusterConfig struct {
	Clusters []Cluster
	Apps     []Application
}

var (
	ERROR_LOADING_UNIT_SOURCE = errors.New("Error loading unit source file")
)

func (cc *ClusterConfig) refresh(baseDirectory string) error {
	for _, app := range cc.Apps {
		for _, unit := range app.Units {
			if err := unit.load(baseDirectory); nil != err {
				return ERROR_LOADING_UNIT_SOURCE
			}

		}
	}

	return nil
}

func (uf *UnitFile) load(baseDirectory string) error {
	if "" != uf.Source {
		bytesContent, err := ioutil.ReadFile(filepath.Join(baseDirectory, uf.Source))
		if nil != err {
			return err
		}

		uf.Content = string(bytesContent)
	}

	return nil
}
