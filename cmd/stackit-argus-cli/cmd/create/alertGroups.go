package create

/*
 * Create an alert group.
 */

import (
	"fmt"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/utils"

	"github.com/spf13/cobra"
)

// AlertGroupsCmd represents the alertGroups command
var AlertGroupsCmd = &cobra.Command{
	Use:   "alertGroup",
	Short: "Create an alert group.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// generate an url
		url := config.GetBaseUrl() + "alertgroups"

		// print debug messages if debug mode is turned on
		if config.IsDebugMode() {
			fmt.Println("create alert group command called")
			fmt.Printf("url to call - %s\n", url)
		}

		// create the alert group
		status := postRequest(url, nil)

		// print response status
		utils.ResponseMessage(status, "alert group", "create")
	},
}
