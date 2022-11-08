package get

/*
 * Get http check.
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
	Short: "Get all http checks configured.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// generate an url
		url := config.GetBaseUrl() + "http-checks"

		// print debug messages if debug mode is turned on
		if config.IsDebugMode() {
			fmt.Println("get http check command called")
			fmt.Printf("url to call - %s\n", url)
		}

		// get http check
		status, body := getRequest(url)

		// print response status
		utils.ResponseMessage(status, "http check", "get")

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
