package get

/*
 * Get http check.
 */

import (
	"encoding/json"
	"github.com/spf13/cobra"
	config2 "github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/output_table"
)

// httpCheck struct is used to unmarshal http check response body
type httpCheck struct {
	HttpChecks []struct {
		Url string `json:"url" header:"url"`
	} `json:"httpChecks"`
}

// printHttpCheckTable prints http checks as a output_table
func printHttpCheckTable(body []byte, outputType config2.OutputType) error {
	var hc httpCheck

	// unmarshal response body
	if err := json.Unmarshal(body, &hc); err != nil {
		return err
	}

	// print the output_table
	output_table.PrintTable(hc.HttpChecks, outputType)

	return nil
}

// HttpCheckCmd represents the HttpCheck command
var HttpCheckCmd = &cobra.Command{
	Use:   "http-check",
	Short: "Get all http checks configured.",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		// generate an url
		url := config2.GetBaseUrl() + "http-checks"

		// get output flag
		outputType := config2.GetOutputType()

		// call the command
		body, err := runCommand(url, "http check", outputType)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		// print output_table output
		if body != nil && outputType != "yaml" && outputType != "json" {
			if err := printHttpCheckTable(body, outputType); err != nil {
				cmd.SilenceUsage = true
				return err
			}
		}

		return nil
	},
}
