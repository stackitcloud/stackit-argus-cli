package delete

/*
 * Delete ping check.
 */

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/config"
)

// PingCheckCmd represents the PingCheck command
var PingCheckCmd = &cobra.Command{
	Use:   "pingCheck <domain>",
	Short: "Delete a ping check.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		// generate an url
		url := config.GetBaseUrl() + fmt.Sprintf("ping-checks/%s", args[0])

		// call command
		if err := runCommand(url, "ping check"); err != nil {
			cmd.SilenceUsage = true

			return err
		}

		return nil
	},
}
