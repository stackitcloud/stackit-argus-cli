package create

/*
 * Create a backup schedule.
 */

import (
	"errors"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/internal/config"
)

// BackupSchedulesCmd represents the backupSchedules command
var BackupSchedulesCmd = &cobra.Command{
	Use:     "backup-schedule",
	Short:   "Create a backup schedule.",
	Args:    cobra.NoArgs,
	Example: "stackit-argus-cli create backupSchedule -t=\"grafana,scrapeConfig\" -f test",
	RunE: func(cmd *cobra.Command, args []string) error {
		if config.GetBodyFile() == "" {
			return errors.New("required flag \"--file(-f)\" not set")
		}

		backupTargets := config.GetTargets()

		// check if target contains wrong values
		for _, target := range backupTargets {
			switch target {
			case "alertConfig", "alertRules", "scrapeConfig", "grafana":
			default:
				cmd.SilenceUsage = true
				return errors.New("wrong targets, possible targets: alertConfig, alertRules, scrapeConfig or grafana")
			}
		}

		// generate an url
		url := config.GetBaseUrl() + "backup-schedules"

		// call command
		if err := runCommand(url, "backup schedules", "backupTarget", backupTargets); err != nil {
			cmd.SilenceUsage = true
			return err
		}

		return nil
	},
}
