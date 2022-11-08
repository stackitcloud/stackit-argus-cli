package delete

/*
 * Delete http check.
 */

import (
	"fmt"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/utils"

	"github.com/spf13/cobra"
)

// HttpCheckCmd represents the HttpCheck command
var HttpCheckCmd = &cobra.Command{
	Use:   "httpCheck",
	Short: "Delete a http check.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// generate an url
		url := config.GetBaseUrl() + fmt.Sprintf("http-checks/%s", args[0])

		// print debug messages if debug mode is turned on
		if config.IsDebugMode() {
			fmt.Println("delete http check command called")
			fmt.Printf("url to call - %s\n", url)
		}

		// delete the http check
		status := deleteRequest(url)

		// print response status
		utils.ResponseMessage(status, "http check", "delete")
	},
}
