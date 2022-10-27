package update

/*
 * Update scrape configs.
 */

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/utils"
)

// ScrapeConfigsCmd represents the scrapeConfigs command
var ScrapeConfigsCmd = &cobra.Command{
	Use:   "scrapeConfigs <jobName>",
	Short: "Update scrape configs.",
	Long:  "Patch scrape configs if job name was nit specified, otherwise update scrape config.",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var method string
		var debugMessage string

		// generate an url
		url := config.GetBaseUrl() + "scrapeconfigs"

		// modify url and debug msg if an argument has been given
		if len(args) == 0 {
			debugMessage = "patch scrape configs command called"
			method = "PATCH"
		} else if len(args) == 1 {
			debugMessage = "update scrape config command called"
			method = "PUT"
			url += fmt.Sprintf("/%s", args[0])
		}

		// print debug messages if debug mode is turned on
		if config.IsDebugMode() {
			fmt.Println(debugMessage)
			fmt.Printf("url to call - %s\n", url)
		}

		// update the scrape config
		status := updateRequest(url, method)

		// print response status
		utils.ResponseMessage(status, "scrape config", "update")
	},
}
