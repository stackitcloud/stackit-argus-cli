package update

/*
 * Update an alert config.
 */

import (
	"fmt"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/utils"

	"github.com/spf13/cobra"
)

// AlertConfigsCmd represents the alertConfigs command
var AlertConfigsCmd = &cobra.Command{
	Use:   "alertConfig",
	Short: "Update an alert config.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// generate an url
		url := config.GetBaseUrl() + "alertconfigs"

		// print debug messages if debug mode is turned on
		if config.IsDebugMode() {
			fmt.Println("update alert config command called")
			fmt.Printf("url to call - %s\n", url)
		}

		// update the alert config
		status := updateRequest(url, "PUT")

		// print response status
		utils.ResponseMessage(status, "alert config", "update")
	},
}
