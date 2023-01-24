package get

/*
 * Get scrape configs.
 */

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/output"
	"github.com/stackitcloud/stackit-argus-cli/internal/config"
)

type scrapeConfigData struct {
	JobName        string `json:"jobName"`
	ScrapeInterval string `json:"scrapeInterval"`
	ScrapeTimeout  string `json:"scrapeTimeout"`
	MetricsPath    string `json:"metricsPath"`
	// wide output attributes
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
	Data scrapeConfigData `json:"data" validate:"required"`
}

// scrapeConfigsList is used to unmarshal scrape configs response body
type scrapeConfigsList struct {
	Data []scrapeConfigData `json:"data" validate:"required"`
}

// scrapeConfigsTable holds structure of scrape configs output
type scrapeConfigsTable struct {
	JobName        string   `header:"job name"`
	ScrapeInterval string   `header:"scrape interval"`
	ScrapeTimeout  string   `header:"scrape timeout"`
	SampleLimit    int      `header:"sample limit"`
	Targets        []string `header:"targets"`
	// wide output attributes
	MetricsPath     string `header:"metrics path"`
	HonorLabels     bool   `header:"honor labels"`
	HonorTimeStamps bool   `header:"honor time stamps"`
}

// printScrapeConfigTable prints scrape config response body as output
func printScrapeConfigTable(body []byte, outputType config.OutputType) error {
	var sc scrapeConfig
	var targets []string

	// unmarshal response body
	if err := json.Unmarshal(body, &sc); err != nil {
		return err
	}

	// fill the output
	for _, sc := range sc.Data.StaticConfigs {
		targets = append(targets, sc.Targets...)
	}
	table := scrapeConfigsTable{
		ScrapeInterval:  sc.Data.ScrapeInterval,
		ScrapeTimeout:   sc.Data.ScrapeTimeout,
		SampleLimit:     sc.Data.SampleLimit,
		Targets:         targets,
		MetricsPath:     sc.Data.MetricsPath,
		HonorLabels:     sc.Data.HonorLabels,
		HonorTimeStamps: sc.Data.HonorTimeStamps,
	}

	// print the output
	if outputType != "wide" {
		output.PrintTable(output.RemoveColumnsFromTable(table,
			[]string{"MetricsPath", "HonorLabels", "HonorTimeStamps", "JobName"}), string(outputType))
	} else {
		output.PrintTable(output.RemoveColumnsFromTable(table,
			[]string{"JobName"}), string(outputType))
	}

	return nil
}

// printScrapeConfigsListTable prints scrape configs response body as output
func printScrapeConfigsListTable(body []byte, outputType config.OutputType) error {
	var scrapeConfigs scrapeConfigsList
	var table []scrapeConfigsTable
	var targets []string

	// unmarshal response body
	if err := json.Unmarshal(body, &scrapeConfigs); err != nil {
		return err
	}

	// fill the output
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

	// print the output
	if outputType != "wide" && outputType != "wide-table" {
		var newTable []interface{}

		for _, data := range table {
			newTable = append(newTable, output.RemoveColumnsFromTable(data,
				[]string{"MetricsPath", "HonorLabels", "HonorTimeStamps"}))
		}
		output.PrintTable(newTable, string(outputType))
	} else {
		output.PrintTable(table, string(outputType))
	}

	return nil
}

// ScrapeConfigsCmd represents the scrapeConfigs command
var ScrapeConfigsCmd = &cobra.Command{
	Use:   "scrape-configs <job-name>",
	Short: "Get scrape config.",
	Long:  "Get list of scrape config if job name was not specified, otherwise get scrape config.",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
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
		body, err := runCommand(url, resource, outputType)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		// print output output
		if body != nil && outputType != "yaml" && outputType != "json" {
			if len(args) == 0 {
				if err := printScrapeConfigsListTable(body, outputType); err != nil {
					cmd.SilenceUsage = true
					return err
				}
			} else {
				if err := printScrapeConfigTable(body, outputType); err != nil {
					cmd.SilenceUsage = true
					return err
				}
			}
		}

		return nil
	},
}
