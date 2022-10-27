package get

/*
 * Get alert groups.
 */

import (
	"fmt"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/utils"

	"github.com/spf13/cobra"
)

// AlertGroupsCmd represents the alertGroups command
var AlertGroupsCmd = &cobra.Command{
	Use:   "alertGroups <groupName>",
	Short: "Get list of alert groups.",
	Long:  "Get list of alert groups if group name was not specified, otherwise get alert group.",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var debugMsg string

		// generate an url
		url := config.GetBaseUrl() + "alertgroups"

		// modify url and debug message depend on arguments
		if len(args) == 1 {
			debugMsg = "get alert group command called"
			url += fmt.Sprintf("/%s", args[0])
		} else if len(args) == 0 {
			debugMsg = "list alert groups command called"
		}

		// print debug messages if debug mode is turned on
		if config.IsDebugMode() {
			fmt.Println(debugMsg)
			fmt.Printf("url to call - %s\n", url)
		}

		// get alert groups
		status, body := getRequest(url)

		// print response status
		utils.ResponseMessage(status, "alert groups", "get")

		// print response body
		if status == 200 {
			fmt.Println(body)
		}
	},
}
