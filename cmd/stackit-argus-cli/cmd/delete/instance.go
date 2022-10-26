package delete

/*
 * Delete an instance.
 */

import (
	"fmt"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/utils"

	"github.com/spf13/cobra"
)

// InstanceCmd represents the instance command
var InstanceCmd = &cobra.Command{
	Use:   "instance <instanceId>",
	Short: "Delete an instance.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// generate an url
		url := config.GetInstancesUrl() + args[0]

		// print debug messages if debug mode is turned on
		if config.IsDebugMode() {
			fmt.Println("delete instance command called")
			fmt.Printf("url to call - %s\n", url)
		}

		// delete the instance
		status := deleteRequest(url)

		// print response status
		utils.ResponseMessage(status, "instance", "delete")
	},
}
