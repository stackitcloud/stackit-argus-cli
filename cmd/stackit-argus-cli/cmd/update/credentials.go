package update

/*
 * Update credentials.
 */

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/config"
)

// CredentialsCmd represents the credentials command
var CredentialsCmd = &cobra.Command{
	Use:   "credentials <username>",
	Short: "Update remote write config for credentials.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		// generate an url
		url := config.GetBaseUrl() + fmt.Sprintf("credentials/%s/remote-write-limits", args[0])

		// call command
		if err := runCommand(url, "credentials", "PUT"); err != nil {
			cmd.SilenceUsage = true
			return err
		}

		return nil
	},
}
