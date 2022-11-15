package delete

/*
 * Delete an instance.
 */

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
)

// InstanceCmd represents the instance command
var InstanceCmd = &cobra.Command{
	Use:   "instance <instanceId>",
	Short: "Delete an instance.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// generate an url
		url := config.GetInstancesUrl() + fmt.Sprintf("/%s", args[0])

		// call command
		runCommand(url, "instance")
	},
}
