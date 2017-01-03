package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// cleanCmd represents the clean command
var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "",
	Long:  ``,
	Run:   runCleanCmd,
}

func init() {
	RootCmd.AddCommand(cleanCmd)
}

func runCleanCmd(cmd *cobra.Command, args []string) {
	fmt.Println("clean called")
}
