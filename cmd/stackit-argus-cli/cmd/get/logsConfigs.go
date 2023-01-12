package get

/*
 * Get logs configs.
 */

import (
	"encoding/json"
	"github.com/spf13/cobra"
	config2 "github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/config"
	output_table "github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/outputTable"
)

// logsConfigs is used to unmarshal logs configs response body
type logsConfigs struct {
	Config struct {
		Retention string `json:"retention" header:"retention"`
	} `json:"config"`
}

// printLogsConfigsTable prints logs configs response body as outputTable
func printLogsConfigsTable(body []byte) error {
	var logsConfigs logsConfigs

	// unmarshal response body
	if err := json.Unmarshal(body, &logsConfigs); err != nil {
		return err
	}

	// print the outputTable
	output_table.PrintTable(logsConfigs.Config)

	return nil
}

// LogsConfigsCmd represents the logsConfigs command
var LogsConfigsCmd = &cobra.Command{
	Use:   "logsConfigs",
	Short: "Get logs configuration.",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		// generate an url
		url := config2.GetBaseUrl() + "logs-configs"

		// get output flag
		outputType := config2.GetOutputType()

		// call the command
		body, err := runCommand(url, "logs configs", outputType)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		// print outputTable output
		if body != nil && (outputType == "" || outputType == "wide") {
			if err := printLogsConfigsTable(body); err != nil {
				cmd.SilenceUsage = true
				return err
			}
		}

		return nil
	},
}
