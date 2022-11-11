package get

/*
 * Get alert records.
 */

import (
	"encoding/json"
	"fmt"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/utils"

	"github.com/spf13/cobra"
)

// recordsList is used to unmarshal records list response body and generate a table out of it
type recordsList struct {
	Data []struct {
		Record string            `json:"record" header:"record"`
		Expr   string            `json:"expr" header:"expr"`
		Labels map[string]string `json:"labels" header:"labels"`
	} `json:"data"`
}

// record is used to unmarshal record response body and generate a table out of it
type record struct {
	Data struct {
		Record string            `json:"record" header:"record"`
		Expr   string            `json:"expr" header:"expr"`
		Labels map[string]string `json:"labels" header:"labels"`
	} `json:"data"`
}

// printRecordsListTable prints records list response body as table
func printRecordsListTable(body []byte) {
	var records recordsList

	// unmarshal response body
	err := json.Unmarshal(body, &records)
	cobra.CheckErr(err)

	// print the table
	utils.PrintTable(records.Data)
}

// printRecordTable prints record response body as table
func printRecordTable(body []byte) {
	var record record

	// unmarshal response body
	err := json.Unmarshal(body, &record)
	cobra.CheckErr(err)

	// print the table
	utils.PrintTable(record.Data)
}

// RecordsCmd represents the alertRecords command
var RecordsCmd = &cobra.Command{
	Use:   "records <groupName> <alertRecord>",
	Short: "Get alert records.",
	Long:  "Get list of alert records if alert record was not specified, otherwise get alert record.",
	Args:  cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
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
		body := runCommand(url, resource, outputType)

		// print table output
		if body != nil && (outputType == "" || outputType == "wide") {
			if len(args) == 1 {
				printRecordsListTable(body)
			} else {
				printRecordTable(body)
			}
		}
	},
}
