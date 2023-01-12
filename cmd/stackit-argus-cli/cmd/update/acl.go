package update

/*
 * Update an acl.
 */

import (
	"errors"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/config"
)

// AclCmd represents the acl command
var AclCmd = &cobra.Command{
	Use:   "acl",
	Short: "Update an acl config.",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		if config.GetBodyFile() == "" {
			return errors.New("required flag \"--file(-f)\" not set")
		}

		// generate an url
		url := config.GetBaseUrl() + "acl"

		// call command
		if err := runCommand(url, "acl", "PUT"); err != nil {
			cmd.SilenceUsage = true

			return err
		}

		return nil
	},
}
