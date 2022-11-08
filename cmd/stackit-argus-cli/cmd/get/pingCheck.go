package get

/*
 * Get ping check.
 */

import (
	"fmt"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/utils"

	"github.com/spf13/cobra"
)

// PingCheckCmd represents the PingCheck command
var PingCheckCmd = &cobra.Command{
	Use:   "PingCheck",
	Short: "Get all ping checks configured.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// generate an url
		url := config.GetBaseUrl() + "ping-checks"

		// print debug messages if debug mode is turned on
		if config.IsDebugMode() {
			fmt.Println("get ping network command called")
			fmt.Printf("url to call - %s\n", url)
		}

		// get ping check
		status, body := getRequest(url)

		// print response status
		utils.ResponseMessage(status, "ping check", "get")

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
