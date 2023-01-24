package delete

/*
 * Delete credentials.
 */

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/internal/config"
)

// CredentialsCmd represents the credentials command
var CredentialsCmd = &cobra.Command{
	Use:   "credentials <username>",
	Short: "Delete credentials.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		// generate an url
		url := config.GetBaseUrl() + fmt.Sprintf("credentials/%s", args[0])

		// check if remote write limits flag is set, if so modify url
		resource := "credentials"
		if config.IsRemoteWriteLimits() {
			resource = "remote write config for credentials"
			url += "/remote-write-limits"
		}

		// call command
		if err := runCommand(url, resource); err != nil {
			cmd.SilenceUsage = true
			return err
		}

		return nil
	},
}
