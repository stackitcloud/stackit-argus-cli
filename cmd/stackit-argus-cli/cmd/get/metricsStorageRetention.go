package get

/*
 * Get metrics storage retention.
 */

import (
	"fmt"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/utils"

	"github.com/spf13/cobra"
)

// MetricsStorageRetentionCmd represents the metricsStorageRetention command
var MetricsStorageRetentionCmd = &cobra.Command{
	Use:   "metricsStorageRetention",
	Short: "Get metric storage retention time.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// generate an url
		url := config.GetBaseUrl() + "metrics-storage-retentions"

		// print debug messages if debug mode is turned on
		if config.IsDebugMode() {
			fmt.Println("list metrics storage retentions command called")
			fmt.Printf("url to call - %s\n", url)
		}

		// get metrics storage retentions
		status, body := getRequest(url)

		// print response status
		utils.ResponseMessage(status, "metrics storage retentions", "get")

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
