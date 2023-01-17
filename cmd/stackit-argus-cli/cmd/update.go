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

// init subcommands
func init() {
	//updateCmd.AddCommand(update.InstanceCmd)
	//updateCmd.AddCommand(update.CredentialsCmd)
	//updateCmd.AddCommand(update.AlertConfigsCmd)
	//updateCmd.AddCommand(update.ReceiversCmd)
	//updateCmd.AddCommand(update.RoutesCmd)
	//updateCmd.AddCommand(update.AlertGroupsCmd)
	//updateCmd.AddCommand(update.AlertRulesCmd)
	//updateCmd.AddCommand(update.AlertRecordsCmd)
	//updateCmd.AddCommand(update.GrafanaConfigsCmd)
	//updateCmd.AddCommand(update.LogsAlertGroupsCmd)
	//updateCmd.AddCommand(update.LogsConfigsCmd)
	//updateCmd.AddCommand(update.MetricsStorageRetentionCmd)
	//updateCmd.AddCommand(update.ScrapeConfigsCmd)
	//updateCmd.AddCommand(update.TracesConfigsCmd)
	//updateCmd.AddCommand(update.AclCmd)
	subCommands := []update.UpdateSubCmd{
		{
			Use:      "scrapeConfigs <jobName>",
			Short:    "Update scrape configs.",
			Long:     "Patch scrape configs if job name was not specified, otherwise update scrape config.",
			Args:     cobra.MaximumNArgs(1),
			MaxArgs:  1,
			Resource: "scrape config",
			BodyFile: true,
			Path:     []string{"scrapeconfigs"},
		},
		{
			Use:      "alertRecords <groupName> <alertRecord>",
			Short:    "Update alert records.",
			Long:     "Patch alert records if alert record was not specified, otherwise update alert record.",
			Args:     cobra.RangeArgs(1, 2),
			MaxArgs:  2,
			Resource: "alert record",
			BodyFile: true,
			Path:     []string{"alertgroups", "records"},
		},
		{
			Use:      "acl",
			Short:    "Update an acl config.",
			Args:     cobra.NoArgs,
			MaxArgs:  0,
			Resource: "acl",
			BodyFile: false,
			Path:     []string{"acl"},
		},
		{
			Use:      "alertConfig",
			Short:    "Update an alert config.",
			Args:     cobra.NoArgs,
			MaxArgs:  0,
			Resource: "alert config",
			BodyFile: true,
			Path:     []string{"alertconfigs"},
		},
		{
			Use:      "alertGroups <groupName>",
			Short:    "Update alert groups.",
			Long:     "Patch alert groups if group name was not specified, otherwise update alert group config.",
			Args:     cobra.MaximumNArgs(1),
			MaxArgs:  1,
			Resource: "alert group",
			BodyFile: true,
			Path:     []string{"alertgroups"},
		},
		{
			Use:      "alertRules <groupName> <alertName>",
			Short:    "Update alert rules.",
			Long:     "Patch alert rules if alert name was not specified, otherwise update an alert rule.",
			Args:     cobra.RangeArgs(1, 2),
			MaxArgs:  2,
			Resource: "alert rule",
			BodyFile: true,
			Path:     []string{"alertgroups", "alertrules"},
		},
		{
			Use:      "credentials <username>",
			Short:    "Update remote write config for credentials.",
			Long:     "Patch alert records if alert record was not specified, otherwise update alert record.",
			Args:     cobra.ExactArgs(1),
			MaxArgs:  1,
			Resource: "alert record",
			BodyFile: true,
			Path:     []string{"alertgroups", "records"},
		},
	}

	for _, command := range subCommands {
		updateCmd.AddCommand(update.GenerateUpdateSubCmd(command))
	}
}
