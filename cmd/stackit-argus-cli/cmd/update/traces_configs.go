package update

/*
 * Update a traces config.
 */

import (
	"errors"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/internal/config"
)

// TracesConfigsCmd represents the tracesConfigs command
var TracesConfigsCmd = &cobra.Command{
	Use:   "traces-config",
	Short: "Update a traces config.",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		if config.GetBodyFile() == "" {
			return errors.New("required flag \"--file(-f)\" not set")
		}

		// generate an url
		url := config.GetBaseUrl() + "traces-configs"

		// call command
		if err := runCommand(url, "traces config", "PUT"); err != nil {
			cmd.SilenceUsage = true
			return err
		}

		return nil
	},
}
