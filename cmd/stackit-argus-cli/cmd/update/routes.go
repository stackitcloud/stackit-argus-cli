package update

/*
 * Update an alert receiver for route.
 */

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/internal/config"
)

// RoutesCmd represents the routes command
var RoutesCmd = &cobra.Command{
	Use:   "route <receiver>",
	Short: "Update alert receiver for route.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if config.GetBodyFile() == "" {
			return errors.New("required flag \"--file(-f)\" not set")
		}

		// generate an url
		url := config.GetBaseUrl() + fmt.Sprintf("alertconfigs/routes/%s", args[0])

		// call command
		if err := runCommand(url, "alert config route", "PUT"); err != nil {
			cmd.SilenceUsage = true
			return err
		}

		return nil
	},
}
