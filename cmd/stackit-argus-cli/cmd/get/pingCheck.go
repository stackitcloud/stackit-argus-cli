package get

/*
 * Get ping check.
 */

import (
	"encoding/json"
	"github.com/spf13/cobra"
	config2 "github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/config"
	output_table "github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/outputTable"
)

// pingCheck struct is used to unmarshal ping check response body
type pingCheck struct {
	PingChecks []struct {
		Url string `json:"url" header:"url"`
	} `json:"pingChecks"`
}

// printPingCheckTable prints ping checks as a outputTable
func printPingCheckTable(body []byte) error {
	var pingCheck pingCheck

	// unmarshal response body
	if err := json.Unmarshal(body, &pingCheck); err != nil {
		return err
	}

	// print the outputTable
	output_table.PrintTable(pingCheck.PingChecks)

	return nil
}

// PingCheckCmd represents the PingCheck command
var PingCheckCmd = &cobra.Command{
	Use:   "pingCheck",
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

		// print outputTable output
		if body != nil && (outputType == "" || outputType == "wide") {
			if err := printPingCheckTable(body); err != nil {
				cmd.SilenceUsage = true
				return err
			}
		}

		return nil
	},
}