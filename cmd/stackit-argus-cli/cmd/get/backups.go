package get

/*
 * Get backups.
 */

import (
	"encoding/json"
	"github.com/spf13/cobra"
	config2 "github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/output_table"
)

var numberOfBackups = 10

// backups struct is used to unmarshal backups response body
type backups struct {
	AlertConfig  []string `json:"alertConfigBackups" header:"alert config"`
	ScrapeConfig []string `json:"scrapeConfigBackups" header:"scrape config"`
	AlertRules   []string `json:"alertRulesBackups" header:"alert rules"`
	Grafana      []string `json:"grafanaBackups" header:"grafana"`
}

// printBackupsTable prints backups response body as an output_table
func printBackupsTable(body []byte, outputType config2.OutputType) error {
	var b backups

	// unmarshal response body
	if err := json.Unmarshal(body, &b); err != nil {
		return err
	}

	if outputType != "wide" && outputType != "wide-table" {
		l := len(b.Grafana)
		if l > numberOfBackups {
			b.Grafana = b.Grafana[l-numberOfBackups : l]
		}
		l = len(b.AlertConfig)
		if l > numberOfBackups {
			b.AlertConfig = b.AlertConfig[l-numberOfBackups : l]
		}
		l = len(b.ScrapeConfig)
		if l > numberOfBackups {
			b.ScrapeConfig = b.ScrapeConfig[l-numberOfBackups : l]
		}
		l = len(b.AlertRules)
		if l > numberOfBackups {
			b.AlertRules = b.AlertRules[l-numberOfBackups : l]
		}
	}

	// print the output_table
	output_table.PrintTable(b, outputType)

	return nil
}

// BackupCmd represents the backup command
var BackupCmd = &cobra.Command{
	Use:   "backups",
	Short: "Get backups.",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		// generate an url
		url := config2.GetBaseUrl() + "backups"

		// get output flag
		outputType := config2.GetOutputType()

		// call the command
		body, err := runCommand(url, "backups", outputType)
		if err != nil {
			cmd.SilenceUsage = true
			return nil
		}

		// print output_table output
		if body != nil && outputType != "yaml" && outputType != "json" {
			if err := printBackupsTable(body, outputType); err != nil {
				cmd.SilenceUsage = true
				return err
			}
		}

		return nil
	},
}
