package create

/*
 * Create network check.
 */

import (
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
)

// NetworkCheckCmd represents the NetworkCheck command
var NetworkCheckCmd = &cobra.Command{
	Use:   "networkCheck",
	Short: "Create a network check.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// generate an url
		url := config.GetBaseUrl() + "network-checks"

		// call command
		err := runCommand(url, "network check", nil)
		cobra.CheckErr(err)
	},
}
