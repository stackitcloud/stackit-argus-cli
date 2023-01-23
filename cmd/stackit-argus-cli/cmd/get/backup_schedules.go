package get

/*
 * Get backup schedules.
 */

import (
	"encoding/json"
	"github.com/spf13/cobra"
	config2 "github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/output_table"
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

// backupSchedulesTable holds structure of backup schedules output_table
type backupSchedulesTable struct {
	Resource   string `header:"resource"`
	Schedule   string `header:"schedule"`
	ScheduleId string `header:"scheduleId"`
}

// printBackupSchedulesTable prints backup schedules response body as output_table
func printBackupSchedulesTable(body []byte, outputType config2.OutputType) error {
	var check [4]bool
	var bs backupSchedules
	var table []backupSchedulesTable

	// unmarshal response body
	if err := json.Unmarshal(body, &bs); err != nil {
		return err
	}

	// fill the output_table
	for i := 0; !check[0] && !check[1] && !check[2] && !check[3]; i++ {
		if i < len(bs.AlertConfig) {
			table = append(table, backupSchedulesTable{
				Resource:   "Alert Config",
				Schedule:   bs.AlertConfig[i].Schedule,
				ScheduleId: bs.AlertConfig[i].ScheduleId,
			})
		} else {
			check[0] = true
		}
		if i < len(bs.AlertRules) {
			table = append(table, backupSchedulesTable{
				Resource:   "Alert Rules",
				Schedule:   bs.AlertRules[i].Schedule,
				ScheduleId: bs.AlertRules[i].ScheduleId,
			})
		} else {
			check[1] = true
		}
		if i < len(bs.ScrapeConfig) {
			table = append(table, backupSchedulesTable{
				Resource:   "Scrape Config",
				Schedule:   bs.ScrapeConfig[i].Schedule,
				ScheduleId: bs.ScrapeConfig[i].ScheduleId,
			})
		} else {
			check[2] = true
		}
		if i < len(bs.Grafana) {
			table = append(table, backupSchedulesTable{
				Resource:   "Grafana",
				Schedule:   bs.Grafana[i].Schedule,
				ScheduleId: bs.Grafana[i].ScheduleId,
			})
		} else {
			check[3] = true
		}
	}

	// print the output_table
	output_table.PrintTable(table, outputType)

	return nil
}

// BackupSchedulesCmd represents the backupSchedules command
var BackupSchedulesCmd = &cobra.Command{
	Use:   "backup-schedules",
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

		// print output_table output
		if body != nil && outputType != "yaml" && outputType != "json" {
			if err := printBackupSchedulesTable(body, outputType); err != nil {
				cmd.SilenceUsage = true
				return err
			}
		}

		return err
	},
}
