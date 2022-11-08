package get

/*
 * Get plans.
 */

import (
	"fmt"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/utils"

	"github.com/spf13/cobra"
)

// PlansCmd represents the plans command
var PlansCmd = &cobra.Command{
	Use:   "plans",
	Short: "Get all plans.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// generate an url
		url := config.GetProjectUrl() + "plans"

		// print debug messages if debug mode is turned on
		if config.IsDebugMode() {
			fmt.Println("list plans command called")
			fmt.Printf("url to call - %s\n", url)
		}

		// get plans
		status, body := getRequest(url)

		// print response status
		utils.ResponseMessage(status, "plans", "get")

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
