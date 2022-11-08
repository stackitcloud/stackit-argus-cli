package get

/*
 * Get alert groups.
 */

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/utils"
)

// alertGroup is used to unmarshal alert group response body
type alertGroup struct {
	Data struct {
		Name     string `json:"name"`
		Interval string `json:"interval"`
		Rules    []struct {
			Record      string            `json:"record"`
			Alert       string            `json:"alert"`
			Expr        string            `json:"expr"`
			For         string            `json:"for"`
			Labels      map[string]string `json:"labels"`
			Annotations map[string]string `json:"annotations"`
		} `json:"rules"`
	} `json:"data"`
}

// alertGroupTable holds structure of alert group table
type alertGroupTable struct {
	Record      string            `header:"record"`
	Alert       string            `header:"alert"`
	For         string            `header:"for"`
	Expr        string            `header:"expr"`
	Labels      map[string]string `header:"labels"`
	Annotations map[string]string `header:"annotations"`
}

// printAlertGroupResponse prints alert group response body as a table
func printAlertGroupResponse(body []byte) {
	var alertGroup alertGroup
	var table []alertGroupTable

	// unmarshal response body
	err := json.Unmarshal(body, &alertGroup)
	cobra.CheckErr(err)

	// fill table with values
	for _, rule := range alertGroup.Data.Rules {
		table = append(table, alertGroupTable{
			Alert:       rule.Alert,
			For:         rule.For,
			Expr:        rule.Expr,
			Labels:      rule.Labels,
			Annotations: rule.Annotations,
		})
	}

	// print group name and interval
	utils.PrintTable(alertGroupsTable{
		Name:     alertGroup.Data.Name,
		Interval: alertGroup.Data.Interval,
	})

	fmt.Println("RULES")

	// print the table
	utils.PrintTable(table)
}

// alertGroups is used to unmarshal alert groups response body
type alertGroups struct {
	Data []struct {
		Name     string `json:"name"`
		Interval string `json:"interval"`
	} `json:"data" yaml:"data"`
}

// alertGroupsTable holds structure of alert groups table
type alertGroupsTable struct {
	Name     string `header:"name"`
	Interval string `header:"interval"`
}

// printAlertGroupsResponse prints alert groups response body as a table
func printAlertGroupsResponse(body []byte) {
	var alertGroups alertGroups
	var table []alertGroupsTable

	// unmarshal response body
	err := json.Unmarshal(body, &alertGroups)
	cobra.CheckErr(err)

	// fill table with values
	for _, data := range alertGroups.Data {
		table = append(table, alertGroupsTable{
			Name:     data.Name,
			Interval: data.Interval,
		})
	}

	// print the table
	utils.PrintTable(table)
}

// AlertGroupsCmd represents the alertGroups command
var AlertGroupsCmd = &cobra.Command{
	Use:   "alertGroups <groupName>",
	Short: "Get list of alert groups.",
	Long:  "Get list of alert groups if group name was not specified, otherwise get alert group.",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var debugMsg string

		// generate an url
		url := config.GetBaseUrl() + "alertgroups"

		// modify url and debug message depend on arguments
		if len(args) == 1 {
			debugMsg = "get alert group command called"
			url += fmt.Sprintf("/%s", args[0])
		} else if len(args) == 0 {
			debugMsg = "list alert groups command called"
		}

		// print debug messages if debug mode is turned on
		if config.IsDebugMode() {
			fmt.Println(debugMsg)
			fmt.Printf("url to call - %s\n", url)
		}

		// get alert groups
		status, body := getRequest(url)

		// print response status
		utils.ResponseMessage(status, "alert group", "get")

		// print response body
		if status == 200 {
			outputType := config.GetOutputType()
			if outputType == "json" || outputType == "yaml" {
				utils.PrintYamlOrJson(body, string(outputType))
			} else if len(args) == 0 {
				printAlertGroupsResponse(body)
			} else {
				printAlertGroupResponse(body)
			}
		}
	},
}
