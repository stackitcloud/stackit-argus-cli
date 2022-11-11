package update

/*
 * Update alert records.
 */

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
)

// AlertRecordsCmd represents the alertRecords command
var AlertRecordsCmd = &cobra.Command{
	Use:   "alertRecords <groupName> <alertRecord>",
	Short: "Update alert records.",
	Long:  "Patch alert records if alert record was not specified, otherwise update alert record.",
	Args:  cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		// generate an url
		url := config.GetBaseUrl() + fmt.Sprintf("alertgroups/%s/records", args[0])

		// modify url and debug msg if an argument has been given
		method := "PATCH"
		resource := "alert records"
		if len(args) == 2 {
			resource = "alert record"
			method = "PUT"
			url += fmt.Sprintf("/%s", args[1])
		}

		// call command
		runCommand(url, resource, method)
	},
}
