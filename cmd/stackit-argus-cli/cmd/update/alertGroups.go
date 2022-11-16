package update

/*
 * Update alert groups.
 */

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
)

// AlertGroupsCmd represents the alertGroups command
var AlertGroupsCmd = &cobra.Command{
	Use:   "alertGroups <groupName>",
	Short: "Update alert groups.",
	Long:  "Patch alert groups if group name was not specified, otherwise update alert group config.",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// generate an url
		url := config.GetBaseUrl() + "alertgroups"

		// modify url and debug msg if an argument has been given
		method := "PATCH"
		resource := "alert groups"
		if len(args) == 1 {
			resource = "alert group"
			method = "PUT"
			url += fmt.Sprintf("/%s", args[0])
		}

		// call command
		err := runCommand(url, resource, method)
		cobra.CheckErr(err)
	},
}
