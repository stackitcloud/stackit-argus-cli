package delete

/*
 * Delete an instance.
 */

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/internal/config"
)

// InstanceCmd represents the instance command
var InstanceCmd = &cobra.Command{
	Use:   "instance <instance-id>",
	Short: "Delete an instance.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		// generate an url
		url := config.GetInstancesUrl() + fmt.Sprintf("/%s", args[0])

		// call command
		if err := runCommand(url, "instance"); err != nil {
			cmd.SilenceUsage = true
			return err
		}

		return nil
	},
}
