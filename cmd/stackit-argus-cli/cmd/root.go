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

var confFile string
var instanceId string
var projectId string
var debugMode bool

var logger = logging.New()

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{ //nolint:gochecknoglobals // CLI command
	Use:     "stackit-argus-cli",
	Short:   "Manage ARGUS resources",
	Long:    `Manage ARGUS resources, like instances with CRUD operations`,
	Version: "1.0.0",
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

	rootCmd.AddCommand(getCmd)
	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(updateCmd)
	rootCmd.AddCommand(deleteCmd)

	rootCmd.PersistentFlags().StringVarP(&confFile, "config", "c", "./.stackit-argus-cli.yaml", "provide config file (default is ./.stackit-argus-cli.yaml)")
	rootCmd.PersistentFlags().StringVarP(&instanceId, "instanceId", "i", "", "provide instance id that should be used")
	rootCmd.PersistentFlags().StringVarP(&projectId, "projectId", "p", "", "provide project id that should be used")
	rootCmd.PersistentFlags().BoolVarP(&debugMode, "debug", "d", false, "turn on debug mode to see more information about the command")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if confFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(confFile)
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
	if err := viper.ReadInConfig(); err != nil {
		logger.Error("cannot read config file", logging.String("err", err.Error()))
		//logger.Info("Configuration file is found", logging.String("Using config file", viper.ConfigFileUsed()))
	}
}
