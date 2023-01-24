package delete

/*
 * Delete a receiver.
 */

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/internal/config"
)

// ReceiversCmd represents the receivers command
var ReceiversCmd = &cobra.Command{
	Use:   "receiver <receiver>",
	Short: "Delete alert config receiver.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		// generate an url
		url := config.GetBaseUrl() + fmt.Sprintf("alertconfigs/receivers/%s", args[0])

		// call command
		if err := runCommand(url, "receiver"); err != nil {
			cmd.SilenceUsage = true
			return err
		}

		return nil
	},
}
