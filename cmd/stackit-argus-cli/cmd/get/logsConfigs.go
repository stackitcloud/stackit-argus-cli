package get

/*
 * Get logs configs.
 */

import (
	"encoding/json"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	output_table "github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/outputTable"
)

// logsConfigs is used to unmarshal logs configs response body
type logsConfigs struct {
	Config struct {
		Retention string `json:"retention" header:"retention"`
	} `json:"config"`
}

// printLogsConfigsTable prints logs configs response body as outputTable
func printLogsConfigsTable(body []byte) {
	var logsConfigs logsConfigs

	// unmarshal response body
	err := json.Unmarshal(body, &logsConfigs)
	cobra.CheckErr(err)

	// print the outputTable
	output_table.PrintTable(logsConfigs.Config)
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
		body, err := runCommand(url, "logs configs", outputType)
		cobra.CheckErr(err)

		// print outputTable output
		if body != nil && (outputType == "" || outputType == "wide") {
			printLogsConfigsTable(body)
		}
	},
}
