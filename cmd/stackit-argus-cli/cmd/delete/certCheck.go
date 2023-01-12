package delete

/*
 * Delete cert check.
 */

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
)

// CertCheckCmd represents the CertCheck command
var CertCheckCmd = &cobra.Command{
	Use:   "certCheck",
	Short: "Delete a cert check.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		// generate an url
		url := config.GetBaseUrl() + fmt.Sprintf("cert-checks/%s", args[0])

		// call command
		if err := runCommand(url, "cert check"); err != nil {
			cmd.SilenceUsage = true

			return err
		}

		return nil
	},
}
