package delete

/*
 * Delete credentials.
 */

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/utils"
)

var remoteWriteLimits bool

// CredentialsCmd represents the credentials command
var CredentialsCmd = &cobra.Command{
	Use:   "credentials <username>",
	Short: "Delete credentials.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var debugMsg string

		// generate an url
		url := config.GetBaseUrl() + fmt.Sprintf("credentials/%s", args[0])

		// check if remote write limits flag is set, if so modify url
		if remoteWriteLimits == true {
			debugMsg = "delete remote write config for credentials command called"
			url += "/remote-write-limits"
		} else {
			debugMsg = "delete technical credentials command called"
		}

		// print url if debug mode is turned on
		if config.IsDebugMode() {
			fmt.Println(debugMsg)
			fmt.Printf("url to call - %s\n", url)
		}

		// delete the credentials
		status := deleteRequest(url)

		// print response status
		utils.ResponseMessage(status, "credentials", "delete")
	},
}

// init flags for the command
func init() {
	CredentialsCmd.Flags().BoolVarP(&remoteWriteLimits, "remote-write-limits", "r", false, "delete remote write config for credentials")
}
