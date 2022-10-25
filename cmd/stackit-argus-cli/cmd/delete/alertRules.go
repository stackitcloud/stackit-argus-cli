package delete

/*
 * Delete an alert rule.
 */

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/utils"
)

// AlertRulesCmd represents the alertRules command
var AlertRulesCmd = &cobra.Command{
	Use:   "alertRule <groupName> <alertName>",
	Short: "Delete alert rules.",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		// generate an url
		url := config.GetBaseUrl() + fmt.Sprintf("alertgroups/%s/alertrules/%s", args[0], args[1])

		// print debug messages if debug mode is turned on
		if config.IsDebugMode() {
			fmt.Println("delete alert rule command called")
			fmt.Printf("url to call - %s\n", url)
		}

		// delete the alert rule
		status := deleteRequest(url)

		// print response status
		utils.ResponseMessage(status, "alert rule", "delete")
	},
}
