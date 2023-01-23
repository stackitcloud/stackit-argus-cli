package get

/*
 * Get logs alert groups.
 */

import (
	"fmt"
	"github.com/spf13/cobra"
	config2 "github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/config"
)

// LogsAlertGroupsCmd represents the logsAlertGroups command
var LogsAlertGroupsCmd = &cobra.Command{
	Use:   "logs-alert-groups <group-name>",
	Short: "Get logs alert groups config.",
	Long:  "Get list of logs alert groups config if group name was not specified, otherwise get logs alert group config.",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		// generate an url
		url := config2.GetBaseUrl() + "logs-alertgroups"

		// modify url and debug message depend on arguments
		resource := "logs alert groups"
		if len(args) == 1 {
			resource = "logs alert group"
			url += fmt.Sprintf("/%s", args[0])
		}

		// get output flag
		outputType := config2.GetOutputType()

		// call the command
		body, err := runCommand(url, resource, outputType)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		if body != nil {
			if len(args) == 0 {
				if err := printAlertGroupsListTable(body, outputType); err != nil {
					cmd.SilenceUsage = true
					return err
				}
			} else {
				if err := printAlertGroupTable(body, outputType); err != nil {
					cmd.SilenceUsage = true
					return err
				}
			}
		}

		return err
	},
}
