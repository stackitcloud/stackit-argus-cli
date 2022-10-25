package delete

/*
 * Delete an alert group.
 */

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/utils"
)

// AlertGroupsCmd represents the alertGroups command
var AlertGroupsCmd = &cobra.Command{
	Use:   "alertGroup <groupName>",
	Short: "Delete alert group config.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// generate an url
		url := config.GetBaseUrl() + fmt.Sprintf("alertgroups/%s", args[0])

		// print url if debug mode is turned on
		if config.IsDebugMode() {
			fmt.Println("delete alert group command called")
			fmt.Printf("url to call - %s\n", url)
		}

		// delete the alert group
		status := deleteRequest(url)

		// print response status
		utils.ResponseMessage(status, "alert group", "delete")
	},
}
