package create

/*
 * Create an alert config receiver.
 */

import (
	"errors"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/internal/config"
)

// ReceiversCmd represents the receivers command
var ReceiversCmd = &cobra.Command{
	Use:   "receiver",
	Short: "Create alert config receiver.",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		if config.GetBodyFile() == "" {
			return errors.New("required flag \"--file(-f)\" not set")
		}

		// generate an url
		url := config.GetBaseUrl() + "alertconfigs/receivers"

		// call command
		if err := runCommand(url, "alert config receiver", "", nil); err != nil {
			cmd.SilenceUsage = true
			return err
		}

		return nil
	},
}
