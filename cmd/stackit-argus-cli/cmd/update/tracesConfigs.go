package update

/*
 * Update a traces config.
 */

import (
	"fmt"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/utils"

	"github.com/spf13/cobra"
)

// TracesConfigsCmd represents the tracesConfigs command
var TracesConfigsCmd = &cobra.Command{
	Use:   "tracesConfig",
	Short: "Update a traces config.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// generate an url
		url := config.GetBaseUrl() + "traces-configs"

		// print debug messages if debug mode is turned on
		if config.IsDebugMode() {
			fmt.Println("update traces config command called")
			fmt.Printf("url to call - %s\n", url)
		}

		// update the traces configs
		status := updateRequest(url, "PUT")

		// print response status
		utils.ResponseMessage(status, "traces config", "update")
	},
}
