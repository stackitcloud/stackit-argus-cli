/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/delete"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:     "delete",
	Short:   "A brief description of your command",
	Long:    "",
	Example: "",
	Run: func(cmd *cobra.Command, args []string) {
		logger.Info("delete command was called")
	},
}

func init() {
	deleteCmd.AddCommand(delete.InstanceCmd)
	deleteCmd.AddCommand(delete.CredentialsCmd)
	deleteCmd.AddCommand(delete.ReceiversCmd)
	deleteCmd.AddCommand(delete.RoutesCmd)
	deleteCmd.AddCommand(delete.AlertGroupsCmd)
	deleteCmd.AddCommand(delete.AlertRulesCmd)
	deleteCmd.AddCommand(delete.AlertRecordsCmd)
	deleteCmd.AddCommand(delete.LogsAlertGroupsCmd)
	deleteCmd.AddCommand(delete.ScrapeConfigsCmd)
	deleteCmd.AddCommand(delete.SystemCredentialsCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
