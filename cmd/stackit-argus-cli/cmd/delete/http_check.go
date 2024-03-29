package delete

/*
 * Delete http check.
 */

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/internal/config"
)

// HttpCheckCmd represents the HttpCheck command
var HttpCheckCmd = &cobra.Command{
	Use:   "http-check <url>",
	Short: "Delete a http check.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		// generate an url
		url := config.GetBaseUrl() + fmt.Sprintf("http-checks/%s", args[0])

		// call command
		if err := runCommand(url, "http check"); err != nil {
			cmd.SilenceUsage = true
			return err
		}

		return nil
	},
}
