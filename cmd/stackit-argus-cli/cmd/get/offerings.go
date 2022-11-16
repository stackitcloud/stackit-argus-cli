package get

/*
 * Get offerings.
 */

import (
	"encoding/json"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	output_table "github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/outputTable"
)

// offerings struct is used to unmarshal offerings response body
type offerings struct {
	Name        string `json:"name" header:"name"`
	Description string `json:"description" header:"description"`
	// wide outputTable attributes
	DocumentationUrl string   `json:"documentationUrl" header:"documentation url"`
	Tags             []string `json:"tags" header:"tags"`
}

// printOfferingsTable prints offerings response body as outputTable
func printOfferingsTable(body []byte, outputType config.OutputType) {
	var offerings offerings

	// unmarshal response body
	err := json.Unmarshal(body, &offerings)
	cobra.CheckErr(err)

	// print the outputTable
	if outputType != "wide" {
		output_table.PrintTable(output_table.RemoveColumnsFromTable(offerings, []string{"DocumentationUrl", "Tags"}))
	} else {
		output_table.PrintTable(offerings)
	}
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
		body, err := runCommand(url, "offerings", outputType)
		cobra.CheckErr(err)

		// print outputTable output
		if body != nil && (outputType == "" || outputType == "wide") {
			printOfferingsTable(body, outputType)
		}
	},
}
