package get

/*
 * Get alert configs receivers.
 */

import (
	"fmt"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/utils"

	"github.com/spf13/cobra"
)

// ReceiversCmd represents the receivers command
var ReceiversCmd = &cobra.Command{
	Use:   "receivers <receiver>",
	Short: "Get alert config receivers.",
	Long:  "Get list of alert config receivers if receiver was not specified, otherwise get alert config receiver.",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var debugMsg string

		// generate an url
		url := config.GetBaseUrl() + "alertconfigs/receivers"

		// modify url and debug message depend on arguments
		if len(args) == 1 {
			debugMsg = "get alert config receiver command called"
			url += fmt.Sprintf("/%s", args[0])
		} else if len(args) == 0 {
			debugMsg = "list alert config receivers command called"
		}

		// print debug messages if debug mode is turned on
		if config.IsDebugMode() {
			fmt.Println(debugMsg)
			fmt.Printf("url to call - %s\n", url)
		}

		// get alert config receivers
		status, body := getRequest(url)

		// print response status
		utils.ResponseMessage(status, "alert config receivers", "get")

		// print response body
		if status == 200 {
			outputType := config.GetOutputType()
			if outputType == "json" || outputType == "yaml" {
				utils.PrintYamlOrJson(body, string(outputType))
			} else {
				fmt.Println(body)
			}
		}
	},
}
