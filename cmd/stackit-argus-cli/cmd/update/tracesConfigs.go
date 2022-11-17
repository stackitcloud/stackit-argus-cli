package update

/*
 * Update a traces config.
 */

import (
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
)

// TracesConfigsCmd represents the tracesConfigs command
var TracesConfigsCmd = &cobra.Command{
	Use:   "tracesConfig",
	Short: "Update a traces config.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// generate an url
		url := config.GetBaseUrl() + "traces-configs"

		// call command
		err := runCommand(url, "traces config", "PUT")
		cobra.CheckErr(err)
	},
}
