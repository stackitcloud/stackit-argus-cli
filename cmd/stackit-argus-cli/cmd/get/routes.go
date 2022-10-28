package get

/*
 * Get alert configs routes.
 */

import (
	"fmt"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/utils"

	"github.com/spf13/cobra"
)

// RoutesCmd represents the routes command
var RoutesCmd = &cobra.Command{
	Use:   "routes <receiver>",
	Short: "Get alert config routes.",
	Long:  "Get list of alert config route if receiver was not specified, otherwise get alert receiver for route.",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var debugMsg string

		// generate an url
		url := config.GetBaseUrl() + "alertconfigs/routes"

		// modify url and debug message depend on arguments
		if len(args) == 1 {
			debugMsg = "get alert config route command called"
			url += fmt.Sprintf("/%s", args[0])
		} else if len(args) == 0 {
			debugMsg = "list alert config routes command called"
		}

		// print debug messages if debug mode is turned on
		if config.IsDebugMode() {
			fmt.Println(debugMsg)
			fmt.Printf("url to call - %s\n", url)
		}

		// get alert config routes
		status, body := getRequest(url)

		// print response status
		utils.ResponseMessage(status, "alert config routes", "get")

		// print response body
		if status == 200 {
			fmt.Println(body)
		}
	},
}
