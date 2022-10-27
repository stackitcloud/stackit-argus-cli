package create

/*
 * Create a backup.
 */

import (
	"fmt"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/utils"

	"github.com/spf13/cobra"
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

		// print debug messages if debug mode is turned on
		if config.IsDebugMode() {
			fmt.Println("create backup command called")
			fmt.Printf("url to call - %s\n", url)
		}

		// create the backup
		status := postRequest(url, backupTargets)

		// print response status
		utils.ResponseMessage(status, "backup", "create")
	},
}

// init flags
func init() {
	BackupCmd.Flags().StringSliceVarP(&backupTargets, "target", "t", nil,
		"defines targets of the backup, possible targets: alertConfig, alertRules, scrapeConfig or grafana")
}
