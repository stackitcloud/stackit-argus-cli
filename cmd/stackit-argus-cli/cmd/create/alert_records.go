package create

/*
 * Create an alert record.
 */

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/internal/config"
)

// AlertRecordsCmd represents the alertRecords command
var AlertRecordsCmd = &cobra.Command{
	Use:   "alert-record <group-name>",
	Short: "Create an alert record.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if config.GetBodyFile() == "" {
			return errors.New("required flag \"--file(-f)\" not set")
		}

		// generate an url
		url := config.GetBaseUrl() + fmt.Sprintf("alertgroups/%s/records", args[0])

		// call command
		if err := runCommand(url, "alert record", "", nil); err != nil {
			cmd.SilenceUsage = true
			return err
		}

		return nil
	},
}
