package get

/*
 * Get backups.
 */

import (
	"encoding/json"
	"fmt"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/utils"

	"github.com/spf13/cobra"
)

// backups struct is used to unmarshal backups response body
type backups struct {
	AlertConfig  []string `json:"alertConfigBackups"`
	ScrapeConfig []string `json:"scrapeConfigBackups"`
	AlertRules   []string `json:"alertRulesBackups"`
	Grafana      []string `json:"grafanaBackups"`
}

// backupsTable holds structure of backups table
type backupsTable struct {
	AlertConfig  []string `header:"alert config"`
	ScrapeConfig []string `header:"scrape config"`
	AlertRules   []string `header:"alert rules"`
	Grafana      []string `header:"grafana"`
}

// printBackups prints backups
func printBackups(body []byte) {
	var backups backups

	// unmarshal response body
	err := json.Unmarshal(body, &backups)
	cobra.CheckErr(err)

	// print the table
	utils.PrintTable(backupsTable{
		AlertConfig:  backups.AlertConfig,
		ScrapeConfig: backups.ScrapeConfig,
		AlertRules:   backups.AlertRules,
		Grafana:      backups.Grafana,
	})
}

// BackupCmd represents the backup command
var BackupCmd = &cobra.Command{
	Use:   "backups",
	Short: "Get backups.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// generate an url
		url := config.GetBaseUrl() + "backups"

		// print debug messages if debug mode is turned on
		if config.IsDebugMode() {
			fmt.Println("get backups command called")
			fmt.Printf("url to call - %s\n", url)
		}

		// get backups
		status, body := getRequest(url)

		// print response status
		utils.ResponseMessage(status, "backups", "get")

		// print response body
		if status == 200 {
			outputType := config.GetOutputType()
			if outputType == "json" || outputType == "yaml" {
				utils.PrintYamlOrJson(body, string(outputType))
			} else {
				printBackups(body)
			}
		}
	},
}
