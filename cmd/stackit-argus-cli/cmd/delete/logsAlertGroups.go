package delete

/*
 * Delete a logs alert group.
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
	Short: "Delete logs alert group config.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// generate an url
		url := config.GetBaseUrl() + fmt.Sprintf("logs-alertgroups/%s", args[0])

		// print url if debug mode is turned on
		if config.IsDebugMode() {
			fmt.Println("delete logs alert group config command called")
			fmt.Printf("url to call - %s\n", url)
		}

		// delete the logs alert group
		status := deleteRequest(url)

		// print response status
		utils.ResponseMessage(status, "logs alert group", "delete")
	},
}
