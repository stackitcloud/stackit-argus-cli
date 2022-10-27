package get

/*
 * Get alert rules.
 */

import (
	"fmt"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/utils"

	"github.com/spf13/cobra"
)

// AlertRulesCmd represents the alertRules command
var AlertRulesCmd = &cobra.Command{
	Use:   "alertRules <groupName> <alertName>",
	Short: "Get alert rules.",
	Long:  "Get list of alert rules if alert name was not specified, otherwise get alert rule.",
	Args:  cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		var debugMsg string

		// generate an url
		url := config.GetBaseUrl() + fmt.Sprintf("alertgroups/%s/alertrules", args[0])

		// modify url and debug message depend on arguments
		if len(args) == 2 {
			debugMsg = "get alert rule command called"
			url += fmt.Sprintf("/%s", args[1])
		} else if len(args) == 1 {
			debugMsg = "list alert rules command called"
		}

		// print debug messages if debug mode is turned on
		if config.IsDebugMode() {
			fmt.Println(debugMsg)
			fmt.Printf("url to call - %s\n", url)
		}

		// get alert rules
		status, body := getRequest(url)

		// print response status
		utils.ResponseMessage(status, "alert rules", "get")

		// print response body
		if status == 200 {
			fmt.Println(body)
		}
	},
}
