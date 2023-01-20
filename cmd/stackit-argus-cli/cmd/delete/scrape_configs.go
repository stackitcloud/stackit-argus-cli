package delete

/*
 * Delete a scrape config.
 */

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/config"
)

// ScrapeConfigsCmd represents the scrapeConfigs command
var ScrapeConfigsCmd = &cobra.Command{
	Use:   "scrape-config <job-name>",
	Short: "Delete scrape config.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		// generate an url
		url := config.GetBaseUrl() + fmt.Sprintf("scrapeconfigs/%s", args[0])

		// call command
		if err := runCommand(url, "scrape config"); err != nil {
			cmd.SilenceUsage = true
			return err
		}

		return nil
	},
}
