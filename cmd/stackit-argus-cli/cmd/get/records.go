package get

/*
 * Get alert records.
 */

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/output"
	"github.com/stackitcloud/stackit-argus-cli/internal/config"
)

// recordsList is used to unmarshal records list response body and generate a output out of it
type recordsList struct {
	Data []struct {
		Record string            `json:"record" header:"record"`
		Expr   string            `json:"expr" header:"expr"`
		Labels map[string]string `json:"labels" header:"labels"`
	} `json:"data" validate:"required"`
}

// record is used to unmarshal record response body and generate a output out of it
type record struct {
	Data struct {
		Record string            `json:"record" header:"record"`
		Expr   string            `json:"expr" header:"expr"`
		Labels map[string]string `json:"labels" header:"labels"`
	} `json:"data" validate:"required"`
}

// printRecordsListTable prints records list response body as output
func printRecordsListTable(body []byte, outputType config.OutputType) error {
	var records recordsList

	// unmarshal response body
	if err := json.Unmarshal(body, &records); err != nil {
		return err
	}

	// print the output
	if outputType != "wide" {
		var table []interface{}

		for _, d := range records.Data {
			table = append(table, output.RemoveColumnsFromTable(d, []string{"Labels"}))
		}
		output.PrintTable(table, string(outputType))
	} else {
		output.PrintTable(records.Data, string(outputType))
	}

	return nil
}

// printRecordTable prints record response body as output
func printRecordTable(body []byte, outputType config.OutputType) error {
	var r record

	// unmarshal response body
	if err := json.Unmarshal(body, &r); err != nil {
		return err
	}

	// print the output
	if outputType != "wide" && outputType != "wide-table" {
		output.PrintTable(output.RemoveColumnsFromTable(r.Data, []string{"Labels"}), string(outputType))
	} else {
		output.PrintTable(r.Data, string(outputType))
	}

	return nil
}

// RecordsCmd represents the alertRecords command
var RecordsCmd = &cobra.Command{
	Use:   "records <group-name> <alert-record>",
	Short: "Get alert records.",
	Long:  "Get list of alert records if alert record was not specified, otherwise get alert record.",
	Args:  cobra.RangeArgs(1, 2),
	RunE: func(cmd *cobra.Command, args []string) error {
		// generate an url
		url := config.GetBaseUrl() + fmt.Sprintf("alertgroups/%s/records", args[0])

		// modify url and debug message depend on arguments
		resource := "alert records"
		if len(args) == 2 {
			resource = "alert record"
			url += fmt.Sprintf("/%s", args[1])
		}

		// get output flag
		outputType := config.GetOutputType()

		// call the command
		body, err := runCommand(url, resource, outputType)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		// print output output
		if body != nil && outputType != "yaml" && outputType != "json" {
			if len(args) == 1 {
				if err := printRecordsListTable(body, outputType); err != nil {
					cmd.SilenceUsage = true
					return err
				}
			} else {
				if err := printRecordTable(body, outputType); err != nil {
					cmd.SilenceUsage = true
					return err
				}
			}
		}

		return nil
	},
}
