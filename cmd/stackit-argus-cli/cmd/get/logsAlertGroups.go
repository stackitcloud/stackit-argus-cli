package get

/*
 * Get logs alert groups.
 */

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
)

// LogsAlertGroupsCmd represents the logsAlertGroups command
var LogsAlertGroupsCmd = &cobra.Command{
	Use:   "logsAlertGroups <groupName>",
	Short: "Get logs alert groups config.",
	Long:  "Get list of logs alert groups config if group name was not specified, otherwise get logs alert group config.",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// generate an url
		url := config.GetBaseUrl() + "logs-alertgroups"

		// modify url and debug message depend on arguments
		resource := "logs alert groups"
		if len(args) == 1 {
			resource = "logs alert group"
			url += fmt.Sprintf("/%s", args[0])
		}

		// get output flag
		outputType := config.GetOutputType()

		// call the command
		body, err := runCommand(url, resource, outputType)
		cobra.CheckErr(err)

		if body != nil {
			if len(args) == 0 {
				printAlertGroupsListTable(body)
			} else {
				printAlertGroupTable(body)
			}
		}
	},
}
