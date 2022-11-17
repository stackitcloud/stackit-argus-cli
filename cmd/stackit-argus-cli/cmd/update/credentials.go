package update

/*
 * Update credentials.
 */

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
)

// CredentialsCmd represents the credentials command
var CredentialsCmd = &cobra.Command{
	Use:   "credentials <username>",
	Short: "Update remote write config for credentials.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// generate an url
		url := config.GetBaseUrl() + fmt.Sprintf("credentilads/%s/remote-write-limits", args[0])

		// call command
		err := runCommand(url, "credentials", "PUT")
		cobra.CheckErr(err)
	},
}
