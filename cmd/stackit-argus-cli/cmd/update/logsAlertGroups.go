package update

/*
 * Update a logs alert group config.
 */

import (
	"fmt"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/utils"

	"github.com/spf13/cobra"
)

// LogsAlertGroupsCmd represents the logsAlertGroups command
var LogsAlertGroupsCmd = &cobra.Command{
	Use:   "logsAlertGroup <groupName>",
	Short: "Update logs alert group config.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// generate an url
		url := config.GetBaseUrl() + fmt.Sprintf("logs-alertgroups/%s", args[0])

		// print debug messages if debug mode is turned on
		if config.IsDebugMode() {
			fmt.Println("update logs alert group config command called")
			fmt.Printf("url to call - %s\n", url)
		}

		// update the logs alert groups
		status := updateRequest(url, "PUT")

		// print response status
		utils.ResponseMessage(status, "logs alert groups", "update")
	},
}
