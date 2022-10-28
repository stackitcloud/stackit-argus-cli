package update

/*
 * Update alert rules.
 */

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/utils"
)

// AlertRulesCmd represents the alertRules command
var AlertRulesCmd = &cobra.Command{
	Use:   "alertRules <groupName> <alertName>",
	Short: "Update alert rules.",
	Long:  "Patch alert rules if alert name was not specified, otherwise update an alert rule.",
	Args:  cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		var method string
		var debugMessage string

		// generate an url
		url := config.GetBaseUrl() + fmt.Sprintf("alertgroups/%s/alertrules", args[0])

		// modify url and debug msg if an argument has been given
		if len(args) == 1 {
			debugMessage = "patch alert rules command called"
			method = "PATCH"
		} else if len(args) == 2 {
			debugMessage = "update alert rule command called"
			method = "PUT"
			url += fmt.Sprintf("/%s", args[1])
		}

		// print debug messages if debug mode is turned on
		if config.IsDebugMode() {
			fmt.Println(debugMessage)
			fmt.Printf("url to call - %s\n", url)
		}

		// update the alert rule
		status := updateRequest(url, method)

		// print response status
		utils.ResponseMessage(status, "alert rule", "update")
	},
}
