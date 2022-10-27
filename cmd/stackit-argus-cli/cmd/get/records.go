package get

/*
 * Get alert records.
 */

import (
	"fmt"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/utils"

	"github.com/spf13/cobra"
)

// RecordsCmd represents the alertRecords command
var RecordsCmd = &cobra.Command{
	Use:   "records <groupName> <alertRecord>",
	Short: "Get alert records.",
	Long:  "Get list of alert records if alert record was not specified, otherwise get alert record.",
	Args:  cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		var debugMsg string

		// generate an url
		url := config.GetBaseUrl() + fmt.Sprintf("alertgroups/%s/records", args[0])

		// modify url and debug message depend on arguments
		if len(args) == 2 {
			debugMsg = "get alert record command called"
			url += fmt.Sprintf("/%s", args[1])
		} else if len(args) == 1 {
			debugMsg = "list alert records command called"
		}

		// print debug messages if debug mode is turned on
		if config.IsDebugMode() {
			fmt.Println(debugMsg)
			fmt.Printf("url to call - %s\n", url)
		}

		// get alert records
		status, body := getRequest(url)

		// print response status
		utils.ResponseMessage(status, "alert records", "get")

		// print response body
		if status == 200 {
			fmt.Println(body)
		}
	},
}
