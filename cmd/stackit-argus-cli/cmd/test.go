/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	logging "github.com/stackitcloud/stackit-argus-cli/internal/log"
	"os"
)

var key string

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:     "test",
	Short:   "A brief description of your command",
	Long:    "",
	Example: "",
	Run: func(cmd *cobra.Command, args []string) {
		logger.Info("test command was called", logging.Any("Environment Variables", os.Environ()))
	},
}

func init() {
	rootCmd.AddCommand(testCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// testCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	testCmd.Flags().StringVarP(&key, "key", "k", "", "specify the key to retrieve")
}
