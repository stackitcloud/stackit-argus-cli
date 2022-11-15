package create

/*
 * Create ping check.
 */

import (
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
)

// PingCheckCmd represents the PingCheck command
var PingCheckCmd = &cobra.Command{
	Use:   "pingCheck",
	Short: "Create a ping check.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// generate an url
		url := config.GetBaseUrl() + "ping-checks"

		// call command
		runCommand(url, "ping check", nil)
	},
}
