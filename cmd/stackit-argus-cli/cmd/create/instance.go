package create

/*
 * Create an instance.
 */

import (
	"errors"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/internal/config"
)

// InstanceCmd represents the instance command
var InstanceCmd = &cobra.Command{
	Use:   "instance",
	Short: "Create a new instance.",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		if config.GetBodyFile() == "" {
			return errors.New("required flag \"--file(-f)\" not set")
		}

		// generate an url
		url := config.GetInstancesUrl()

		// call command
		if err := runCommand(url, "instance", "", nil); err != nil {
			cmd.SilenceUsage = true
			return err
		}

		return nil
	},
}
