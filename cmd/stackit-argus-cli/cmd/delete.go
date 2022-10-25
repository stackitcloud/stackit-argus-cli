package cmd

/*
 * Delete subcommand command implementation (stackit-argus-cli delete).
 */

import (
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/delete"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an ARGUS resource.",
	Args:  cobra.NoArgs,
}

// init subcommands and flags
func init() {
	// init subcommands
	deleteCmd.AddCommand(delete.InstanceCmd)
	deleteCmd.AddCommand(delete.CredentialsCmd)
	deleteCmd.AddCommand(delete.ReceiversCmd)
	deleteCmd.AddCommand(delete.RoutesCmd)
	deleteCmd.AddCommand(delete.AlertGroupsCmd)
	deleteCmd.AddCommand(delete.AlertRulesCmd)
	deleteCmd.AddCommand(delete.AlertRecordsCmd)
	deleteCmd.AddCommand(delete.LogsAlertGroupsCmd)
	deleteCmd.AddCommand(delete.ScrapeConfigsCmd)
}
