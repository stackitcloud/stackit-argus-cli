package update

/*
 * Update alert rules.
 */

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
)

// AlertRulesCmd represents the alertRules command
var AlertRulesCmd = &cobra.Command{
	Use:   "alertRules <groupName> <alertName>",
	Short: "Update alert rules.",
	Long:  "Patch alert rules if alert name was not specified, otherwise update an alert rule.",
	Args:  cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		// generate an url
		url := config.GetBaseUrl() + fmt.Sprintf("alertgroups/%s/alertrules", args[0])

		// modify url and debug msg if an argument has been given
		method := "PATCH"
		resource := "alert rules"
		if len(args) == 2 {
			resource = "alert rule"
			method = "PUT"
			url += fmt.Sprintf("/%s", args[1])
		}

		// call command
		runCommand(url, resource, method)
	},
}
