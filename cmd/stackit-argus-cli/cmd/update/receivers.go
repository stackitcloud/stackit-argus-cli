package update

/*
 * Update a receiver.
 */

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
)

// ReceiversCmd represents the receivers command
var ReceiversCmd = &cobra.Command{
	Use:   "receiver <receiver>",
	Short: "Update an alert config receiver.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// generate an url
		url := config.GetBaseUrl() + fmt.Sprintf("alertconfigs/receivers/%s", args[0])

		// call command
		err := runCommand(url, "alert config receiver", "PUT")
		cobra.CheckErr(err)
	},
}
