package get

/*
 * Get network check.
 */

import (
	"encoding/json"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/output"
	"github.com/stackitcloud/stackit-argus-cli/internal/config"
)

// networkCheck struct is used to unmarshal network check response body
type networkCheck struct {
	NetworkChecks []struct {
		Address string `json:"address" header:"address"`
	} `json:"networkChecks"`
}

// printNetworkCheckTable prints network checks as an output
func printNetworkCheckTable(body []byte, outputType config.OutputType) error {
	var nc networkCheck

	// unmarshal response body
	if err := json.Unmarshal(body, &nc); err != nil {
		return err
	}

	// print the output
	output.PrintTable(nc.NetworkChecks, string(outputType))

	return nil
}

// NetworkCheckCmd represents the NetworkCheck command
var NetworkCheckCmd = &cobra.Command{
	Use:   "network-check",
	Short: "Get all network checks configured.",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		// generate an url
		url := config.GetBaseUrl() + "network-checks"

		// get output flag
		outputType := config.GetOutputType()

		// call the command
		body, err := runCommand(url, "network check", outputType)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		// print output output
		if body != nil && outputType != "yaml" && outputType != "json" {
			if err := printNetworkCheckTable(body, outputType); err != nil {
				cmd.SilenceUsage = true
				return err
			}
		}

		return nil
	},
}
