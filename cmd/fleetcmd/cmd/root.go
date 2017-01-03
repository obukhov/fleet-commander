package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

var cfgFile string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "fleetcmd",
	Short: "Tool for managing coreos/fleet cluster applications",
	Long:  ``,

	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	RootCmd.Flags().StringP("working-dir", "w", "", "Working dir path")

	cobra.OnInitialize(initConfig, initDefaultFlags)

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags, which, if defined here,
	// will be global for your application.

	//RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.fleetcmd.yaml)")
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
}

func initDefaultFlags() {
	log.Println("initDefaultFlags", RootCmd.Flag("working-dir"))
	if RootCmd.Flag("working-dir").Value.String() == "" {
		log.Println("RootCmd.Flag(\"working-dir\") is nil")
		wd, _ := os.Getwd()
		RootCmd.Flag("working-dir").Value.Set(wd)
	}

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	//viper.SetConfigName(".fleetcmd") // name of config file (without extension)
	//viper.AddConfigPath("$HOME")  // adding home directory as first search path
	//viper.AutomaticEnv()          // read in environment variables that match
	//
	//// If a config file is found, read it in.
	//if err := viper.ReadInConfig(); err == nil {
	//	fmt.Println("Using config file:", viper.ConfigFileUsed())
	//}
}
