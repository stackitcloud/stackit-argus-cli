package get

/*
 * Get backup retentions.
 */

import (
	"encoding/json"
	"fmt"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/utils"

	"github.com/spf13/cobra"
)

// backupRetentions is used to unmarshal backup retentions response body
type backupRetentions struct {
	AlertConfig  string `json:"alertConfigBackupRetention"`
	ScrapeConfig string `json:"scrapeConfigBackupRetention"`
	AlertRules   string `json:"alertRulesBackupRetention"`
	Grafana      string `json:"grafanaBackupRetention"`
}

// backupRetentionsTable holds structure of backup retentions table
type backupRetentionsTable struct {
	AlertConfig  string `header:"alert config"`
	ScrapeConfig string `header:"scrape config"`
	AlertRules   string `header:"alert rules"`
	Grafana      string `header:"grafana"`
}

// printBackupRetentions prints backup retentions
func printBackupRetentions(body []byte) {
	var backupRetentions backupRetentions

	// unmarshal response body
	err := json.Unmarshal(body, &backupRetentions)
	cobra.CheckErr(err)

	// print the table
	utils.PrintTable(backupRetentionsTable{
		AlertConfig:  backupRetentions.AlertConfig,
		ScrapeConfig: backupRetentions.ScrapeConfig,
		AlertRules:   backupRetentions.AlertRules,
		Grafana:      backupRetentions.Grafana,
	})
}

// BackupRetentionsCmd represents the backupRetentions command
var BackupRetentionsCmd = &cobra.Command{
	Use:   "backupRetentions",
	Short: "Get backup retentions.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// generate an url
		url := config.GetBaseUrl() + "backup-retentions"

		// print debug messages if debug mode is turned on
		if config.IsDebugMode() {
			fmt.Println("get backup retentions command called")
			fmt.Printf("url to call - %s\n", url)
		}

		// get backup retentions
		status, body := getRequest(url)

		// print response status
		utils.ResponseMessage(status, "backup retentions", "get")

		// print response body
		if status == 200 {
			outputType := config.GetOutputType()
			if outputType == "json" || outputType == "yaml" {
				utils.PrintYamlOrJson(body, string(outputType))
			} else {
				printBackupRetentions(body)
			}
		}
	},
}
