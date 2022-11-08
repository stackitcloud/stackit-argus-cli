package create

/*
 * Create http check.
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
	Short: "Create a http check.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// generate an url
		url := config.GetBaseUrl() + "http-checks"

		// print debug messages if debug mode is turned on
		if config.IsDebugMode() {
			fmt.Println("create http check command called")
			fmt.Printf("url to call - %s\n", url)
		}

		// create the http check
		status := postRequest(url, nil)

		// print response status
		utils.ResponseMessage(status, "http check", "create")
	},
}
