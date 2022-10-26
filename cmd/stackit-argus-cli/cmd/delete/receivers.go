package delete

/*
 * Delete a receiver.
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
	Short: "Delete alert config receiver.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// generate an url
		url := config.GetBaseUrl() + fmt.Sprintf("alertconfigs/receivers/%s", args[0])

		// print debug messages if debug mode is turned on
		if config.IsDebugMode() {
			fmt.Println("delete alert config receiver command called")
			fmt.Printf("url to call - %s\n", url)
		}

		// delete the receiver
		status := deleteRequest(url)

		// print response status
		utils.ResponseMessage(status, "alert config receiver", "delete")
	},
}
