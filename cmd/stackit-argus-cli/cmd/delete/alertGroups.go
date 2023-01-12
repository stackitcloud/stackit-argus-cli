package delete

/*
 * Delete an alert group.
 */

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/config"
)

// AlertGroupsCmd represents the alertGroups command
var AlertGroupsCmd = &cobra.Command{
	Use:   "alertGroup <groupName>",
	Short: "Delete alert group config.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		// generate an url
		url := config.GetBaseUrl() + fmt.Sprintf("alertgroups/%s", args[0])

		// call command
		if err := runCommand(url, "alert group"); err != nil {
			cmd.SilenceUsage = true

			return err
		}

		return nil
	},
}
