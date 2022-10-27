package update

/*
 * Update credentials.
 */

import (
	"fmt"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/utils"

	"github.com/spf13/cobra"
)

// CredentialsCmd represents the credentials command
var CredentialsCmd = &cobra.Command{
	Use:   "credentials <username>",
	Short: "Update remote write config for credentials.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// generate an url
		url := config.GetBaseUrl() + fmt.Sprintf("credentilads/%s/remote-write-limits", args[0])

		// print debug messages if debug mode is turned on
		if config.IsDebugMode() {
			fmt.Println("update credentials command called")
			fmt.Printf("url to call - %s\n", url)
		}

		// update the credentials
		status := updateRequest(url, "PUT")

		// print response status
		utils.ResponseMessage(status, "credentials", "update")
	},
}
