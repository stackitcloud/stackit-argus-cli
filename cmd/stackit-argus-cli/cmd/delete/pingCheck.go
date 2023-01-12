package delete

/*
 * Delete ping check.
 */

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
)

// PingCheckCmd represents the PingCheck command
var PingCheckCmd = &cobra.Command{
	Use:   "pingCheck",
	Short: "Delete a ping check.",
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
