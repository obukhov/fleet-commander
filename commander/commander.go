package commander

import (
	"github.com/coreos/fleet/client"
	"github.com/obukhov/fleet-commander/appdef"
	"errors"
	"github.com/coreos/fleet/schema"
)

var (
	ERROR_CLUSTER_NOT_DEFINED = errors.New("Cluser not defined")
)

type Commander struct {
	clusterConfigSource appdef.ClusterConfigSource
	clusterClients      map[string]client.API
}

func NewCommander(ccs appdef.ClusterConfigSource) *Commander {
	return &Commander{
		clusterConfigSource:ccs,
	}
}

func (c *Commander) Refresh() error {
	c.clusterConfigSource.Refresh()
	clusters, err := ClusterMapByConfig(c.clusterConfigSource.Config().Clusters)
	if nil != err {
		return err
	}

	c.clusterClients = clusters

	return nil
}

func (c *Commander) Check() ([]AppStatus, error) {
	result := make([]AppStatus, 0)
	if err := c.Refresh(); nil != err {
		return result, err
	}

	for _, app := range c.clusterConfigSource.Config().Apps {
		appCheckStatus := AppStatus{}
		appCheckStatus.Name = app.Name

		clusterAPIClient, clusterFound := c.clusterClients[app.Cluster]
		if false == clusterFound {
			return result, ERROR_CLUSTER_NOT_DEFINED
		}

		actualUnitsSlice, err := clusterAPIClient.Units()
		if nil != err {
			return result, err
		}

		actualUnits := make(map[string]*schema.Unit)
		for _, actualUnit := range actualUnitsSlice {
			actualUnits[actualUnit.Name] = actualUnit
		}

		for _, unit := range app.Units {
			unitStatus := UnitStatus{Name:unit.Name, IsFound:false, IsContentIdentical:false}

			if actualUnit, unitFound := actualUnits[unit.Name]; true == unitFound {
				unitStatus.IsFound = true
				unitStatus.Status = actualUnit.DesiredState

				if unit.Content == schema.MapSchemaUnitOptionsToUnitFile(actualUnit.Options).String() {
					unitStatus.IsContentIdentical = true
				}
			}

			appCheckStatus.UnitStatus = append(appCheckStatus.UnitStatus, unitStatus)
		}

		result = append(result, appCheckStatus)
	}

	return result, nil
}
