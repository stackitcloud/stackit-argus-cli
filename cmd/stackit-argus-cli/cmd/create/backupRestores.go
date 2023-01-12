package create

/*
 * Create a backup restore.
 */

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
)

// BackupRestoresCmd represents the backupRestores command
var BackupRestoresCmd = &cobra.Command{
	Use:   "backupRestore <backupDate>",
	Short: "Restore a backup.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		restoreTargets := config.GetTargets()

		if restoreTargets == nil {
			return errors.New("required flag \"--target(-t)\" not set")
		}

		if len(restoreTargets) != 1 {
			return errors.New("target requires only one element")
		}

		// check if target contains wrong values
		for _, target := range restoreTargets {
			switch target {
			case "alertConfig", "alertRules", "scrapeConfig", "grafana":
			default:
				cmd.SilenceUsage = true

				return errors.New("wrong targets, possible targets: alertConfig, alertRules, scrapeConfig or grafana")
			}
		}

		// generate an url
		url := config.GetBaseUrl() + fmt.Sprintf("backup-restores/%s", args[0])

		// call command
		if err := runCommand(url, "backup restores", "restoreTarget", restoreTargets); err != nil {
			cmd.SilenceUsage = true

			return err
		}

		return nil
	},
}
