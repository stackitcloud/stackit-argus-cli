package create

/*
 * Create cert check.
 */

import (
	"fmt"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/utils"

	"github.com/spf13/cobra"
)

// CertCheckCmd represents the CertCheck command
var CertCheckCmd = &cobra.Command{
	Use:   "certCheck",
	Short: "Create a cert check.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// generate an url
		url := config.GetBaseUrl() + "cert-checks"

		// print debug messages if debug mode is turned on
		if config.IsDebugMode() {
			fmt.Println("create cert check command called")
			fmt.Printf("url to call - %s\n", url)
		}

		// create the cert check
		status := postRequest(url, nil)

		// print response status
		utils.ResponseMessage(status, "cert check", "create")
	},
}
