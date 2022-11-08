package get

/*
 * Get scrape configs.
 */

import (
	"fmt"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/utils"

	"github.com/spf13/cobra"
)

// ScrapeConfigsCmd represents the scrapeConfigs command
var ScrapeConfigsCmd = &cobra.Command{
	Use:   "scrapeConfigs <jobName>",
	Short: "Get scrape config.",
	Long:  "Get list of scrape config if job name was not specified, otherwise get scrape config.",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var debugMsg string

		// generate an url
		url := config.GetBaseUrl() + "scrapeconfigs"

		// modify url and debug message depend on arguments
		if len(args) == 1 {
			debugMsg = "get scrape config command called"
			url += fmt.Sprintf("/%s", args[0])
		} else if len(args) == 0 {
			debugMsg = "list scrape configs command called"
		}

		// print debug messages if debug mode is turned on
		if config.IsDebugMode() {
			fmt.Println(debugMsg)
			fmt.Printf("url to call - %s\n", url)
		}

		// get scrape configs
		status, body := getRequest(url)

		// print response status
		utils.ResponseMessage(status, "scrape configs", "get")

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
