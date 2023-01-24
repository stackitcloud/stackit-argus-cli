package create

/*
 * Create cert check.
 */

import (
	"errors"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/internal/config"
)

// CertCheckCmd represents the CertCheck command
var CertCheckCmd = &cobra.Command{
	Use:   "cert-check",
	Short: "Create a cert check.",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		if config.GetBodyFile() == "" {
			return errors.New("required flag \"--file(-f)\" not set")
		}

		// generate an url
		url := config.GetBaseUrl() + "cert-checks"

		// call command
		if err := runCommand(url, "cert check", "", nil); err != nil {
			cmd.SilenceUsage = true
			return err
		}

		return nil
	},
}
