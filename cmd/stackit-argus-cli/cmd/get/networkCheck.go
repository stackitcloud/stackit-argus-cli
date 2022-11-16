package get

/*
 * Get network check.
 */

import (
	"encoding/json"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	output_table "github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/outputTable"
)

// networkCheck struct is used to unmarshal network check response body
type networkCheck struct {
	NetworkChecks []struct {
		Address string `json:"address" header:"address"`
	} `json:"networkChecks"`
}

// printNetworkCheckTable prints network checks as a outputTable
func printNetworkCheckTable(body []byte) {
	var networkCheck networkCheck

	// unmarshal response body
	err := json.Unmarshal(body, &networkCheck)
	cobra.CheckErr(err)

	// print the outputTable
	output_table.PrintTable(networkCheck.NetworkChecks)
}

// NetworkCheckCmd represents the NetworkCheck command
var NetworkCheckCmd = &cobra.Command{
	Use:   "networkCheck",
	Short: "Get all network checks configured.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// generate an url
		url := config.GetBaseUrl() + "network-checks"

		// get output flag
		outputType := config.GetOutputType()

		// call the command
		body, err := runCommand(url, "network check", outputType)
		cobra.CheckErr(err)

		// print outputTable output
		if body != nil && (outputType == "" || outputType == "wide") {
			printNetworkCheckTable(body)
		}
	},
}
