package delete

/*
 * Delete credentials.
 */

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
)

var remoteWriteLimits bool

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
		if remoteWriteLimits == true {
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

// init flags for the command
func init() {
	CredentialsCmd.Flags().BoolVarP(&remoteWriteLimits, "remote-write-limits", "r", false, "delete remote write config for credentials")
}
