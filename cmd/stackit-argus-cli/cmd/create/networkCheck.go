package create

/*
 * Create network check.
 */

import (
	"fmt"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/utils"

	"github.com/spf13/cobra"
)

// NetworkCheckCmd represents the NetworkCheck command
var NetworkCheckCmd = &cobra.Command{
	Use:   "networkCheck",
	Short: "Create a network check.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// generate an url
		url := config.GetBaseUrl() + "network-checks"

		// print debug messages if debug mode is turned on
		if config.IsDebugMode() {
			fmt.Println("create network check command called")
			fmt.Printf("url to call - %s\n", url)
		}

		// create the network check
		status := postRequest(url, nil)

		// print response status
		utils.ResponseMessage(status, "network check", "create")
	},
}
