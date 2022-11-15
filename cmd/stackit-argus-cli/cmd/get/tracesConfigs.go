package get

/*
 * Get traces configs.
 */

import (
	"encoding/json"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/utils"

	"github.com/spf13/cobra"
)

// tracesConfigs is used to unmarshal traces configs response body and generate a table out of it
type tracesConfigs struct {
	Config struct {
		Retention string `json:"retention" header:"retention"`
	} `json:"config"`
}

// printTracesConfigsListTable prints traces configs response body as table
func printTracesConfigsListTable(body []byte) {
	var tracesConfigs tracesConfigs

	// unmarshal response body
	err := json.Unmarshal(body, &tracesConfigs)
	cobra.CheckErr(err)

	// print the table
	utils.PrintTable(tracesConfigs.Config)
}

// TracesConfigsCmd represents the tracesConfigs command
var TracesConfigsCmd = &cobra.Command{
	Use:   "tracesConfigs",
	Short: "Get traces config.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// generate an url
		url := config.GetBaseUrl() + "traces-configs"

		// get output flag
		outputType := config.GetOutputType()

		// call the command
		body := runCommand(url, "traces configs", outputType)

		// print table output
		if body != nil && (outputType == "" || outputType == "wide") {
			printTracesConfigsListTable(body)
		}
	},
}
