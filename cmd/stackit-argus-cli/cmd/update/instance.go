package update

/*
 * Update an instance.
 */

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/config"
)

// InstanceCmd represents the instance command
var InstanceCmd = &cobra.Command{
	Use:   "instance <instanceId>",
	Short: "Update an instance.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if config.GetBodyFile() == "" {
			return errors.New("required flag \"--file(-f)\" not set")
		}

		// generate an url
		url := config.GetInstancesUrl() + fmt.Sprintf("/%s", args[0])

		// call command
		if err := runCommand(url, "instance", "PUT"); err != nil {
			cmd.SilenceUsage = true
			return err
		}

		return nil
	},
}
