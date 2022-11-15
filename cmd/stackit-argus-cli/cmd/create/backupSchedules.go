package create

/*
 * Create a backup schedule.
 */

import (
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
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

		// call command
		runCommand(url, "backup schedules", backupTargets)
	},
}

// init flags
func init() {
	BackupSchedulesCmd.Flags().StringSliceVarP(&backupTargets, "target", "t", nil,
		"defines targets of the backup schedules, possible targets: alertConfig, alertRules, scrapeConfig or grafana")
}
