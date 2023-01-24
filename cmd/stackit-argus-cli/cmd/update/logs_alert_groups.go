package update

/*
 * Update a logs alert group config.
 */

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/internal/config"
)

// LogsAlertGroupsCmd represents the logsAlertGroups command
var LogsAlertGroupsCmd = &cobra.Command{
	Use:   "logs-alert-group <group-name>",
	Short: "Update logs alert group config.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if config.GetBodyFile() == "" {
			return errors.New("required flag \"--file(-f)\" not set")
		}

		// generate an url
		url := config.GetBaseUrl() + fmt.Sprintf("logs-alertgroups/%s", args[0])

		// call command
		if err := runCommand(url, "logs alert group", "PUT"); err != nil {
			cmd.SilenceUsage = true
			return err
		}

		return nil
	},
}
