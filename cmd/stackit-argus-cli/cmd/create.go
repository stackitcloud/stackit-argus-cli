/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/create"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:     "create",
	Short:   "A brief description of your command",
	Long:    "",
	Example: "",
	Run: func(cmd *cobra.Command, args []string) {
		logger.Info("create command was called")
	},
}

func init() {
	createCmd.AddCommand(create.InstanceCmd)
	createCmd.AddCommand(create.CredentialsCmd)
	createCmd.AddCommand(create.ReceiversCmd)
	createCmd.AddCommand(create.RoutesCmd)
	createCmd.AddCommand(create.AlertGroupsCmd)
	createCmd.AddCommand(create.AlertRulesCmd)
	createCmd.AddCommand(create.AlertRecordsCmd)
	createCmd.AddCommand(create.BackupCmd)
	createCmd.AddCommand(create.BackupRestoresCmd)
	createCmd.AddCommand(create.BackupSchedulesCmd)
	createCmd.AddCommand(create.LogsAlertGroupsCmd)
	createCmd.AddCommand(create.ScrapeConfigsCmd)
	createCmd.AddCommand(create.SystemCredentialsCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	createCmd.PersistentFlags().StringP("file", "f", "", "provide file with request body")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
