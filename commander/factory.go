package commander

import (
	"errors"

	"github.com/coreos/fleet/client"
	"github.com/obukhov/fleet-commander/appdef"
)

var (
	ERROR_CLUSTER_NAME_CONFLICT = errors.New("Cluster name conflict")
)

func APIClientByConfig(cluster appdef.Cluster) (client.API, error) {
	return client.GetClient(&client.ClientConfig{
		SSHConfig: &client.SSHConfig{
			Tunnel: "",
			SSHUserName: "",
			StrictHostKeyChecking: false,
			KnownHostsFile: "",
			SshTimeout: 3.0,
		},

		ClientDriver: client.ClientDriverAPI,
		EndPoint: cluster.Endpoint,
		ReqTimeout: 10.0,

		CAFile: "",
		CertFile: "",
		KeyFile: "",

		EtcdKeyPrefix: "",
		ExperimentalAPI: false,
	})
}

func ClusterMapByConfig(clusterConfigs []appdef.Cluster) (map[string]client.API, error) {
	result := make(map[string]client.API)

	for _, config := range clusterConfigs {
		if _, found := result[config.Name]; true == found {
			return nil, ERROR_CLUSTER_NAME_CONFLICT
		}

		apiClient, err := APIClientByConfig(config)
		if nil != err {
			return nil, err
		}

		result[config.Name] = apiClient
	}

	return result, nil
}
