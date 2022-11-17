package update

/*
 * Update scrape configs.
 */

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
)

// ScrapeConfigsCmd represents the scrapeConfigs command
var ScrapeConfigsCmd = &cobra.Command{
	Use:   "scrapeConfigs <jobName>",
	Short: "Update scrape configs.",
	Long:  "Patch scrape configs if job name was nit specified, otherwise update scrape config.",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
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
		err := runCommand(url, resource, method)
		cobra.CheckErr(err)
	},
}
