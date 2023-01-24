package create

/*
 * Create a logs alert group.
 */

import (
	"errors"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/internal/config"
)

// LogsAlertGroupsCmd represents the logsAlertGroups command
var LogsAlertGroupsCmd = &cobra.Command{
	Use:   "logs-alert-group",
	Short: "Create a logs alert group config.",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		if config.GetBodyFile() == "" {
			return errors.New("required flag \"--file(-f)\" not set")
		}

		// generate an url
		url := config.GetBaseUrl() + "logs-alertgroups"

		// call command
		if err := runCommand(url, "logs alert groups", "", nil); err != nil {
			cmd.SilenceUsage = true
			return err
		}

		return nil
	},
}
