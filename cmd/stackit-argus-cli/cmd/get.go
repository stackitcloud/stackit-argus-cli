package cmd

/*
 * Get subcommand command implementation (stackit-argus-cli get).
 */

import (
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/get"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve information about an ARGUS resource.",
	Args:  cobra.NoArgs,
}

// init subcommands and flags
func init() {
	// add subcommands
	getCmd.AddCommand(get.InstanceCmd)
	getCmd.AddCommand(get.CredentialsCmd)
	getCmd.AddCommand(get.AlertConfigsCmd)
	getCmd.AddCommand(get.ReceiversCmd)
	getCmd.AddCommand(get.RoutesCmd)
	getCmd.AddCommand(get.AlertGroupsCmd)
	getCmd.AddCommand(get.AlertRulesCmd)
	getCmd.AddCommand(get.RecordsCmd)
	getCmd.AddCommand(get.BackupCmd)
	getCmd.AddCommand(get.BackupRetentionsCmd)
	getCmd.AddCommand(get.BackupSchedulesCmd)
	getCmd.AddCommand(get.GrafanaConfigsCmd)
	getCmd.AddCommand(get.LogsAlertGroupsCmd)
	getCmd.AddCommand(get.LogsConfigsCmd)
	getCmd.AddCommand(get.MetricsStorageRetentionCmd)
	getCmd.AddCommand(get.ScrapeConfigsCmd)
	getCmd.AddCommand(get.TracesConfigsCmd)
	getCmd.AddCommand(get.AclCmd)
	getCmd.AddCommand(get.PlansCmd)
	getCmd.AddCommand(get.OfferingsCmd)
	getCmd.AddCommand(get.CertCheckCmd)
	getCmd.AddCommand(get.HttpCheckCmd)
	getCmd.AddCommand(get.NetworkCheckCmd)
	getCmd.AddCommand(get.PingCheckCmd)
}
