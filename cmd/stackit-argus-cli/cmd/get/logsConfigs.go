package get

/*
 * Get logs configs.
 */

import (
	"encoding/json"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/utils"

	"github.com/spf13/cobra"
)

// logsConfigs is used to unmarshal logs configs response body
type logsConfigs struct {
	Config struct {
		Retention string `json:"retention" header:"retention"`
	} `json:"config"`
}

// printLogsConfigsTable prints logs configs response body as table
func printLogsConfigsTable(body []byte) {
	var logsConfigs logsConfigs

	// unmarshal response body
	err := json.Unmarshal(body, &logsConfigs)
	cobra.CheckErr(err)

	// print the table
	utils.PrintTable(logsConfigs.Config)
}

// LogsConfigsCmd represents the logsConfigs command
var LogsConfigsCmd = &cobra.Command{
	Use:   "logsConfigs",
	Short: "Get logs configuration.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// generate an url
		url := config.GetBaseUrl() + "logs-configs"

		// get output flag
		outputType := config.GetOutputType()

		// call the command
		body := runCommand(url, "logs configs", outputType)

		// print table output
		if body != nil && (outputType == "" || outputType == "wide") {
			printLogsConfigsTable(body)
		}
	},
}
