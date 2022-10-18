/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	logging "github.com/stackitcloud/stackit-argus-cli/internal/log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const ProjectId = "PROJECT_ID"

var cfgFile string

var logger = logging.New()

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{ //nolint:gochecknoglobals // CLI command
	Use:     "stackit-argus-cli",
	Short:   "Manage ARGUS resources",
	Long:    `Manage ARGUS resources, like instances with CRUD operations`,
	Version: "1.0.0",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() error {
	err := rootCmd.Execute()
	if err != nil {
		logger.Error("an error occurred", logging.String("err", err.Error()))

		return err
	}

	return nil
}

func init() { //nolint:gochecknoinits // cobra CLI
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.stackit-argus-cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.

	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".stackit-argus-cli" (without extension).
		viper.AddConfigPath(home)
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")
		viper.SetConfigName(".stackit-argus-cli")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		logger.Info("Configuration file is found", logging.String("Using config file", viper.ConfigFileUsed()))
	}
}

func getBaseUrl() string {
	var url string

	if viper.Get(loggedIn) == "" {
		pId := viper.GetString(ProjectId)
		url = "https://api-dev.stackit.cloud/argus-service/v1/projects/" + pId + "/instances"
	} else {
		url = "https://api-dev.stackit.cloud/argus-service/v1/projects/" + viper.GetString(loggedIn) + "/instances"
	}

	return url
}
