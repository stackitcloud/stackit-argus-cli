package delete

/*
 * Delete a scrape config.
 */

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
)

// ScrapeConfigsCmd represents the scrapeConfigs command
var ScrapeConfigsCmd = &cobra.Command{
	Use:   "scrapeConfig",
	Short: "Delete scrape config.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// generate an url
		url := config.GetBaseUrl() + fmt.Sprintf("scrapeconfigs/%s", args[0])

		// call command
		err := runCommand(url, "scrape config")
		cobra.CheckErr(err)
	},
}
