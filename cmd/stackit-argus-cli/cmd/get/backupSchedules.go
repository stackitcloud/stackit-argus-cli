package get

/*
 * Get backup schedules.
 */

import (
	"encoding/json"
	"github.com/spf13/cobra"
	config2 "github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/config"
	output_table "github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/outputTable"
)

// schedule is used to unmarshal backup schedules response
type schedule struct {
	Schedule   string `json:"schedule"`
	ScheduleId string `json:"scheduleId"`
}

// backupSchedules is used to unmarshal backup schedules response
type backupSchedules struct {
	AlertConfig  []schedule `json:"alertConfigBackupSchedules"`
	ScrapeConfig []schedule `json:"scrapeConfigBackupSchedules"`
	AlertRules   []schedule `json:"alertRulesBackupSchedules"`
	Grafana      []schedule `json:"grafanaBackupSchedules"`
}

// backupSchedulesTable holds structure of backup schedules outputTable
type backupSchedulesTable struct {
	Resource   string `header:"resource"`
	Schedule   string `header:"schedule"`
	ScheduleId string `header:"scheduleId"`
}

// printBackupSchedulesTable prints backup schedules response body as outputTable
func printBackupSchedulesTable(body []byte) error {
	var check [4]bool
	var backupSchedules backupSchedules
	var table []backupSchedulesTable

	// unmarshal response body
	if err := json.Unmarshal(body, &backupSchedules); err != nil {
		return err
	}

	// fill the outputTable
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

	// print the outputTable
	output_table.PrintTable(table)

	return nil
}

// BackupSchedulesCmd represents the backupSchedules command
var BackupSchedulesCmd = &cobra.Command{
	Use:   "backupSchedules",
	Short: "Get backup schedules.",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		// generate an url
		url := config2.GetBaseUrl() + "backup-schedules"

		// get output flag
		outputType := config2.GetOutputType()

		// call the command
		body, err := runCommand(url, "backup schedules", outputType)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		// print outputTable output
		if body != nil && (outputType == "" || outputType == "wide") {
			if err := printBackupSchedulesTable(body); err != nil {
				cmd.SilenceUsage = true
				return err
			}
		}

		return err
	},
}
