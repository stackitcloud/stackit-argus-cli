package get

/*
 * Get alert configs.
 */

import (
	"fmt"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/utils"

	"github.com/spf13/cobra"
)

// AlertConfigsCmd represents the alertConfigs command
var AlertConfigsCmd = &cobra.Command{
	Use:   "alertConfigs",
	Short: "Get alert configs.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// generate an url
		url := config.GetBaseUrl() + "alertconfigs"

		// print debug messages if debug mode is turned on
		if config.IsDebugMode() {
			fmt.Println("get alert configs command called")
			fmt.Printf("url to call - %s\n", url)
		}

		// get alert configs
		status, body := getRequest(url)

		// print response status
		utils.ResponseMessage(status, "alert configs", "get")

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
