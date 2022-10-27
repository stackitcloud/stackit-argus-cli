package create

/*
 * Create a backup restore.
 */

import (
	"fmt"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/utils"

	"github.com/spf13/cobra"
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

		// print debug messages if debug mode is turned on
		if config.IsDebugMode() {
			fmt.Println("restore backup command called")
			fmt.Printf("url to call - %s\n", url)
		}

		// create the restore
		status := postRequest(url, nil)

		// print response status
		utils.ResponseMessage(status, "restore backup", "create")
	},
}

// init flags
func init() {
	BackupRestoresCmd.Flags().StringVarP(&restoreTarget, "target", "t", "",
		"defines restore target, possible targets: alertConfig, alertRules, scrapeConfig or grafana")
	err := BackupRestoresCmd.MarkFlagRequired("target")
	cobra.CheckErr(err)
}
