package create

/*
 * Create a scrape config.
 */

import (
	"fmt"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/utils"

	"github.com/spf13/cobra"
)

// ScrapeConfigsCmd represents the scrapeConfigs command
var ScrapeConfigsCmd = &cobra.Command{
	Use:   "scrapeConfig",
	Short: "Create a scrape config.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// generate an url
		url := config.GetBaseUrl() + "scrapeconfigs"

		// print debug messages if debug mode is turned on
		if config.IsDebugMode() {
			fmt.Println("create scrape config command called")
			fmt.Printf("url to call - %s\n", url)
		}

		// create the scrape config
		status := postRequest(url, nil)

		// print response status
		utils.ResponseMessage(status, "scrape config", "create")
	},
}
