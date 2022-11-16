package delete

/*
 * Delete a route.
 */

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
)

// RoutesCmd represents the routes command
var RoutesCmd = &cobra.Command{
	Use:   "route <receiver>",
	Short: "Delete an alert config route.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// generate an url
		url := config.GetBaseUrl() + fmt.Sprintf("alertconfigs/routes/%s", args[0])

		// call command
		err := runCommand(url, "alert config route")
		cobra.CheckErr(err)
	},
}
