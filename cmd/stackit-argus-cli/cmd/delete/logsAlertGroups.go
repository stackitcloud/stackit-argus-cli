package delete

/*
 * Delete a logs alert group.
 */

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
)

// LogsAlertGroupsCmd represents the logsAlertGroups command
var LogsAlertGroupsCmd = &cobra.Command{
	Use:   "logsAlertGroup <groupName>",
	Short: "Delete logs alert group config.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// generate an url
		url := config.GetBaseUrl() + fmt.Sprintf("logs-alertgroups/%s", args[0])

		// call command
		err := runCommand(url, "logs alert group")
		cobra.CheckErr(err)
	},
}
