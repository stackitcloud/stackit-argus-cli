/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	logging "github.com/stackitcloud/stackit-argus-cli/internal/log"
)

var projectName string
var projectId string

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:     "register",
	Short:   "A brief description of your command",
	Long:    "",
	Example: "",
	Run: func(cmd *cobra.Command, args []string) {
		if projectName == "" || projectId == "" {
			logger.Info("Please set both the name and id flag")

			return
		} else {
			projects := viper.Sub("projects")
			if projects == nil {
				logger.Error("projects configuration is not found")

				return
			}
			viper.Set("projects."+projectName, projectId)
			if err := viper.WriteConfigAs(viper.ConfigFileUsed()); err != nil {
				logger.Fatal("fatal saving project", logging.String("err", err.Error()))
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// registerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// registerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	registerCmd.Flags().StringVarP(&projectName, "projectName", "n", "", "specify the name of the project to save")
	registerCmd.Flags().StringVarP(&projectId, "projectId", "i", "", "specify the id of the project to save")
}
