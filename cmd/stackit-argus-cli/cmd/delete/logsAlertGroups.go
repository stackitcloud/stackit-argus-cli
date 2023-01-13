package delete

/*
 * Delete a logs alert group.
 */

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/config"
)

// LogsAlertGroupsCmd represents the logsAlertGroups command
var LogsAlertGroupsCmd = &cobra.Command{
	Use:   "logsAlertGroup <groupName>",
	Short: "Delete logs alert group config.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		// generate an url
		url := config.GetBaseUrl() + fmt.Sprintf("logs-alertgroups/%s", args[0])

		// call command
		if err := runCommand(url, "logs alert group"); err != nil {
			cmd.SilenceUsage = true
			return err
		}

		return nil
	},
}
