package update

/*
 * Update an alert config.
 */

import (
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
)

// AlertConfigsCmd represents the alertConfigs command
var AlertConfigsCmd = &cobra.Command{
	Use:   "alertConfig",
	Short: "Update an alert config.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// generate an url
		url := config.GetBaseUrl() + "alertconfigs"

		// call command
		err := runCommand(url, "alert config", "PUT")
		cobra.CheckErr(err)
	},
}
