package main

import (
	"fmt"
	"os"

	"github.com/obukhov/fleet-commander/appdef"
)

var VERSION string // Makefile sets this using linker flag, must be uninitialized

func main() {
	clusterConfig := appdef.NewClusterConfigSourceFile("../../example/cluster.yaml")

	if err := clusterConfig.Refresh(); nil != err {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println(clusterConfig.Config())
}
