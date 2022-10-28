package update

/*
 * Update alert records.
 */

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/utils"
)

// AlertRecordsCmd represents the alertRecords command
var AlertRecordsCmd = &cobra.Command{
	Use:   "alertRecords <groupName> <alertRecord>",
	Short: "Update alert records.",
	Long:  "Patch alert records if alert record was not specified, otherwise update alert record.",
	Args:  cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		var method string
		var debugMessage string

		// generate an url
		url := config.GetBaseUrl() + fmt.Sprintf("alertgroups/%s/records", args[0])

		// modify url and debug msg if an argument has been given
		if len(args) == 1 {
			debugMessage = "patch alert records command called"
			method = "PATCH"
		} else if len(args) == 2 {
			debugMessage = "update alert record command called"
			method = "PUT"
			url += fmt.Sprintf("/%s", args[1])
		}

		// print debug messages if debug mode is turned on
		if config.IsDebugMode() {
			fmt.Println(debugMessage)
			fmt.Printf("url to call - %s\n", url)
		}

		// update the alert record
		status := updateRequest(url, method)

		// print response status
		utils.ResponseMessage(status, "alert record", "update")
	},
}
