/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/update"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:     "update",
	Short:   "A brief description of your command",
	Long:    "",
	Example: "",
	Run: func(cmd *cobra.Command, args []string) {
		logger.Info("update command was called")
	},
}

func init() {
	updateCmd.AddCommand(update.InstanceCmd)
	updateCmd.AddCommand(update.CredentialsCmd)
	updateCmd.AddCommand(update.AlertConfigsCmd)
	updateCmd.AddCommand(update.ReceiversCmd)
	updateCmd.AddCommand(update.RoutesCmd)
	updateCmd.AddCommand(update.AlertGroupsCmd)
	updateCmd.AddCommand(update.AlertRulesCmd)
	updateCmd.AddCommand(update.AlertRecordsCmd)
	updateCmd.AddCommand(update.GrafanaConfigsCmd)
	updateCmd.AddCommand(update.LogsAlertGroupsCmd)
	updateCmd.AddCommand(update.LogsConfigsCmd)
	updateCmd.AddCommand(update.MetricsStorageRetentionCmd)
	updateCmd.AddCommand(update.ScrapeConfigsCmd)
	updateCmd.AddCommand(update.TracesConfigsCmd)
	updateCmd.AddCommand(update.AclCmd)

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	updateCmd.PersistentFlags().StringP("file", "f", "", "provide file with request body")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
