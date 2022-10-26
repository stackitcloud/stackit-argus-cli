package get

/*
 * Get credentials.
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
	Short: "Get technical user credentials.",
	Long:  "Get list of all credentials if username was not specified, otherwise get technical credentials.",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var debugMsg string

		// generate an url
		url := config.GetBaseUrl() + "credentials"

		// modify url and debug message depend on argument and flag
		if remoteWriteLimits == true {
			if len(args) == 1 {
				debugMsg = "get remote write limits credentials command called"
				url += fmt.Sprintf("/%s/remote-write-limits", args[0])
			} else {
				cobra.CheckErr("get credentials with remote write limits must have one argument")
			}
		} else {
			if len(args) == 1 {
				debugMsg = "get credential command called"
				url += fmt.Sprintf("/%s", args[0])
			} else if len(args) == 0 {
				debugMsg = "list credentials command called"
			}
		}

		// print debug messages if debug mode is turned on
		if config.IsDebugMode() {
			fmt.Println(debugMsg)
			fmt.Printf("url to call - %s\n", url)
		}

		// get credentials
		status, body := getRequest(url)

		// print response status
		utils.ResponseMessage(status, "credentials", "get")

		// print response body
		if status == 200 {
			fmt.Println(body)
		}
	},
}

// init flags for the command
func init() {
	CredentialsCmd.Flags().BoolVarP(&remoteWriteLimits, "remote-write-limits", "r", false, "get remote write limits credentials")
}
