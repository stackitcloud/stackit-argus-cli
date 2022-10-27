package create

/*
 * Create an instance.
 */

import (
	"fmt"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/utils"

	"github.com/spf13/cobra"
)

// InstanceCmd represents the instance command
var InstanceCmd = &cobra.Command{
	Use:   "instance",
	Short: "Create a new instance.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// generate an url
		url := config.GetInstancesUrl()

		// print debug messages if debug mode is turned on
		if config.IsDebugMode() {
			fmt.Println("create instance command called")
			fmt.Printf("url to call - %s\n", url)
		}

		// create the instance
		status := postRequest(url, nil)

		// print response status
		utils.ResponseMessage(status, "instance", "create")
	},
}
