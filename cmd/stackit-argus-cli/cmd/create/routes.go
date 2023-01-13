package create

/*
 * Create an alert config route.
 */

import (
	"errors"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/config"
)

// RoutesCmd represents the routes command
var RoutesCmd = &cobra.Command{
	Use:   "route",
	Short: "Create alert config route in routes of route.",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		if config.GetBodyFile() == "" {
			return errors.New("required flag \"--file(-f)\" not set")
		}

		// generate an url
		url := config.GetBaseUrl() + "alertconfigs/routes"

		// call command
		if err := runCommand(url, "alert config route", "", nil); err != nil {
			cmd.SilenceUsage = true
			return err
		}

		return nil
	},
}
