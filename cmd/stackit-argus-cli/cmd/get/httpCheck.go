package get

/*
 * Get http check.
 */

import (
	"encoding/json"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/utils"

	"github.com/spf13/cobra"
)

// httpCheck struct is used to unmarshal http check response body
type httpCheck struct {
	HttpChecks []struct {
		Url string `json:"url" header:"url"`
	} `json:"httpChecks"`
}

// printHttpCheckTable prints http checks as a table
func printHttpCheckTable(body []byte) {
	var httpCheck httpCheck

	// unmarshal response body
	err := json.Unmarshal(body, &httpCheck)
	cobra.CheckErr(err)

	// print the table
	utils.PrintTable(httpCheck.HttpChecks)
}

// HttpCheckCmd represents the HttpCheck command
var HttpCheckCmd = &cobra.Command{
	Use:   "httpCheck",
	Short: "Get all http checks configured.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// generate an url
		url := config.GetBaseUrl() + "http-checks"

		// get output flag
		outputType := config.GetOutputType()

		// call the command
		body := runCommand(url, "http check", outputType)

		// print table output
		if body != nil && (outputType == "" || outputType == "wide") {
			printHttpCheckTable(body)
		}
	},
}
