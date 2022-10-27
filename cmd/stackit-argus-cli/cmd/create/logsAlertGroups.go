package create

/*
 * Create a logs alert group.
 */

import (
	"fmt"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/utils"

	"github.com/spf13/cobra"
)

// LogsAlertGroupsCmd represents the logsAlertGroups command
var LogsAlertGroupsCmd = &cobra.Command{
	Use:   "logsAlertGroup",
	Short: "Create a logs alert group config.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// generate an url
		url := config.GetBaseUrl() + "logs-alertgroups"

		// print debug messages if debug mode is turned on
		if config.IsDebugMode() {
			fmt.Println("create logs alert group config command called")
			fmt.Printf("url to call - %s\n", url)
		}

		// create the logs alertgroup
		status := postRequest(url, nil)

		// print response status
		utils.ResponseMessage(status, "logs alertgroup", "create")
	},
}
