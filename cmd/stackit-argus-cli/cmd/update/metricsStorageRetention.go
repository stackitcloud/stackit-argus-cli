package update

/*
 * Update a metrics storage retention.
 */

import (
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
)

// MetricsStorageRetentionCmd represents the metricsStorageRetention command
var MetricsStorageRetentionCmd = &cobra.Command{
	Use:   "metricsStorageRetention",
	Short: "Update metric storage retention time.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// generate an url
		url := config.GetBaseUrl() + "metrics-storage-retentions"

		// call command
		err := runCommand(url, "metrics storage retention", "PUT")
		cobra.CheckErr(err)
	},
}
