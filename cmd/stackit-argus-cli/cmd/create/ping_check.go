package create

/*
 * Create ping check.
 */

import (
	"errors"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/config"
)

// PingCheckCmd represents the PingCheck command
var PingCheckCmd = &cobra.Command{
	Use:   "ping-check",
	Short: "Create a ping check.",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		if config.GetBodyFile() == "" {
			return errors.New("required flag \"--file(-f)\" not set")
		}

		// generate an url
		url := config.GetBaseUrl() + "ping-checks"

		// call command
		if err := runCommand(url, "ping check", "", nil); err != nil {
			cmd.SilenceUsage = true
			return err
		}

		return nil
	},
}
