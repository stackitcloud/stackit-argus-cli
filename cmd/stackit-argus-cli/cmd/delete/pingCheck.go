package delete

/*
 * Delete ping check.
 */

import (
	"fmt"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/utils"

	"github.com/spf13/cobra"
)

// PingCheckCmd represents the PingCheck command
var PingCheckCmd = &cobra.Command{
	Use:   "pingCheck",
	Short: "Delete a ping check.",
	Run: func(cmd *cobra.Command, args []string) {
		// generate an url
		url := config.GetBaseUrl() + fmt.Sprintf("ping-checks/%s", args[0])

		// print debug messages if debug mode is turned on
		if config.IsDebugMode() {
			fmt.Println("delete ping check command called")
			fmt.Printf("url to call - %s\n", url)
		}

		// delete the ping check
		status := deleteRequest(url)

		// print response status
		utils.ResponseMessage(status, "ping check", "delete")
	},
}
