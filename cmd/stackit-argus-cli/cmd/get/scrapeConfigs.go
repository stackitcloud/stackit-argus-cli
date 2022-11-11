package get

/*
 * Get scrape configs.
 */

import (
	"encoding/json"
	"fmt"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/utils"

	"github.com/spf13/cobra"
)

// tracesConfigs is used to unmarshal scrape config response body and generate a table out of it
type scrapeConfig struct {
	Data struct {
		JobName        string `json:"jobName" header:"job name"`
		Scheme         string `json:"scheme" header:"scheme"`
		ScrapeInterval string `json:"scrapeInterval" header:"scrape interval"`
		ScrapeTimeout  string `json:"scrapeTimeout" header:"scrape timeout"`
		MetricsPath    string `json:"metricsPath" header:"metrics path"`
	} `json:"data"`
}

// tracesConfigs is used to unmarshal scrape configs response body and generate a table out of it
type scrapeConfigsList struct {
	Data []struct {
		JobName        string `json:"jobName" header:"job name"`
		Scheme         string `json:"scheme" header:"scheme"`
		ScrapeInterval string `json:"scrapeInterval" header:"scrape interval"`
		ScrapeTimeout  string `json:"scrapeTimeout" header:"scrape timeout"`
		MetricsPath    string `json:"metricsPath" header:"metrics path"`
	} `json:"data"`
}

// printScrapeConfigTable prints scrape config response body as table
func printScrapeConfigTable(body []byte) {
	var scrapeConfig scrapeConfig

	// unmarshal response body
	err := json.Unmarshal(body, &scrapeConfig)
	cobra.CheckErr(err)

	// print the table
	utils.PrintTable(scrapeConfig.Data)
}

// printScrapeConfigsListTable prints scrape configs response body as table
func printScrapeConfigsListTable(body []byte) {
	var scrapeConfigs scrapeConfigsList

	// unmarshal response body
	err := json.Unmarshal(body, &scrapeConfigs)
	cobra.CheckErr(err)

	// print the table
	utils.PrintTable(scrapeConfigs.Data)
}

// ScrapeConfigsCmd represents the scrapeConfigs command
var ScrapeConfigsCmd = &cobra.Command{
	Use:   "scrapeConfigs <jobName>",
	Short: "Get scrape config.",
	Long:  "Get list of scrape config if job name was not specified, otherwise get scrape config.",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// generate an url
		url := config.GetBaseUrl() + "scrapeconfigs"

		// modify url and debug message depend on arguments
		resource := "scrape configs"
		if len(args) == 1 {
			resource = "scrape config"
			url += fmt.Sprintf("/%s", args[0])
		}

		// get output flag
		outputType := config.GetOutputType()

		// call the command
		body := runCommand(url, resource, outputType)

		// print table output
		if body != nil && (outputType == "" || outputType == "wide") {
			if len(args) == 0 {
				printScrapeConfigsListTable(body)
			} else {
				printScrapeConfigTable(body)
			}
		}
	},
}
