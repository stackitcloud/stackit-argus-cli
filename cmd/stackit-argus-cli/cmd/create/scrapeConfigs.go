package create

/*
 * Create a scrape config.
 */

import (
	"errors"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/config"
)

// ScrapeConfigsCmd represents the scrapeConfigs command
var ScrapeConfigsCmd = &cobra.Command{
	Use:   "scrapeConfig",
	Short: "Create a scrape config.",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		if config.GetBodyFile() == "" {
			return errors.New("required flag \"--file(-f)\" not set")
		}

		// generate an url
		url := config.GetBaseUrl() + "scrapeconfigs"

		// call command
		if err := runCommand(url, "scrape configs", "", nil); err != nil {
			cmd.SilenceUsage = true

			return err
		}

		return nil
	},
}
