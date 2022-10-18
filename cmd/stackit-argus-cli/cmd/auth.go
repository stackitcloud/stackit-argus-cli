/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/rest"
	logging "github.com/stackitcloud/stackit-argus-cli/internal/log"
)

var token string

// authCmd represents the auth command
var authCmd = &cobra.Command{
	Use:     "auth",
	Short:   "Authorize to the project",
	Long:    "Authorize to the project and shows credentials.",
	Example: "  stackit-argus-cli auth <projectId>",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		token = rest.Authorize(args[0])
		if token != "" {
			viper.Set("token", token)
			if err := viper.WriteConfigAs(viper.ConfigFileUsed()); err != nil {
				logger.Fatal("fatal saving token", logging.String("err", err.Error()))
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(authCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// authCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// authCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
