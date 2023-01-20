package get

/*
 * Get alert groups.
 */

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	config2 "github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/output_table"
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

// alertGroupTable holds structure of alert group output_table
type alertGroupTable struct {
	Interval string `header:"interval"`
	Rules    int    `header:"rules"`
}

// alertGroupsTable holds structure of alert groups output_table
type alertGroupsTable struct {
	Name     string `header:"name"`
	Interval string `header:"interval"`
	Rules    int    `header:"rules"`
}

// printAlertGroupTable prints alert group response body as a output_table
func printAlertGroupTable(body []byte) error {
	var alertGroup alertGroup

	// unmarshal response body
	if err := json.Unmarshal(body, &alertGroup); err != nil {
		return err
	}

	// print the output_table
	output_table.PrintTable(alertGroupTable{
		Interval: alertGroup.Data.Interval,
		Rules:    len(alertGroup.Data.Rules),
	})

	return nil
}

// printAlertGroupsListTable prints alert groups response body as a output_table
func printAlertGroupsListTable(body []byte) error {
	var alertGroups alertGroups
	var table []alertGroupsTable

	// unmarshal response body
	if err := json.Unmarshal(body, &alertGroups); err != nil {
		return err
	}

	// fill the output_table with values
	for _, data := range alertGroups.Data {
		table = append(table, alertGroupsTable{
			Name:     data.Name,
			Interval: data.Interval,
			Rules:    len(data.Rules),
		})
	}

	// print the output_table
	output_table.PrintTable(table)

	return nil
}

// AlertGroupsCmd represents the alertGroups command
var AlertGroupsCmd = &cobra.Command{
	Use:   "alert-groups <group-name>",
	Short: "Get list of alert groups.",
	Long:  "Get list of alert groups if group name was not specified, otherwise get alert group.",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		// generate an url
		url := config2.GetBaseUrl() + "alertgroups"

		// modify url and resource depend on arguments
		resource := "alert groups"
		if len(args) == 1 {
			resource = "alert group"
			url += fmt.Sprintf("/%s", args[0])
		}

		// get output flag
		outputType := config2.GetOutputType()

		// call the command
		body, err := runCommand(url, resource, outputType)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		// print output_table output
		if body != nil && (outputType == "" || outputType == "wide") {
			if len(args) == 0 {
				if err := printAlertGroupsListTable(body); err != nil {
					cmd.SilenceUsage = true
					return err
				}
			} else {
				if err := printAlertGroupTable(body); err != nil {
					cmd.SilenceUsage = true
					return err
				}
			}
		}

		return nil
	},
}
