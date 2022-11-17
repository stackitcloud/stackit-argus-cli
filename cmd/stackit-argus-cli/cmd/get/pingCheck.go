package get

/*
 * Get ping check.
 */

import (
	"encoding/json"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	output_table "github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/outputTable"
)

// pingCheck struct is used to unmarshal ping check response body
type pingCheck struct {
	PingChecks []struct {
		Url string `json:"url" header:"url"`
	} `json:"pingChecks"`
}

// printPingCheckTable prints ping checks as a outputTable
func printPingCheckTable(body []byte) {
	var pingCheck pingCheck

	// unmarshal response body
	err := json.Unmarshal(body, &pingCheck)
	cobra.CheckErr(err)

	// print the outputTable
	output_table.PrintTable(pingCheck.PingChecks)
}

// PingCheckCmd represents the PingCheck command
var PingCheckCmd = &cobra.Command{
	Use:   "pingCheck",
	Short: "Get all ping checks configured.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// generate an url
		url := config.GetBaseUrl() + "ping-checks"

		// get output flag
		outputType := config.GetOutputType()

		// call the command
		body, err := runCommand(url, "ping check", outputType)
		cobra.CheckErr(err)

		// print outputTable output
		if body != nil && (outputType == "" || outputType == "wide") {
			printPingCheckTable(body)
		}
	},
}
