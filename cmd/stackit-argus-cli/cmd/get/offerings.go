package get

/*
 * Get offerings.
 */

import (
	"fmt"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/utils"

	"github.com/spf13/cobra"
)

// OfferingsCmd represents the plansOfferings command
var OfferingsCmd = &cobra.Command{
	Use:   "offerings",
	Short: "Get all offerings.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// generate an url
		url := config.GetProjectUrl() + "offerings"

		// print debug messages if debug mode is turned on
		if config.IsDebugMode() {
			fmt.Println("list offerings command called")
			fmt.Printf("url to call - %s\n", url)
		}

		// get offerings
		status, body := getRequest(url)

		// print response status
		utils.ResponseMessage(status, "offerings", "get")

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
