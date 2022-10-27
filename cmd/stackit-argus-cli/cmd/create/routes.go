package create

/*
 * Create an alert config route.
 */

import (
	"fmt"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/utils"

	"github.com/spf13/cobra"
)

// RoutesCmd represents the routes command
var RoutesCmd = &cobra.Command{
	Use:   "route",
	Short: "Create alert config route in routes of route.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// generate an url
		url := config.GetBaseUrl() + "alertconfigs/routes"

		// print debug messages if debug mode is turned on
		if config.IsDebugMode() {
			fmt.Println("create alert config route command called")
			fmt.Printf("url to call - %s\n", url)
		}

		// create the route
		status := postRequest(url, nil)

		// print response status
		utils.ResponseMessage(status, "route", "create")
	},
}
