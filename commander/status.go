package commander

type UnitStatus struct {
	Name               string
	IsFound            bool
	IsContentIdentical bool
	Status             string
}

type AppStatus struct {
	Name       string
	UnitStatus map[string]UnitStatus
}
