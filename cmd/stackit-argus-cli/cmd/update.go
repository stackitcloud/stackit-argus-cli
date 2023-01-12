package cmd

/*
 * Update subcommand command implementation (stackit-argus-cli update).
 */

import (
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/update"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update an ARGUS resource.",
	Args:  cobra.NoArgs,
}

// init subcommands and flags
func init() {
	// init subcommands
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
}
