package update

/*
 * Update an alert receiver for route.
 */

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
)

// RoutesCmd represents the routes command
var RoutesCmd = &cobra.Command{
	Use:   "route <receiver>",
	Short: "Update alert receiver for route.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// generate an url
		url := config.GetBaseUrl() + fmt.Sprintf("alertconfigs/routes/%s", args[0])

		// call command
		runCommand(url, "route", "PUT")
	},
}
