package get

/*
 * Get backup retentions.
 */

import (
	"encoding/json"
	"github.com/spf13/cobra"
	config2 "github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/output_table"
)

// backupRetentions is used to unmarshal backup retentions response body
type backupRetentions struct {
	AlertConfig  string `json:"alertConfigBackupRetention" header:"alert config"`
	ScrapeConfig string `json:"scrapeConfigBackupRetention" header:"scrape config"`
	AlertRules   string `json:"alertRulesBackupRetention" header:"alert rules"`
	Grafana      string `json:"grafanaBackupRetention" header:"grafana"`
}

// printBackupRetentionsTable prints backup retentions
func printBackupRetentionsTable(body []byte, outputType config2.OutputType) error {
	var br backupRetentions

	// unmarshal response body
	if err := json.Unmarshal(body, &br); err != nil {
		return err
	}

	// print the output_table
	output_table.PrintTable(br, outputType)

	return nil
}

// BackupRetentionsCmd represents the backupRetentions command
var BackupRetentionsCmd = &cobra.Command{
	Use:   "backup-retentions",
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

		// print output_table output
		if body != nil && outputType != "yaml" && outputType != "json" {
			if err := printBackupRetentionsTable(body, outputType); err != nil {
				cmd.SilenceUsage = true
				return err
			}
		}

		return err
	},
}
