package delete

/*
 * Delete an alert rule.
 */

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/config"
)

// AlertRulesCmd represents the alertRules command
var AlertRulesCmd = &cobra.Command{
	Use:   "alertRule <groupName> <alertName>",
	Short: "Delete alert rules.",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		// generate an url
		url := config.GetBaseUrl() + fmt.Sprintf("alertgroups/%s/alertrules/%s", args[0], args[1])

		// call command
		if err := runCommand(url, "alert rule"); err != nil {
			cmd.SilenceUsage = true

			return err
		}

		return nil
	},
}
