package get

/*
 * Get logs configs.
 */

import (
	"encoding/json"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/output"
	"github.com/stackitcloud/stackit-argus-cli/internal/config"
)

// logsConfigs is used to unmarshal logs configs response body
type logsConfigs struct {
	Config struct {
		Retention string `json:"retention" header:"retention"`
	} `json:"config"`
}

// printLogsConfigsTable prints logs configs response body as output
func printLogsConfigsTable(body []byte, outputType config.OutputType) error {
	var lc logsConfigs

	// unmarshal response body
	if err := json.Unmarshal(body, &lc); err != nil {
		return err
	}

	// print the output
	output.PrintTable(lc.Config, string(outputType))

	return nil
}

// LogsConfigsCmd represents the logsConfigs command
var LogsConfigsCmd = &cobra.Command{
	Use:   "logs-configs",
	Short: "Get logs configuration.",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		// generate an url
		url := config.GetBaseUrl() + "logs-configs"

		// get output flag
		outputType := config.GetOutputType()

		// call the command
		body, err := runCommand(url, "logs configs", outputType)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		// print output output
		if body != nil && outputType != "yaml" && outputType != "json" {
			if err := printLogsConfigsTable(body, outputType); err != nil {
				cmd.SilenceUsage = true
				return err
			}
		}

		return nil
	},
}
