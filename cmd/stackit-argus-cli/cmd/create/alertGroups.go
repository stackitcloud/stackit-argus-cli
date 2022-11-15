package create

/*
 * Create an alert group.
 */

import (
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
)

// AlertGroupsCmd represents the alertGroups command
var AlertGroupsCmd = &cobra.Command{
	Use:   "alertGroup",
	Short: "Create an alert group.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// generate an url
		url := config.GetBaseUrl() + "alertgroups"

		// call command
		runCommand(url, "alert group", nil)
	},
}
