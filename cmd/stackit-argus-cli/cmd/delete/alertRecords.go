package delete

/*
 * Delete an alert record.
 */

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
)

// AlertRecordsCmd represents the alertRecords command
var AlertRecordsCmd = &cobra.Command{
	Use:   "alertRecord <groupName> <alertRecord>",
	Short: "Delete alert records.",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		// generate an url
		url := config.GetBaseUrl() + fmt.Sprintf("alertgroups/%s/records/%s", args[0], args[1])

		// call command
		if err := runCommand(url, "alert record"); err != nil {
			cmd.SilenceUsage = true

			return err
		}

		return nil
	},
}
