package get

/*
 * Get backups.
 */

import (
	"encoding/json"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	output_table "github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/outputTable"
)

// backups struct is used to unmarshal backups response body
type backups struct {
	AlertConfig  []string `json:"alertConfigBackups" header:"alert config"`
	ScrapeConfig []string `json:"scrapeConfigBackups" header:"scrape config"`
	AlertRules   []string `json:"alertRulesBackups" header:"alert rules"`
	Grafana      []string `json:"grafanaBackups" header:"grafana"`
}

// printBackupsTable prints backups response body as a outputTable
func printBackupsTable(body []byte, outputType config.OutputType) {
	var backups backups

	// unmarshal response body
	err := json.Unmarshal(body, &backups)
	cobra.CheckErr(err)

	if outputType != "wide" {
		l := len(backups.Grafana)
		if l > 10 {
			backups.Grafana = backups.Grafana[l-10 : l]
		}
		l = len(backups.AlertConfig)
		if l > 10 {
			backups.AlertConfig = backups.AlertConfig[l-10 : l]
		}
		l = len(backups.ScrapeConfig)
		if l > 10 {
			backups.ScrapeConfig = backups.ScrapeConfig[l-10 : l]
		}
		l = len(backups.AlertRules)
		if l > 10 {
			backups.AlertRules = backups.AlertRules[l-10 : l]
		}
	}

	// print the outputTable
	output_table.PrintTable(backups)
}

// BackupCmd represents the backup command
var BackupCmd = &cobra.Command{
	Use:   "backups",
	Short: "Get backups.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// generate an url
		url := config.GetBaseUrl() + "backups"

		// get output flag
		outputType := config.GetOutputType()

		// call the command
		body, err := runCommand(url, "backups", outputType)
		cobra.CheckErr(err)

		// print outputTable output
		if body != nil && (outputType == "" || outputType == "wide") {
			printBackupsTable(body, outputType)
		}
	},
}
