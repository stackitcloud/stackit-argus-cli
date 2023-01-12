package get

/*
 * Get http check.
 */

import (
	"encoding/json"
	"github.com/spf13/cobra"
	config2 "github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/config"
	output_table "github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/outputTable"
)

// httpCheck struct is used to unmarshal http check response body
type httpCheck struct {
	HttpChecks []struct {
		Url string `json:"url" header:"url"`
	} `json:"httpChecks"`
}

// printHttpCheckTable prints http checks as a outputTable
func printHttpCheckTable(body []byte) error {
	var httpCheck httpCheck

	// unmarshal response body
	if err := json.Unmarshal(body, &httpCheck); err != nil {
		return err
	}

	// print the outputTable
	output_table.PrintTable(httpCheck.HttpChecks)

	return nil
}

// HttpCheckCmd represents the HttpCheck command
var HttpCheckCmd = &cobra.Command{
	Use:   "httpCheck",
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

		// print outputTable output
		if body != nil && (outputType == "" || outputType == "wide") {
			if err := printHttpCheckTable(body); err != nil {
				cmd.SilenceUsage = true
				return err
			}
		}

		return nil
	},
}
