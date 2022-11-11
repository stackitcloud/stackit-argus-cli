package get

/*
 * Get backup retentions.
 */

import (
	"encoding/json"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/utils"

	"github.com/spf13/cobra"
)

// backupRetentions is used to unmarshal backup retentions response body
type backupRetentions struct {
	AlertConfig  string `json:"alertConfigBackupRetention" header:"alert config"`
	ScrapeConfig string `json:"scrapeConfigBackupRetention" header:"scrape config"`
	AlertRules   string `json:"alertRulesBackupRetention" header:"alert rules"`
	Grafana      string `json:"grafanaBackupRetention" header:"grafana"`
}

// printBackupRetentionsTable prints backup retentions
func printBackupRetentionsTable(body []byte) {
	var backupRetentions backupRetentions

	// unmarshal response body
	err := json.Unmarshal(body, &backupRetentions)
	cobra.CheckErr(err)

	// print the table
	utils.PrintTable(backupRetentions)
}

// BackupRetentionsCmd represents the backupRetentions command
var BackupRetentionsCmd = &cobra.Command{
	Use:   "backupRetentions",
	Short: "Get backup retentions.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// generate an url
		url := config.GetBaseUrl() + "backup-retentions"

		// get output flag
		outputType := config.GetOutputType()

		// call the command
		body := runCommand(url, "backup retentions", outputType)

		// print table output
		if body != nil && (outputType == "" || outputType == "wide") {
			printBackupRetentionsTable(body)
		}
	},
}
