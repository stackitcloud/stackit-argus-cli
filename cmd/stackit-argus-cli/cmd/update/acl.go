package update

/*
 * Update an acl.
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
	Short: "Update an acl config.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// generate an url
		url := config.GetBaseUrl() + "acl"

		// print debug messages if debug mode is turned on
		if config.IsDebugMode() {
			fmt.Println("update acl config command called")
			fmt.Printf("url to call - %s\n", url)
		}

		// update the acl
		status := updateRequest(url, "PUT")

		// print response status
		utils.ResponseMessage(status, "acl", "update")
	},
}
