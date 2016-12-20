package main

import (
	"fmt"
	"os"

	"github.com/obukhov/fleet-commander/appdef"
	"github.com/obukhov/fleet-commander/commander"
)

var VERSION string // Makefile sets this using linker flag, must be uninitialized

func main() {
	clusterConfig := appdef.NewClusterConfigSourceFile("../../example/cluster.yaml")
	cmdr := commander.NewCommander(clusterConfig)

	appStatus, err := cmdr.Check()

	if nil != err {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println(appStatus)
}