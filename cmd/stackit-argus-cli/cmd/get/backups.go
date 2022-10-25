package get

/*
 * Get backups.
 */

import (
	"fmt"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/utils"

	"github.com/spf13/cobra"
)

// BackupCmd represents the backup command
var BackupCmd = &cobra.Command{
	Use:   "backups",
	Short: "Get backups.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// generate an url
		url := config.GetBaseUrl() + "backups"

		// print debug messages if debug mode is turned on
		if config.IsDebugMode() {
			fmt.Println("get backups command called")
			fmt.Printf("url to call - %s\n", url)
		}

		// get backups
		status, body := getRequest(url)

		// print response status
		utils.ResponseMessage(status, "backups", "get")

		// print response body
		if status == 200 {
			fmt.Print(body)
		}
	},
}
