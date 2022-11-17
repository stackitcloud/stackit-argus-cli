package create

/*
 * Create an alert record.
 */

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
)

// AlertRecordsCmd represents the alertRecords command
var AlertRecordsCmd = &cobra.Command{
	Use:   "alertRecord <groupName>",
	Short: "Create an alert record.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// generate an url
		url := config.GetBaseUrl() + fmt.Sprintf("alertgroups/%s/records", args[0])

		// call command
		err := runCommand(url, "alert record", nil)
		cobra.CheckErr(err)
	},
}
