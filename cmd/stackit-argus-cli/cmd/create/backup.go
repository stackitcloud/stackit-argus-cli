package create

/*
 * Create a backup.
 */

import (
	"errors"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
)

// BackupCmd represents the backup command
var BackupCmd = &cobra.Command{
	Use:     "backup",
	Short:   "Create a backup.",
	Args:    cobra.NoArgs,
	Example: "stackit-argus-cli create backup -t=\"grafana,scrapeConfig\" -f test",
	RunE: func(cmd *cobra.Command, args []string) error {
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
		url := config.GetBaseUrl() + "backups"

		// call command
		if err := runCommand(url, "backups", "backupTarget", backupTargets); err != nil {
			cmd.SilenceUsage = true

			return err
		}

		return nil
	},
}
