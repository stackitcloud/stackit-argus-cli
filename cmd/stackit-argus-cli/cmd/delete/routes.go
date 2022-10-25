package delete

/*
 * Delete a route.
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
	Short: "Delete an alert config route.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// generate an url
		url := config.GetBaseUrl() + fmt.Sprintf("alertconfigs/routes/%s", args[0])

		// print url if debug mode is turned on
		if config.IsDebugMode() {
			fmt.Println("delete alert config route command called")
			fmt.Printf("url to call - %s\n", url)
		}

		// delete the route
		status := deleteRequest(url)

		// print response status
		utils.ResponseMessage(status, "alert config route", "delete")
	},
}
