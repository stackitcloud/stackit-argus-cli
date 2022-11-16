package delete

/*
 * Delete an alert rule.
 */

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
)

// AlertRulesCmd represents the alertRules command
var AlertRulesCmd = &cobra.Command{
	Use:   "alertRule <groupName> <alertName>",
	Short: "Delete alert rules.",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		// generate an url
		url := config.GetBaseUrl() + fmt.Sprintf("alertgroups/%s/alertrules/%s", args[0], args[1])

		// call command
		err := runCommand(url, "alert rule")
		cobra.CheckErr(err)
	},
}
