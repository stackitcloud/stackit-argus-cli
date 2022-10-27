package update

/*
 * Update a grafana config.
 */

import (
	"fmt"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/utils"

	"github.com/spf13/cobra"
)

// GrafanaConfigsCmd represents the grafanaConfigs command
var GrafanaConfigsCmd = &cobra.Command{
	Use:   "grafanaConfig",
	Short: "Update grafana config.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// generate an url
		url := config.GetBaseUrl() + "grafana-configs"

		// print debug messages if debug mode is turned on
		if config.IsDebugMode() {
			fmt.Println("update grafana config command called")
			fmt.Printf("url to call - %s\n", url)
		}

		// update the grafana config
		status := updateRequest(url, "PUT")

		// print response status
		utils.ResponseMessage(status, "grafana config", "update")
	},
}
