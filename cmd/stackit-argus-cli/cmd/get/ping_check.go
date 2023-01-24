package get

/*
 * Get ping check.
 */

import (
	"encoding/json"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/output"
	"github.com/stackitcloud/stackit-argus-cli/internal/config"
)

// pingCheck struct is used to unmarshal ping check response body
type pingCheck struct {
	PingChecks []struct {
		Url string `json:"url" header:"url"`
	} `json:"pingChecks"`
}

// printPingCheckTable prints ping checks as an output
func printPingCheckTable(body []byte, outputType config.OutputType) error {
	var pc pingCheck

	// unmarshal response body
	if err := json.Unmarshal(body, &pc); err != nil {
		return err
	}

	// print the output
	output.PrintTable(pc.PingChecks, string(outputType))

	return nil
}

// PingCheckCmd represents the PingCheck command
var PingCheckCmd = &cobra.Command{
	Use:   "ping-check",
	Short: "Get all ping checks configured.",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		// generate an url
		url := config.GetBaseUrl() + "ping-checks"

		// get output flag
		outputType := config.GetOutputType()

		// call the command
		body, err := runCommand(url, "ping check", outputType)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		// print output output
		if body != nil && outputType != "yaml" && outputType != "json" {
			if err := printPingCheckTable(body, outputType); err != nil {
				cmd.SilenceUsage = true
				return err
			}
		}

		return nil
	},
}
