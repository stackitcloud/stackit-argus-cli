package create

/*
 * Create an alert config receiver.
 */

import (
	"fmt"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/utils"

	"github.com/spf13/cobra"
)

// ReceiversCmd represents the receivers command
var ReceiversCmd = &cobra.Command{
	Use:   "receiver",
	Short: "Create alert config receiver.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// generate an url
		url := config.GetBaseUrl() + "alertconfigs/receivers"

		// print debug messages if debug mode is turned on
		if config.IsDebugMode() {
			fmt.Println("create alert config receiver command called")
			fmt.Printf("url to call - %s\n", url)
		}

		// create the receiver
		status := postRequest(url, nil)

		// print response status
		utils.ResponseMessage(status, "receiver", "create")
	},
}
