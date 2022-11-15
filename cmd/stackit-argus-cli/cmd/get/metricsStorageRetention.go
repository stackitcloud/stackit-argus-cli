package get

/*
 * Get metrics storage retention.
 */

import (
	"encoding/json"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/utils"

	"github.com/spf13/cobra"
)

// metricsStorageRetention is used to unmarshal metrics storage retention response body
type metricsStorageRetention struct {
	MetricsRetentionTimeRaw string `json:"metricsRetentionTimeRaw" header:"metrics retention time raw"`
	MetricsRetentionTime5M  string `json:"metricsRetentionTime5m" header:"metrics retention time 5m"`
	MetricsRetentionTime1H  string `json:"metricsRetentionTime1h" header:"metrics retention time 1h"`
}

// printMetricsStorageRetentionTable prints routes response body as table
func printMetricsStorageRetentionTable(body []byte) {
	var metricsStorageRetention metricsStorageRetention

	// unmarshal response body
	err := json.Unmarshal(body, &metricsStorageRetention)
	cobra.CheckErr(err)

	// print the table
	utils.PrintTable(metricsStorageRetention)
}

// MetricsStorageRetentionCmd represents the metricsStorageRetention command
var MetricsStorageRetentionCmd = &cobra.Command{
	Use:   "metricsRetention",
	Short: "Get metric storage retention time.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// generate an url
		url := config.GetBaseUrl() + "metrics-storage-retentions"

		// get output flag
		outputType := config.GetOutputType()

		// call the command
		body := runCommand(url, "metrics storage retentions", outputType)

		// print table output
		if body != nil && (outputType == "" || outputType == "wide") {
			printMetricsStorageRetentionTable(body)
		}
	},
}
