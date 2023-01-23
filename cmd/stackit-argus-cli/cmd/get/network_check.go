package get

/*
 * Get network check.
 */

import (
	"encoding/json"
	"github.com/spf13/cobra"
	config2 "github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/output_table"
)

// networkCheck struct is used to unmarshal network check response body
type networkCheck struct {
	NetworkChecks []struct {
		Address string `json:"address" header:"address"`
	} `json:"networkChecks"`
}

// printNetworkCheckTable prints network checks as an output_table
func printNetworkCheckTable(body []byte, outputType config2.OutputType) error {
	var nc networkCheck

	// unmarshal response body
	if err := json.Unmarshal(body, &nc); err != nil {
		return err
	}

	// print the output_table
	output_table.PrintTable(nc.NetworkChecks, outputType)

	return nil
}

// NetworkCheckCmd represents the NetworkCheck command
var NetworkCheckCmd = &cobra.Command{
	Use:   "network-check",
	Short: "Get all network checks configured.",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		// generate an url
		url := config2.GetBaseUrl() + "network-checks"

		// get output flag
		outputType := config2.GetOutputType()

		// call the command
		body, err := runCommand(url, "network check", outputType)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		// print output_table output
		if body != nil && outputType != "yaml" && outputType != "json" {
			if err := printNetworkCheckTable(body, outputType); err != nil {
				cmd.SilenceUsage = true
				return err
			}
		}

		return nil
	},
}
