package create

/*
 * Create an alert rule.
 */

import (
	"fmt"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/utils"

	"github.com/spf13/cobra"
)

// AlertRulesCmd represents the alertRules command
var AlertRulesCmd = &cobra.Command{
	Use:   "alertRule <groupName>",
	Short: "Create an alert rule.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// generate an url
		url := config.GetBaseUrl() + fmt.Sprintf("alertgroups/%s/alertrules", args[0])

		// print debug messages if debug mode is turned on
		if config.IsDebugMode() {
			fmt.Println("create alert rule command called")
			fmt.Printf("url to call - %s\n", url)
		}

		// create the alert rule
		status := postRequest(url, nil)

		// print response status
		utils.ResponseMessage(status, "alert rule", "create")
	},
}
