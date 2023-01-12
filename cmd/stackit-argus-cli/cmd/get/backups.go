package get

/*
 * Get backups.
 */

import (
	"encoding/json"
	"github.com/spf13/cobra"
	config2 "github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/config"
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
func printBackupsTable(body []byte, outputType config2.OutputType) error {
	var backups backups

	// unmarshal response body
	if err := json.Unmarshal(body, &backups); err != nil {
		return err
	}

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

		// print outputTable output
		if body != nil && (outputType == "" || outputType == "wide") {
			if err := printBackupsTable(body, outputType); err != nil {
				cmd.SilenceUsage = true
				return err
			}
		}

		return nil
	},
}
