package create

/*
 * Create a backup restore.
 */

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
)

var restoreTarget string

// BackupRestoresCmd represents the backupRestores command
var BackupRestoresCmd = &cobra.Command{
	Use:   "backupRestore",
	Short: "Restore a backup.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// check if target contains wrong values
		switch restoreTarget {
		case "alertConfig", "alertRules", "scrapeConfig", "grafana":
		default:
			cobra.CheckErr("wrong target, possible targets: alertConfig, alertRules, scrapeConfig or grafana")
		}

		// generate an url
		url := config.GetBaseUrl() + fmt.Sprintf("backup-restores/%s", args[0])

		// call command
		err := runCommand(url, "backup restores", []string{restoreTarget})
		cobra.CheckErr(err)
	},
}

// init flags
func init() {
	BackupRestoresCmd.Flags().StringVarP(&restoreTarget, "target", "t", "",
		"defines restore target, possible targets: alertConfig, alertRules, scrapeConfig or grafana")
	err := BackupRestoresCmd.MarkFlagRequired("target")
	cobra.CheckErr(err)
}
