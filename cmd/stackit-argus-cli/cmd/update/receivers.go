package update

/*
 * Update a receiver.
 */

import (
	"fmt"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/utils"

	"github.com/spf13/cobra"
)

// ReceiversCmd represents the receivers command
var ReceiversCmd = &cobra.Command{
	Use:   "receiver <receiver>",
	Short: "Update an alert config receiver.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// generate an url
		url := config.GetBaseUrl() + fmt.Sprintf("/alertconfigs/receivers/%s", args[0])

		// print debug messages if debug mode is turned on
		if config.IsDebugMode() {
			fmt.Println("update receiver command called")
			fmt.Printf("url to call - %s\n", url)
		}

		// update the receiver
		status := updateRequest(url, "PUT")

		// print response status
		utils.ResponseMessage(status, "receiver", "update")
	},
}
