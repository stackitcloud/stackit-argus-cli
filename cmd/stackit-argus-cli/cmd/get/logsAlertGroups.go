package get

/*
 * Get logs alert groups.
 */

import (
	"fmt"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/utils"

	"github.com/spf13/cobra"
)

// LogsAlertGroupsCmd represents the logsAlertGroups command
var LogsAlertGroupsCmd = &cobra.Command{
	Use:   "logsAlertGroups <groupName>",
	Short: "Get logs alert groups config.",
	Long:  "Get list of logs alert groups config if group name was not specified, otherwise get logs alert group config.",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var debugMsg string

		// generate an url
		url := config.GetBaseUrl() + "logs-alertgroups"

		// modify url and debug message depend on arguments
		if len(args) == 0 {
			debugMsg = "list logs alert groups command called"
		} else if len(args) == 1 {
			fmt.Println("get logs alert group config")
			debugMsg = "get logs alert group command called"
			url += fmt.Sprintf("/%s", args[0])
		}

		// print debug messages if debug mode is turned on
		if config.IsDebugMode() {
			fmt.Println(debugMsg)
			fmt.Printf("url to call - %s\n", url)
		}

		// get logs alert groups
		status, body := getRequest(url)

		// print response status
		utils.ResponseMessage(status, "logs alert groups", "get")

		// print response body
		if status == 200 {
			fmt.Println(body)
		}
	},
}
