package update

/*
 * Update a metrics storage retention.
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
	Short: "Update metric storage retention time.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// generate an url
		url := config.GetBaseUrl() + "metrics-storage-retentions"

		// print debug messages if debug mode is turned on
		if config.IsDebugMode() {
			fmt.Println("update metric storage retention command called")
			fmt.Printf("url to call - %s\n", url)
		}

		// update the metrics storage retention
		status := updateRequest(url, "PUT")

		// print response status
		utils.ResponseMessage(status, "metrics storage retention", "update")
	},
}
