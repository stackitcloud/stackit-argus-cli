package delete

/*
 * Delete a scrape config.
 */

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/utils"
)

// ScrapeConfigsCmd represents the scrapeConfigs command
var ScrapeConfigsCmd = &cobra.Command{
	Use:   "scrapeConfig",
	Short: "Delete scrape config.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// generate an url
		url := config.GetBaseUrl() + fmt.Sprintf("scrapeconfigs/%s", args[0])

		// print debug messages if debug mode is turned on
		if config.IsDebugMode() {
			fmt.Println("delete scrape config command called")
			fmt.Printf("url to call - %s\n", url)
		}

		// delete the scrape config
		status := deleteRequest(url)

		// print response status
		utils.ResponseMessage(status, "scrape config", "delete")
	},
}
