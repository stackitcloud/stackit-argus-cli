package delete

/*
 * Delete network check.
 */

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/internal/config"
)

// NetworkCheckCmd represents the NetworkCheck command
var NetworkCheckCmd = &cobra.Command{
	Use:   "network-check <address>",
	Short: "Delete a network check.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		// generate an url
		url := config.GetBaseUrl() + fmt.Sprintf("network-checks/%s", args[0])

		// call command
		if err := runCommand(url, "network check"); err != nil {
			cmd.SilenceUsage = true
			return err
		}

		return nil
	},
}
