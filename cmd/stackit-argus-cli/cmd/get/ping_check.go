package get

/*
 * Get ping check.
 */

import (
	"encoding/json"
	"github.com/spf13/cobra"
	config2 "github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/output_table"
)

// pingCheck struct is used to unmarshal ping check response body
type pingCheck struct {
	PingChecks []struct {
		Url string `json:"url" header:"url"`
	} `json:"pingChecks"`
}

// printPingCheckTable prints ping checks as an output_table
func printPingCheckTable(body []byte, outputType config2.OutputType) error {
	var pc pingCheck

	// unmarshal response body
	if err := json.Unmarshal(body, &pc); err != nil {
		return err
	}

	// print the output_table
	output_table.PrintTable(pc.PingChecks, outputType)

	return nil
}

// PingCheckCmd represents the PingCheck command
var PingCheckCmd = &cobra.Command{
	Use:   "ping-check",
	Short: "Get all ping checks configured.",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		// generate an url
		url := config2.GetBaseUrl() + "ping-checks"

		// get output flag
		outputType := config2.GetOutputType()

		// call the command
		body, err := runCommand(url, "ping check", outputType)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		// print output_table output
		if body != nil && outputType != "yaml" && outputType != "json" {
			if err := printPingCheckTable(body, outputType); err != nil {
				cmd.SilenceUsage = true
				return err
			}
		}

		return nil
	},
}
