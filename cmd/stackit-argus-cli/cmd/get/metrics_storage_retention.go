package get

/*
 * Get metrics storage retention.
 */

import (
	"encoding/json"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/output"
	"github.com/stackitcloud/stackit-argus-cli/internal/config"
)

// metricsStorageRetention is used to unmarshal metrics storage retention response body
type metricsStorageRetention struct {
	MetricsRetentionTimeRaw string `json:"metricsRetentionTimeRaw" header:"metrics retention time raw"`
	MetricsRetentionTime5M  string `json:"metricsRetentionTime5m" header:"metrics retention time 5m"`
	MetricsRetentionTime1H  string `json:"metricsRetentionTime1h" header:"metrics retention time 1h"`
}

// printMetricsStorageRetentionTable prints routes response body as output
func printMetricsStorageRetentionTable(body []byte, outputType config.OutputType) error {
	var msr metricsStorageRetention

	// unmarshal response body
	if err := json.Unmarshal(body, &msr); err != nil {
		return err
	}

	// print the output
	output.PrintTable(msr, string(outputType))

	return nil
}

// MetricsStorageRetentionCmd represents the metricsStorageRetention command
var MetricsStorageRetentionCmd = &cobra.Command{
	Use:   "metrics-storage-retention",
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

		// print output output
		if body != nil && outputType != "yaml" && outputType != "json" {
			if err := printMetricsStorageRetentionTable(body, outputType); err != nil {
				cmd.SilenceUsage = true
				return err
			}
		}

		return nil
	},
}
