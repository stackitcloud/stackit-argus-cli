package update

/*
 * Update a grafana config.
 */

import (
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
)

// GrafanaConfigsCmd represents the grafanaConfigs command
var GrafanaConfigsCmd = &cobra.Command{
	Use:   "grafanaConfig",
	Short: "Update grafana config.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// generate an url
		url := config.GetBaseUrl() + "grafana-configs"

		// call command
		err := runCommand(url, "grafana config", "PUT")
		cobra.CheckErr(err)
	},
}
