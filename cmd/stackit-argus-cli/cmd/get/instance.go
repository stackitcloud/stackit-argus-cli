package get

/*
 * Get instances.
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
	Short: "Get project instance.",
	Long:  "Get list of all project's instances if instance id was not specified, otherwise get instance.",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var debugMsg string

		// generate an url
		url := config.GetInstancesUrl()

		// modify url and debug message depend on argument and flag
		if len(args) == 0 {
			debugMsg = "list instances command called"
		} else if len(args) == 1 {
			debugMsg = "get instance command called"
			url += fmt.Sprintf("/%s", args[0])
		}

		// print debug messages if debug mode is turned on
		if config.IsDebugMode() {
			fmt.Println(debugMsg)
			fmt.Printf("url to call - %s\n", url)
		}

		// get instances
		status, body := getRequest(url)

		// print response status
		utils.ResponseMessage(status, "instances", "get")

		// print response body
		if status == 200 {
			fmt.Println(body)
		}
	},
}
