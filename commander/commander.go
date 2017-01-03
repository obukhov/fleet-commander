package commander

import (
	"errors"
	"github.com/coreos/fleet/client"
	"github.com/coreos/fleet/schema"
	"github.com/coreos/fleet/unit"
	"github.com/obukhov/fleet-commander/appdef"
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
		clusterConfigSource: ccs,
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

func (c *Commander) Check() (map[string]AppStatus, error) {
	result := make(map[string]AppStatus, 0)
	if err := c.Refresh(); nil != err {
		return result, err
	}

	for _, app := range c.clusterConfigSource.Config().Apps {
		appCheckStatus := AppStatus{Name: app.Name, UnitStatus: make(map[string]UnitStatus)} // todo constructor
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

		for _, defUnit := range app.Units {
			unitStatus := UnitStatus{Name: defUnit.Name, IsFound: false, IsContentIdentical: false}

			if actualUnit, unitFound := actualUnits[defUnit.Name]; true == unitFound {
				unitStatus.IsFound = true
				unitStatus.Status = actualUnit.DesiredState

				if defUnit.Content == schema.MapSchemaUnitOptionsToUnitFile(actualUnit.Options).String() {
					unitStatus.IsContentIdentical = true
				}
			}

			appCheckStatus.UnitStatus[defUnit.Name] = unitStatus
		}

		result[app.Name] = appCheckStatus
	}

	return result, nil
}

func (c *Commander) Fix(appStatuses map[string]AppStatus) error {
	for _, app := range c.clusterConfigSource.Config().Apps {
		clusterAPIClient, clusterFound := c.clusterClients[app.Cluster]
		if false == clusterFound {
			return ERROR_CLUSTER_NOT_DEFINED
		}

		for _, defUnit := range app.Units {
			status := appStatuses[app.Name].UnitStatus[defUnit.Name]

			if false == status.IsFound || false == status.IsContentIdentical {
				uf, err := unit.NewUnitFile(defUnit.Content)

				if nil != err {
					return err
				}

				options := schema.MapUnitFileToSchemaUnitOptions(uf)
				schemaUnit := &schema.Unit{
					Name:         defUnit.Name,
					Options:      options,
					DesiredState: "launched",
				}

				if err := clusterAPIClient.CreateUnit(schemaUnit); nil != err {
					return err
				}
			}
		}
	}

	return nil
}
