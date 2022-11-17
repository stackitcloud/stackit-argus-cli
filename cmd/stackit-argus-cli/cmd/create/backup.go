package create

/*
 * Create a backup.
 */

import (
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
)

var backupTargets []string

// BackupCmd represents the backup command
var BackupCmd = &cobra.Command{
	Use:     "backup",
	Short:   "Create a backup.",
	Args:    cobra.NoArgs,
	Example: "stackit-argus-cli create backup -t=\"grafana,scrapeConfig\" -f test",
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
		url := config.GetBaseUrl() + "backups"

		// call command
		err := runCommand(url, "backups", nil)
		cobra.CheckErr(err)
	},
}

// init flags
func init() {
	BackupCmd.Flags().StringSliceVarP(&backupTargets, "target", "t", nil,
		"defines targets of the backup, possible targets: alertConfig, alertRules, scrapeConfig or grafana")
}
