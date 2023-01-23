package get

/*
 * Get offerings.
 */

import (
	"encoding/json"
	"github.com/spf13/cobra"
	config2 "github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/output_table"
)

// offerings struct is used to unmarshal offerings response body
type offerings struct {
	Name        string `json:"name" header:"name"`
	Description string `json:"description" header:"description"`
	// wide output_table attributes
	DocumentationUrl string   `json:"documentationUrl" header:"documentation url"`
	Tags             []string `json:"tags" header:"tags"`
}

// printOfferingsTable prints offerings response body as output_table
func printOfferingsTable(body []byte, outputType config2.OutputType) error {
	var o offerings

	// unmarshal response body
	if err := json.Unmarshal(body, &o); err != nil {
		return err
	}

	// print the output_table
	if outputType != "wide" && outputType != "wide-table" {
		output_table.PrintTable(output_table.RemoveColumnsFromTable(o, []string{"DocumentationUrl", "Tags"}), outputType)
	} else {
		output_table.PrintTable(o, outputType)
	}

	return nil
}

// OfferingsCmd represents the plansOfferings command
var OfferingsCmd = &cobra.Command{
	Use:   "offerings",
	Short: "Get all offerings.",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		// generate an url
		url := config2.GetProjectUrl() + "/offerings"

		// get output flag
		outputType := config2.GetOutputType()

		// call the command
		body, err := runCommand(url, "offerings", outputType)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		// print output_table output
		if body != nil && outputType != "yaml" && outputType != "json" {
			if err := printOfferingsTable(body, outputType); err != nil {
				cmd.SilenceUsage = true
				return err
			}
		}

		return nil
	},
}
