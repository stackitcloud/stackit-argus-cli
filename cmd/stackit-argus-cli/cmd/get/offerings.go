package get

/*
 * Get offerings.
 */

import (
	"encoding/json"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/utils"

	"github.com/spf13/cobra"
)

// offerings struct is used to unmarshal offerings response body
type offerings struct {
	Name             string   `json:"name" header:"name"`
	Description      string   `json:"description" header:"description"`
	DocumentationUrl string   `json:"documentationUrl" header:"documentationUrl"`
	Tags             []string `json:"tags" header:"tags"`
}

// printOfferingsTable prints offerings response body as table
func printOfferingsTable(body []byte) {
	var offerings offerings

	// unmarshal response body
	err := json.Unmarshal(body, &offerings)
	cobra.CheckErr(err)

	// print the table
	utils.PrintTable(offerings)
}

// OfferingsCmd represents the plansOfferings command
var OfferingsCmd = &cobra.Command{
	Use:   "offerings",
	Short: "Get all offerings.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// generate an url
		url := config.GetProjectUrl() + "/offerings"

		// get output flag
		outputType := config.GetOutputType()

		// call the command
		body := runCommand(url, "offerings", outputType)

		// print table output
		if body != nil && (outputType == "" || outputType == "wide") {
			printOfferingsTable(body)
		}
	},
}
