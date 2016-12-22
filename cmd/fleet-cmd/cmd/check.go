package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/obukhov/fleet-commander/appdef"
	"github.com/obukhov/fleet-commander/commander"
	"os"
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check [application name]",
	Short: "Check running status of application",
	Long: `Check running status of application. If application name is omitted all aplications in config files are
checked. Only outputs result of check, no changes in cluster are done.`,
	Run: runCheckCommand,
}

func init() {
	RootCmd.AddCommand(checkCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// checkCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// checkCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

func runCheckCommand(cmd *cobra.Command, args []string) {
	clusterConfig := appdef.NewClusterConfigSourceFile("../../example/cluster.yaml")
	cmdr := commander.NewCommander(clusterConfig)

	appStatus, err := cmdr.Check()

	if nil != err {
		fmt.Println("Error checking:", err)
		os.Exit(1)
	}

	// todo move to fix command
	if err := cmdr.Fix(appStatus); nil != err {
		fmt.Println("Error fixing:", err)
		os.Exit(1)
	}

	fmt.Println(appStatus)
}
