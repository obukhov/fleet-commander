package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run [application name]",
	Short: "",
	Long:  ``,
	Run:   runRunCmd,
}

func init() {
	RootCmd.AddCommand(runCmd)
}

func runRunCmd(cmd *cobra.Command, args []string) {
	fmt.Println("run called")
}
