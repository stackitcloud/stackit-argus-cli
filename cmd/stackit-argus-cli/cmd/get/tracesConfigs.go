package get

/*
 * Get traces configs.
 */

import (
	"fmt"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/utils"

	"github.com/spf13/cobra"
)

// TracesConfigsCmd represents the tracesConfigs command
var TracesConfigsCmd = &cobra.Command{
	Use:   "tracesConfigs",
	Short: "Get traces config.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// generate an url
		url := config.GetBaseUrl() + "traces-configs"

		// print debug messages if debug mode is turned on
		if config.IsDebugMode() {
			fmt.Println("list traces configs command called")
			fmt.Printf("url to call - %s\n", url)
		}

		// get traces configs
		status, body := getRequest(url)

		// print response status
		utils.ResponseMessage(status, "traces configs", "get")

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
