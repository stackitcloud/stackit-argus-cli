package get

/*
 * Get metrics storage retention.
 */

import (
	"encoding/json"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	output_table "github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/outputTable"
)

// metricsStorageRetention is used to unmarshal metrics storage retention response body
type metricsStorageRetention struct {
	MetricsRetentionTimeRaw string `json:"metricsRetentionTimeRaw" header:"metrics retention time raw"`
	MetricsRetentionTime5M  string `json:"metricsRetentionTime5m" header:"metrics retention time 5m"`
	MetricsRetentionTime1H  string `json:"metricsRetentionTime1h" header:"metrics retention time 1h"`
}

// printMetricsStorageRetentionTable prints routes response body as outputTable
func printMetricsStorageRetentionTable(body []byte) error {
	var metricsStorageRetention metricsStorageRetention

	// unmarshal response body
	if err := json.Unmarshal(body, &metricsStorageRetention); err != nil {
		return err
	}

	// print the outputTable
	output_table.PrintTable(metricsStorageRetention)

	return nil
}

// MetricsStorageRetentionCmd represents the metricsStorageRetention command
var MetricsStorageRetentionCmd = &cobra.Command{
	Use:   "metricsRetention",
	Short: "Get metric storage retention time.",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		// generate an url
		url := config.GetBaseUrl() + "metrics-storage-retentions"

		// get output flag
		outputType := config.GetOutputType()

		// call the command
		body, err := runCommand(url, "metrics storage retentions", outputType)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		// print outputTable output
		if body != nil && (outputType == "" || outputType == "wide") {
			if err := printMetricsStorageRetentionTable(body); err != nil {
				cmd.SilenceUsage = true
				return err
			}
		}

		return nil
	},
}
