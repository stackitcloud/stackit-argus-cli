package get

/*
 * Get an acl.
 */

import (
	"fmt"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/utils"

	"github.com/spf13/cobra"
)

// AclCmd represents the acl command
var AclCmd = &cobra.Command{
	Use:   "acl",
	Short: "Get acl for the instance.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// generate an url
		url := config.GetBaseUrl() + "acl"

		// print debug messages if debug mode is turned on
		if config.IsDebugMode() {
			fmt.Println("get acl command called")
			fmt.Printf("url to call - %s\n", url)
		}

		// get acl
		status, body := getRequest(url)

		// print response status
		utils.ResponseMessage(status, "acl", "get")

		// print response body
		if status == 200 {
			fmt.Print(body)
		}
	},
}
