package delete

/*
 * Delete a route.
 */

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/config"
)

// RoutesCmd represents the routes command
var RoutesCmd = &cobra.Command{
	Use:   "route <receiver>",
	Short: "Delete an alert config route.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		// generate an url
		url := config.GetBaseUrl() + fmt.Sprintf("alertconfigs/routes/%s", args[0])

		// call command
		if err := runCommand(url, "alert config route"); err != nil {
			cmd.SilenceUsage = true

			return err
		}

		return nil
	},
}
