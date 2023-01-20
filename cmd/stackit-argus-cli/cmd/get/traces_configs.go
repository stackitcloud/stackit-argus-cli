package get

/*
 * Get traces configs.
 */

import (
	"encoding/json"
	"github.com/spf13/cobra"
	config2 "github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/output_table"
)

// tracesConfigs is used to unmarshal traces configs response body and generate a output_table out of it
type tracesConfigs struct {
	Config struct {
		Retention string `json:"retention" header:"retention"`
	} `json:"config"`
}

// printTracesConfigsListTable prints traces configs response body as output_table
func printTracesConfigsListTable(body []byte) error {
	var tracesConfigs tracesConfigs

	// unmarshal response body
	if err := json.Unmarshal(body, &tracesConfigs); err != nil {
		return err
	}

	// print the output_table
	output_table.PrintTable(tracesConfigs.Config)

	return nil
}

// TracesConfigsCmd represents the tracesConfigs command
var TracesConfigsCmd = &cobra.Command{
	Use:   "traces-configs",
	Short: "Get traces config.",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		// generate an url
		url := config2.GetBaseUrl() + "traces-configs"

		// get output flag
		outputType := config2.GetOutputType()

		// call the command
		body, err := runCommand(url, "traces configs", outputType)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		// print output_table output
		if body != nil && (outputType == "" || outputType == "wide") {
			if err := printTracesConfigsListTable(body); err != nil {
				cmd.SilenceUsage = true
				return err
			}
		}

		return nil
	},
}
