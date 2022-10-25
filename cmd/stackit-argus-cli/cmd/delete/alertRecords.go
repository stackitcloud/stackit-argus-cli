package delete

/*
 * Delete an alert record.
 */

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/utils"
)

// AlertRecordsCmd represents the alertRecords command
var AlertRecordsCmd = &cobra.Command{
	Use:   "alertRecord <groupName> <alertRecord>",
	Short: "Delete alert records.",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		// generate an url
		url := config.GetBaseUrl() + fmt.Sprintf("alertgroups/%s/records/%s", args[0], args[1])

		// print url if debug mode is turned on
		if config.IsDebugMode() {
			fmt.Println("delete alert record command called")
			fmt.Printf("url to call - %s\n", url)
		}

		// delete the alert record
		status := deleteRequest(url)

		// print response status
		utils.ResponseMessage(status, "alert record", "delete")
	},
}
