package create

/*
 * Create a credentials.
 */

import (
	"fmt"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/utils"

	"github.com/spf13/cobra"
)

// CredentialsCmd represents the credentials command
var CredentialsCmd = &cobra.Command{
	Use:   "credentials",
	Short: "Create technical user credentials.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// generate an url
		url := config.GetBaseUrl() + "credentials"

		// print debug messages if debug mode is turned on
		if config.IsDebugMode() {
			fmt.Println("create technical user credentials command called")
			fmt.Printf("url to call - %s\n", url)
		}

		// create the credentials
		status := postRequest(url, nil)

		// print response status
		utils.ResponseMessage(status, "credentials", "create")
	},
}
