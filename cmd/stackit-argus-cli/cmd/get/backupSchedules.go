package get

/*
 * Get backup schedules.
 */

import (
	"encoding/json"
	"fmt"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/utils"

	"github.com/spf13/cobra"
)

type schedule struct {
	Schedule   string `json:"schedule"`
	ScheduleId string `json:"scheduleId"`
}

type backupSchedules struct {
	AlertConfig  []schedule `json:"alertConfigBackupSchedules"`
	ScrapeConfig []schedule `json:"scrapeConfigBackupSchedules"`
	AlertRules   []schedule `json:"alertRulesBackupSchedules"`
	Grafana      []schedule `json:"grafanaBackupSchedules"`
}

type backupSchedulesTable struct {
	Resource   string `header:"resource"`
	Schedule   string `header:"schedule"`
	ScheduleId string `header:"scheduleId"`
}

func printBackupSchedules(body []byte) {
	var check [4]bool
	var backupSchedules backupSchedules
	var table []backupSchedulesTable

	// unmarshal response body
	err := json.Unmarshal(body, &backupSchedules)
	cobra.CheckErr(err)

	// fill the table
	for i := 0; !check[0] && !check[1] && !check[2] && !check[3]; i++ {
		if i < len(backupSchedules.AlertConfig) {
			table = append(table, backupSchedulesTable{
				Resource:   "Alert Config",
				Schedule:   backupSchedules.AlertConfig[i].Schedule,
				ScheduleId: backupSchedules.AlertConfig[i].ScheduleId,
			})
		} else {
			check[0] = true
		}
		if i < len(backupSchedules.AlertRules) {
			table = append(table, backupSchedulesTable{
				Resource:   "Alert Rules",
				Schedule:   backupSchedules.AlertRules[i].Schedule,
				ScheduleId: backupSchedules.AlertRules[i].ScheduleId,
			})
		} else {
			check[1] = true
		}
		if i < len(backupSchedules.ScrapeConfig) {
			table = append(table, backupSchedulesTable{
				Resource:   "Scrape Config",
				Schedule:   backupSchedules.ScrapeConfig[i].Schedule,
				ScheduleId: backupSchedules.ScrapeConfig[i].ScheduleId,
			})
		} else {
			check[2] = true
		}
		if i < len(backupSchedules.Grafana) {
			table = append(table, backupSchedulesTable{
				Resource:   "Grafana",
				Schedule:   backupSchedules.Grafana[i].Schedule,
				ScheduleId: backupSchedules.Grafana[i].ScheduleId,
			})
		} else {
			check[3] = true
		}
	}

	// print the table
	utils.PrintTable(table)
}

// BackupSchedulesCmd represents the backupSchedules command
var BackupSchedulesCmd = &cobra.Command{
	Use:   "backupSchedules",
	Short: "Get backup schedules.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// generate an url
		url := config.GetBaseUrl() + "backup-schedules"

		// print debug messages if debug mode is turned on
		if config.IsDebugMode() {
			fmt.Println("get backup schedules command called")
			fmt.Printf("url to call - %s\n", url)
		}

		// get backup schedules
		status, body := getRequest(url)

		// print response status
		utils.ResponseMessage(status, "backup schedules", "get")

		// print response body
		if status == 200 {
			outputType := config.GetOutputType()
			if outputType == "json" || outputType == "yaml" {
				utils.PrintYamlOrJson(body, string(outputType))
			} else {
				printBackupSchedules(body)
			}
		}
	},
}
