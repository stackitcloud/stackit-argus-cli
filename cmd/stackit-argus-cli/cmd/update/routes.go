package update

/*
 * Update an alert receiver for route.
 */

import (
	"fmt"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/utils"

	"github.com/spf13/cobra"
)

// RoutesCmd represents the routes command
var RoutesCmd = &cobra.Command{
	Use:   "route <receiver>",
	Short: "Update alert receiver for route.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// generate an url
		url := config.GetBaseUrl() + fmt.Sprintf("/alertconfigs/routes/%s", args[0])

		// print debug messages if debug mode is turned on
		if config.IsDebugMode() {
			fmt.Println("update route command called")
			fmt.Printf("url to call - %s\n", url)
		}

		// update the route
		status := updateRequest(url, "PUT")

		// print response status
		utils.ResponseMessage(status, "route", "update")
	},
}
