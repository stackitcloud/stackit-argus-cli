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

type scrapeConfigData struct {
	JobName        string `json:"jobName"`
	ScrapeInterval string `json:"scrapeInterval"`
	ScrapeTimeout  string `json:"scrapeTimeout"`
	MetricsPath    string `json:"metricsPath"`
	// wide table attributes
	Scheme          string `json:"scheme"`
	SampleLimit     int    `json:"sampleLimit"`
	HonorLabels     bool   `json:"honorLabels"`
	HonorTimeStamps bool   `json:"honorTimeStamps"`

	StaticConfigs []struct {
		Targets []string `json:"targets"`
	} `json:"staticConfigs"`
}

// scrapeConfig is used to unmarshal scrape config response body
type scrapeConfig struct {
	Data scrapeConfigData `json:"data"`
}

// scrapeConfigsList is used to unmarshal scrape configs response body
type scrapeConfigsList struct {
	Data []scrapeConfigData `json:"data"`
}

// scrapeConfigsTable holds structure of scrape configs table
type scrapeConfigsTable struct {
	JobName        string   `header:"job name"`
	ScrapeInterval string   `header:"scrape interval"`
	ScrapeTimeout  string   `header:"scrape timeout"`
	SampleLimit    int      `header:"sample limit"`
	Targets        []string `header:"targets"`
	// wide table attributes
	MetricsPath     string `header:"metrics path"`
	HonorLabels     bool   `header:"honor labels"`
	HonorTimeStamps bool   `header:"honor time stamps"`
}

// printScrapeConfigTable prints scrape config response body as table
func printScrapeConfigTable(body []byte, outputType config.OutputType) {
	var scrapeConfig scrapeConfig
	var targets []string

	// unmarshal response body
	err := json.Unmarshal(body, &scrapeConfig)
	cobra.CheckErr(err)

	// fill the table
	for _, sc := range scrapeConfig.Data.StaticConfigs {
		targets = append(targets, sc.Targets...)
	}
	table := scrapeConfigsTable{
		ScrapeInterval:  scrapeConfig.Data.ScrapeInterval,
		ScrapeTimeout:   scrapeConfig.Data.ScrapeTimeout,
		SampleLimit:     scrapeConfig.Data.SampleLimit,
		Targets:         targets,
		MetricsPath:     scrapeConfig.Data.MetricsPath,
		HonorLabels:     scrapeConfig.Data.HonorLabels,
		HonorTimeStamps: scrapeConfig.Data.HonorTimeStamps,
	}

	// print the table
	if outputType != "wide" {
		utils.PrintTable(utils.RemoveColumnsFromTable(table,
			[]string{"MetricsPath", "HonorLabels", "HonorTimeStamps", "JobName"}))
	} else {
		utils.PrintTable(utils.RemoveColumnsFromTable(table,
			[]string{"JobName"}))
	}
}

// printScrapeConfigsListTable prints scrape configs response body as table
func printScrapeConfigsListTable(body []byte, outputType config.OutputType) {
	var scrapeConfigs scrapeConfigsList
	var table []scrapeConfigsTable
	var targets []string

	// unmarshal response body
	err := json.Unmarshal(body, &scrapeConfigs)
	cobra.CheckErr(err)

	// fill the table
	for _, data := range scrapeConfigs.Data {
		for _, sc := range data.StaticConfigs {
			targets = append(targets, sc.Targets...)
		}
		table = append(table, scrapeConfigsTable{
			JobName:         data.JobName,
			ScrapeInterval:  data.ScrapeInterval,
			ScrapeTimeout:   data.ScrapeTimeout,
			SampleLimit:     data.SampleLimit,
			Targets:         targets,
			MetricsPath:     data.MetricsPath,
			HonorLabels:     data.HonorLabels,
			HonorTimeStamps: data.HonorTimeStamps,
		})
		targets = nil
	}

	// print the table
	if outputType != "wide" {
		var newTable []interface{}

		for _, data := range table {
			newTable = append(newTable, utils.RemoveColumnsFromTable(data,
				[]string{"MetricsPath", "HonorLabels", "HonorTimeStamps"}))
		}
		utils.PrintTable(newTable)
	} else {
		utils.PrintTable(table)
	}
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
				printScrapeConfigsListTable(body, outputType)
			} else {
				printScrapeConfigTable(body, outputType)
			}
		}
	},
}
