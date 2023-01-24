package update

/*
 * Update scrape configs.
 */

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/internal/config"
)

// ScrapeConfigsCmd represents the scrapeConfigs command
var ScrapeConfigsCmd = &cobra.Command{
	Use:   "scrape-configs <job-name>",
	Short: "Update scrape configs.",
	Long:  "Patch scrape configs if job name was nit specified, otherwise update scrape config.",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if config.GetBodyFile() == "" {
			return errors.New("required flag \"--file(-f)\" not set")
		}

		// generate an url
		url := config.GetBaseUrl() + "scrapeconfigs"

		// modify url and debug msg if an argument has been given
		method := "PATCH"
		resource := "scrape configs"
		if len(args) == 1 {
			resource = "scrape config"
			method = "PUT"
			url += fmt.Sprintf("/%s", args[0])
		}

		// call command
		if err := runCommand(url, resource, method); err != nil {
			cmd.SilenceUsage = true
			return err
		}

		return nil
	},
}
