package update

/*
 * Update a logs config.
 */

import (
	"fmt"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/utils"

	"github.com/spf13/cobra"
)

// LogsConfigsCmd represents the logsConfigs command
var LogsConfigsCmd = &cobra.Command{
	Use:   "logsConfigs",
	Short: "Update a logs config.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// generate an url
		url := config.GetBaseUrl() + "logs-configs"

		// print debug messages if debug mode is turned on
		if config.IsDebugMode() {
			fmt.Println("update logs config command called")
			fmt.Printf("url to call - %s\n", url)
		}

		// update the logs configs
		status := updateRequest(url, "PUT")

		// print response status
		utils.ResponseMessage(status, "logs configs", "update")
	},
}
