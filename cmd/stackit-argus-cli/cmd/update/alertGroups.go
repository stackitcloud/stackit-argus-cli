package update

/*
 * Update alert groups.
 */

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/utils"
)

// AlertGroupsCmd represents the alertGroups command
var AlertGroupsCmd = &cobra.Command{
	Use:   "alertGroups <groupName>",
	Short: "Update alert groups.",
	Long:  "Patch alert groups if group name was not specified, otherwise update alert group config.",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var method string
		var debugMessage string

		// generate an url
		url := config.GetBaseUrl() + "alertgroups"

		// modify url and debug msg if an argument has been given
		if len(args) == 0 {
			debugMessage = "patch alert groups command called"
			method = "PATCH"
		} else if len(args) == 1 {
			debugMessage = "update alert group config command called"
			method = "PUT"
			url += fmt.Sprintf("/%s", args[0])
		}

		// print debug messages if debug mode is turned on
		if config.IsDebugMode() {
			fmt.Println(debugMessage)
			fmt.Printf("url to call - %s\n", url)
		}

		// update the alert group
		status := updateRequest(url, method)

		// print response status
		utils.ResponseMessage(status, "alert group", "update")
	},
}
