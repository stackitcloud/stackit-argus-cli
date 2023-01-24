package get

/*
 * Get traces configs.
 */

import (
	"encoding/json"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/output"
	"github.com/stackitcloud/stackit-argus-cli/internal/config"
)

// tracesConfigs is used to unmarshal traces configs response body and generate a output out of it
type tracesConfigs struct {
	Config struct {
		Retention string `json:"retention" header:"retention"`
	} `json:"config"`
}

// printTracesConfigsListTable prints traces configs response body as output
func printTracesConfigsListTable(body []byte, outputType config.OutputType) error {
	var tc tracesConfigs

	// unmarshal response body
	if err := json.Unmarshal(body, &tc); err != nil {
		return err
	}

	// print the output
	output.PrintTable(tc.Config, string(outputType))

	return nil
}

// TracesConfigsCmd represents the tracesConfigs command
var TracesConfigsCmd = &cobra.Command{
	Use:   "traces-configs",
	Short: "Get traces config.",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		// generate an url
		url := config.GetBaseUrl() + "traces-configs"

		// get output flag
		outputType := config.GetOutputType()

		// call the command
		body, err := runCommand(url, "traces configs", outputType)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		// print output output
		if body != nil && outputType != "yaml" && outputType != "json" {
			if err := printTracesConfigsListTable(body, outputType); err != nil {
				cmd.SilenceUsage = true
				return err
			}
		}

		return nil
	},
}
