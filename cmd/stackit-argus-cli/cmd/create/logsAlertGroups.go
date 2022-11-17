package create

/*
 * Create a logs alert group.
 */

import (
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
)

// LogsAlertGroupsCmd represents the logsAlertGroups command
var LogsAlertGroupsCmd = &cobra.Command{
	Use:   "logsAlertGroup",
	Short: "Create a logs alert group config.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// generate an url
		url := config.GetBaseUrl() + "logs-alertgroups"

		// call command
		err := runCommand(url, "logs alert groups", nil)
		cobra.CheckErr(err)
	},
}
