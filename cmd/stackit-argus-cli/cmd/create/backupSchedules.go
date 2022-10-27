package create

/*
 * Create a backup schedule.
 */

import (
	"fmt"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/utils"

	"github.com/spf13/cobra"
)

// BackupSchedulesCmd represents the backupSchedules command
var BackupSchedulesCmd = &cobra.Command{
	Use:     "backupSchedule",
	Short:   "Create a backup schedule.",
	Args:    cobra.NoArgs,
	Example: "stackit-argus-cli create backupSchedule -t=\"grafana,scrapeConfig\" -f test",
	Run: func(cmd *cobra.Command, args []string) {
		// check if target contains wrong values
		for _, target := range backupTargets {
			switch target {
			case "alertConfig", "alertRules", "scrapeConfig", "grafana":
			default:
				cobra.CheckErr("wrong targets, possible targets: alertConfig, alertRules, scrapeConfig or grafana")
			}
		}

		// generate an url
		url := config.GetBaseUrl() + "backup-schedules"

		// print debug messages if debug mode is turned on
		if config.IsDebugMode() {
			fmt.Println("create backup schedule command called")
			fmt.Printf("url to call - %s\n", url)
		}

		// create the backup schedule
		status := postRequest(url, backupTargets)

		// print response status
		utils.ResponseMessage(status, "backup schedule", "create")
	},
}

// init flags
func init() {
	BackupSchedulesCmd.Flags().StringSliceVarP(&backupTargets, "target", "t", nil,
		"defines targets of the backup schedules, possible targets: alertConfig, alertRules, scrapeConfig or grafana")
}
