package appdef

type UnitFile struct {
	Name    string
	Path    string
	Content string
}

type Application struct {
	Name string
	Unit []UnitFile
}

type Cluster struct {
	Endpoint string
}

type ClusterConfig struct {
	Clusters     map[string]Cluster
	Applications []Application
}
