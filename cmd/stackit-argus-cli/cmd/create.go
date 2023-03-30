package cmd

/*
 * Create subcommand command implementation (stackit-argus-cli create).
 */

import (
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/create"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new ARGUS resource.",
	Args:  cobra.NoArgs,
}

// init subcommands
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
	createCmd.AddCommand(create.CertCheckCmd)
	createCmd.AddCommand(create.HttpCheckCmd)
	createCmd.AddCommand(create.NetworkCheckCmd)
	createCmd.AddCommand(create.PingCheckCmd)
	createCmd.AddCommand(create.GrafanaAnnotationCmd)
}
