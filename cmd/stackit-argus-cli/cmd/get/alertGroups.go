package get

/*
 * Get alert groups.
 */

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	output_table "github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/outputTable"
)

// show number of rules

// alertGroups is used to unmarshal alert groups response body
type alertGroups struct {
	Data []struct {
		Name     string     `json:"name"`
		Interval string     `json:"interval"`
		Rules    []struct{} `json:"rules"`
	} `json:"data" yaml:"data" validate:"required"`
}

// alertGroup is used to unmarshal alert group response body
type alertGroup struct {
	Data struct {
		Name     string     `json:"name"`
		Interval string     `json:"interval"`
		Rules    []struct{} `json:"rules"`
	} `json:"data" validate:"required"`
}

// alertGroupTable holds structure of alert group outputTable
type alertGroupTable struct {
	Interval string `header:"interval"`
	Rules    int    `header:"rules"`
}

// alertGroupsTable holds structure of alert groups outputTable
type alertGroupsTable struct {
	Name     string `header:"name"`
	Interval string `header:"interval"`
	Rules    int    `header:"rules"`
}

// printAlertGroupTable prints alert group response body as a outputTable
func printAlertGroupTable(body []byte) {
	var alertGroup alertGroup

	// unmarshal response body
	err := json.Unmarshal(body, &alertGroup)
	cobra.CheckErr(err)

	// print the outputTable
	output_table.PrintTable(alertGroupTable{
		Interval: alertGroup.Data.Interval,
		Rules:    len(alertGroup.Data.Rules),
	})
}

// printAlertGroupsListTable prints alert groups response body as a outputTable
func printAlertGroupsListTable(body []byte) {
	var alertGroups alertGroups
	var table []alertGroupsTable

	// unmarshal response body
	err := json.Unmarshal(body, &alertGroups)
	cobra.CheckErr(err)

	// fill the outputTable with values
	for _, data := range alertGroups.Data {
		table = append(table, alertGroupsTable{
			Name:     data.Name,
			Interval: data.Interval,
			Rules:    len(data.Rules),
		})
	}

	// print the outputTable
	output_table.PrintTable(table)
}

// AlertGroupsCmd represents the alertGroups command
var AlertGroupsCmd = &cobra.Command{
	Use:   "alertGroups <groupName>",
	Short: "Get list of alert groups.",
	Long:  "Get list of alert groups if group name was not specified, otherwise get alert group.",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// generate an url
		url := config.GetBaseUrl() + "alertgroups"

		// modify url and resource depend on arguments
		resource := "alert groups"
		if len(args) == 1 {
			resource = "alert group"
			url += fmt.Sprintf("/%s", args[0])
		}

		// get output flag
		outputType := config.GetOutputType()

		// call the command
		body, err := runCommand(url, resource, outputType)
		cobra.CheckErr(err)

		// print outputTable output
		if body != nil && (outputType == "" || outputType == "wide") {
			if len(args) == 0 {
				printAlertGroupsListTable(body)
			} else {
				printAlertGroupTable(body)
			}
		}
	},
}
