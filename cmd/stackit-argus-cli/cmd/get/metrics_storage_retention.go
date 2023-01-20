package get

/*
 * Get metrics storage retention.
 */

import (
	"encoding/json"
	"github.com/spf13/cobra"
	config2 "github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/output_table"
)

// metricsStorageRetention is used to unmarshal metrics storage retention response body
type metricsStorageRetention struct {
	MetricsRetentionTimeRaw string `json:"metricsRetentionTimeRaw" header:"metrics retention time raw"`
	MetricsRetentionTime5M  string `json:"metricsRetentionTime5m" header:"metrics retention time 5m"`
	MetricsRetentionTime1H  string `json:"metricsRetentionTime1h" header:"metrics retention time 1h"`
}

// printMetricsStorageRetentionTable prints routes response body as output_table
func printMetricsStorageRetentionTable(body []byte) error {
	var metricsStorageRetention metricsStorageRetention

	// unmarshal response body
	if err := json.Unmarshal(body, &metricsStorageRetention); err != nil {
		return err
	}

	// print the output_table
	output_table.PrintTable(metricsStorageRetention)

	return nil
}

// MetricsStorageRetentionCmd represents the metricsStorageRetention command
var MetricsStorageRetentionCmd = &cobra.Command{
	Use:   "metrics-storage-retention",
	Short: "Get metric storage retention time.",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		// generate an url
		url := config2.GetBaseUrl() + "metrics-storage-retentions"

		// get output flag
		outputType := config2.GetOutputType()

		// call the command
		body, err := runCommand(url, "metrics storage retentions", outputType)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		// print output_table output
		if body != nil && (outputType == "" || outputType == "wide") {
			if err := printMetricsStorageRetentionTable(body); err != nil {
				cmd.SilenceUsage = true
				return err
			}
		}

		return nil
	},
}
