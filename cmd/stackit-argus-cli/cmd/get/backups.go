package get

/*
 * Get backups.
 */

import (
	"encoding/json"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/utils"

	"github.com/spf13/cobra"
)

// backups struct is used to unmarshal backups response body
type backups struct {
	AlertConfig  []string `json:"alertConfigBackups" header:"alert config"`
	ScrapeConfig []string `json:"scrapeConfigBackups" header:"scrape config"`
	AlertRules   []string `json:"alertRulesBackups" header:"alert rules"`
	Grafana      []string `json:"grafanaBackups" header:"grafana"`
}

// printBackupsTable prints backups response body as a table
func printBackupsTable(body []byte) {
	var backups backups

	// unmarshal response body
	err := json.Unmarshal(body, &backups)
	cobra.CheckErr(err)

	// print the table
	utils.PrintTable(backups)
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
		body := runCommand(url, "backups", outputType)

		// print table output
		if body != nil && (outputType == "" || outputType == "wide") {
			printBackupRetentionsTable(body)
		}
	},
}
