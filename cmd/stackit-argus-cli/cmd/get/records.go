package get

/*
 * Get alert records.
 */

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	config2 "github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/config"
	output_table "github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/outputTable"
)

// recordsList is used to unmarshal records list response body and generate a outputTable out of it
type recordsList struct {
	Data []struct {
		Record string            `json:"record" header:"record"`
		Expr   string            `json:"expr" header:"expr"`
		Labels map[string]string `json:"labels" header:"labels"`
	} `json:"data" validate:"required"`
}

// record is used to unmarshal record response body and generate a outputTable out of it
type record struct {
	Data struct {
		Record string            `json:"record" header:"record"`
		Expr   string            `json:"expr" header:"expr"`
		Labels map[string]string `json:"labels" header:"labels"`
	} `json:"data" validate:"required"`
}

// printRecordsListTable prints records list response body as outputTable
func printRecordsListTable(body []byte, outputType config2.OutputType) error {
	var records recordsList

	// unmarshal response body
	if err := json.Unmarshal(body, &records); err != nil {
		return err
	}

	// print the outputTable
	if outputType != "wide" {
		var table []interface{}

		for _, d := range records.Data {
			table = append(table, output_table.RemoveColumnsFromTable(d, []string{"Labels"}))
		}
		output_table.PrintTable(table)
	} else {
		output_table.PrintTable(records.Data)
	}

	return nil
}

// printRecordTable prints record response body as outputTable
func printRecordTable(body []byte, outputType config2.OutputType) error {
	var record record

	// unmarshal response body
	if err := json.Unmarshal(body, &record); err != nil {
		return err
	}

	// print the outputTable
	if outputType != "wide" {
		output_table.PrintTable(output_table.RemoveColumnsFromTable(record.Data, []string{"Labels"}))
	} else {
		output_table.PrintTable(record.Data)
	}

	return nil
}

// RecordsCmd represents the alertRecords command
var RecordsCmd = &cobra.Command{
	Use:   "records <groupName> <alertRecord>",
	Short: "Get alert records.",
	Long:  "Get list of alert records if alert record was not specified, otherwise get alert record.",
	Args:  cobra.RangeArgs(1, 2),
	RunE: func(cmd *cobra.Command, args []string) error {
		// generate an url
		url := config2.GetBaseUrl() + fmt.Sprintf("alertgroups/%s/records", args[0])

		// modify url and debug message depend on arguments
		resource := "alert records"
		if len(args) == 2 {
			resource = "alert record"
			url += fmt.Sprintf("/%s", args[1])
		}

		// get output flag
		outputType := config2.GetOutputType()

		// call the command
		body, err := runCommand(url, resource, outputType)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		// print outputTable output
		if body != nil && (outputType == "" || outputType == "wide") {
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
