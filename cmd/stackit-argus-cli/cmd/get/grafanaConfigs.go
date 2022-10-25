package get

/*
 * Get grafana configs.
 */

import (
	"fmt"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/utils"

	"github.com/spf13/cobra"
)

// GrafanaConfigsCmd represents the grafanaConfigs command
var GrafanaConfigsCmd = &cobra.Command{
	Use:   "grafanaConfigs",
	Short: "Get grafana config.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// generate an url
		url := config.GetBaseUrl() + "grafana-configs"

		// print debug messages if debug mode is turned on
		if config.IsDebugMode() {
			fmt.Println("get grafana configs command called")
			fmt.Printf("url to call - %s\n", url)
		}

		// get grafana configs
		status, body := getRequest(url)

		// print response status
		utils.ResponseMessage(status, "grafana configs", "get")

		// print response body
		if status == 200 {
			fmt.Print(body)
		}
	},
}
