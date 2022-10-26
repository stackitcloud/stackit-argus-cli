package get

/*
 * Get backup schedules.
 */

import (
	"fmt"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/utils"

	"github.com/spf13/cobra"
)

// BackupSchedulesCmd represents the backupSchedules command
var BackupSchedulesCmd = &cobra.Command{
	Use:   "backupSchedules",
	Short: "Get backup schedules.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// generate an url
		url := config.GetBaseUrl() + "backup-schedules"

		// print debug messages if debug mode is turned on
		if config.IsDebugMode() {
			fmt.Println("get backup schedules command called")
			fmt.Printf("url to call - %s\n", url)
		}

		// get backup schedules
		status, body := getRequest(url)

		// print response status
		utils.ResponseMessage(status, "backup schedules", "get")

		// print response body
		if status == 200 {
			fmt.Print(body)
		}
	},
}
