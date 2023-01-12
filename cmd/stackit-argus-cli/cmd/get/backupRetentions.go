package get

/*
 * Get backup retentions.
 */

import (
	"encoding/json"
	"github.com/spf13/cobra"
	config2 "github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/config"
	output_table "github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/outputTable"
)

// backupRetentions is used to unmarshal backup retentions response body
type backupRetentions struct {
	AlertConfig  string `json:"alertConfigBackupRetention" header:"alert config"`
	ScrapeConfig string `json:"scrapeConfigBackupRetention" header:"scrape config"`
	AlertRules   string `json:"alertRulesBackupRetention" header:"alert rules"`
	Grafana      string `json:"grafanaBackupRetention" header:"grafana"`
}

// printBackupRetentionsTable prints backup retentions
func printBackupRetentionsTable(body []byte) error {
	var backupRetentions backupRetentions

	// unmarshal response body
	if err := json.Unmarshal(body, &backupRetentions); err != nil {
		return err
	}

	// print the outputTable
	output_table.PrintTable(backupRetentions)

	return nil
}

// BackupRetentionsCmd represents the backupRetentions command
var BackupRetentionsCmd = &cobra.Command{
	Use:   "backupRetentions",
	Short: "Get backup retentions.",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		// generate an url
		url := config2.GetBaseUrl() + "backup-retentions"

		// get output flag
		outputType := config2.GetOutputType()

		// call the command
		body, err := runCommand(url, "backup retentions", outputType)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		// print outputTable output
		if body != nil && (outputType == "" || outputType == "wide") {
			if err := printBackupRetentionsTable(body); err != nil {
				cmd.SilenceUsage = true
				return err
			}
		}

		return err
	},
}
