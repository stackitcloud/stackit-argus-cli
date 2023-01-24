package create

/*
 * Create a credentials.
 */

import (
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/internal/config"
)

// CredentialsCmd represents the credentials command
var CredentialsCmd = &cobra.Command{
	Use:   "credentials",
	Short: "Create technical user credentials.",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		// generate an url
		url := config.GetBaseUrl() + "credentials"

		// call command
		if err := runCommand(url, "credentials", "", nil); err != nil {
			cmd.SilenceUsage = true
			return err
		}

		return nil
	},
}
