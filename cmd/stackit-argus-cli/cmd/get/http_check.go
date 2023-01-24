package get

/*
 * Get http check.
 */

import (
	"encoding/json"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/output"
	"github.com/stackitcloud/stackit-argus-cli/internal/config"
)

// httpCheck struct is used to unmarshal http check response body
type httpCheck struct {
	HttpChecks []struct {
		Url string `json:"url" header:"url"`
	} `json:"httpChecks"`
}

// printHttpCheckTable prints http checks as a output
func printHttpCheckTable(body []byte, outputType config.OutputType) error {
	var hc httpCheck

	// unmarshal response body
	if err := json.Unmarshal(body, &hc); err != nil {
		return err
	}

	// print the output
	output.PrintTable(hc.HttpChecks, string(outputType))

	return nil
}

// HttpCheckCmd represents the HttpCheck command
var HttpCheckCmd = &cobra.Command{
	Use:   "http-check",
	Short: "Get all http checks configured.",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		// generate an url
		url := config.GetBaseUrl() + "http-checks"

		// get output flag
		outputType := config.GetOutputType()

		// call the command
		body, err := runCommand(url, "http check", outputType)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		// print output output
		if body != nil && outputType != "yaml" && outputType != "json" {
			if err := printHttpCheckTable(body, outputType); err != nil {
				cmd.SilenceUsage = true
				return err
			}
		}

		return nil
	},
}
