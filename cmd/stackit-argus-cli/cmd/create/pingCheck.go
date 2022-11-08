package create

/*
 * Create ping check.
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
	Short: "Create a ping check.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// generate an url
		url := config.GetBaseUrl() + "ping-checks"

		// print debug messages if debug mode is turned on
		if config.IsDebugMode() {
			fmt.Println("create ping check command called")
			fmt.Printf("url to call - %s\n", url)
		}

		// create the ping check
		status := postRequest(url, nil)

		// print response status
		utils.ResponseMessage(status, "ping check", "create")
	},
}
