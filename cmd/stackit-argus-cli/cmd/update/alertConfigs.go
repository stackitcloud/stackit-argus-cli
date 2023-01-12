package update

/*
 * Update an alert config.
 */

import (
	"errors"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/config"
)

// AlertConfigsCmd represents the alertConfigs command
var AlertConfigsCmd = &cobra.Command{
	Use:   "alertConfig",
	Short: "Update an alert config.",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		if config.GetBodyFile() == "" {
			return errors.New("required flag \"--file(-f)\" not set")
		}

		// generate an url
		url := config.GetBaseUrl() + "alertconfigs"

		// call command
		if err := runCommand(url, "alert config", "PUT"); err != nil {
			cmd.SilenceUsage = true

			return err
		}

		return nil
	},
}
