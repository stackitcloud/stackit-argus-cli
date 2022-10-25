package get

/*
 * Get backup retentions.
 */

import (
	"fmt"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/utils"

	"github.com/spf13/cobra"
)

// BackupRetentionsCmd represents the backupRetentions command
var BackupRetentionsCmd = &cobra.Command{
	Use:   "backupRetentions",
	Short: "Get backup retentions.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// generate an url
		url := config.GetBaseUrl() + "backup-retentions"

		// print debug messages if debug mode is turned on
		if config.IsDebugMode() {
			fmt.Println("get backup retentions command called")
			fmt.Printf("url to call - %s\n", url)
		}

		// get backup retentions
		status, body := getRequest(url)

		// print response status
		utils.ResponseMessage(status, "backup retentions", "get")

		// print response body
		if status == 200 {
			fmt.Print(body)
		}
	},
}
