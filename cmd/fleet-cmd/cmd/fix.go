package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// fixCmd represents the fix command
var fixCmd = &cobra.Command{
	Use:   "fix [application name]",
	Short: "",
	Long: ``,
	Run: runFixCmd,
}

func init() {
	RootCmd.AddCommand(fixCmd)
}

func runFixCmd(cmd *cobra.Command, args []string) {
	fmt.Println("fix called")
}
