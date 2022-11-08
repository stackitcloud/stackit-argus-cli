package get

/*
 * Get alert rules.
 */

import (
	"encoding/json"
	"fmt"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/utils"

	"github.com/spf13/cobra"
)

// alertRules is used to unmarshal alert rules response body
type alertRules struct {
	Data []struct {
		Record      string            `json:"record"`
		Alert       string            `json:"alert"`
		Expr        string            `json:"expr"`
		For         string            `json:"for"`
		Labels      map[string]string `json:"labels"`
		Annotations map[string]string `json:"annotations"`
	} `json:"data"`
}

// alertRule is used to unmarshal alert rule response body
type alertRule struct {
	Data struct {
		Record      string            `json:"record"`
		Alert       string            `json:"alert"`
		Expr        string            `json:"expr"`
		For         string            `json:"for"`
		Labels      map[string]string `json:"labels"`
		Annotations map[string]string `json:"annotations"`
	} `json:"data"`
}

// alertRulesTable holds structure of alert rules table
type alertRulesTable struct {
	Alert       string            `header:"alert"`
	For         string            `header:"for"`
	Expr        string            `header:"expr"`
	Labels      map[string]string `header:"labels"`
	Annotations map[string]string `header:"annotations"`
}

// printAlertRules prints alert rules table
func printAlertRules(body []byte) {
	var alertRules alertRules
	var table []alertRulesTable

	// unmarshal response body
	err := json.Unmarshal(body, &alertRules)
	cobra.CheckErr(err)

	// fill table with values
	for _, data := range alertRules.Data {
		table = append(table, alertRulesTable{
			Alert:       data.Alert,
			For:         data.For,
			Expr:        data.Expr,
			Labels:      data.Labels,
			Annotations: data.Annotations,
		})
	}

	// print the table
	utils.PrintTable(table)
}

// printAlertRule prints alert rule table
func printAlertRule(body []byte) {
	var alertRule alertRule

	// unmarshal response body
	err := json.Unmarshal(body, &alertRule)
	cobra.CheckErr(err)

	// print the table
	utils.PrintTable(alertRulesTable{
		Alert:       alertRule.Data.Alert,
		For:         alertRule.Data.For,
		Expr:        alertRule.Data.Expr,
		Labels:      alertRule.Data.Labels,
		Annotations: alertRule.Data.Annotations,
	})
}

// AlertRulesCmd represents the alertRules command
var AlertRulesCmd = &cobra.Command{
	Use:   "alertRules <groupName> <alertName>",
	Short: "Get alert rules.",
	Long:  "Get list of alert rules if alert name was not specified, otherwise get alert rule.",
	Args:  cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		var debugMsg string

		// generate an url
		url := config.GetBaseUrl() + fmt.Sprintf("alertgroups/%s/alertrules", args[0])

		// modify url and debug message depend on arguments
		if len(args) == 2 {
			debugMsg = "get alert rule command called"
			url += fmt.Sprintf("/%s", args[1])
		} else if len(args) == 1 {
			debugMsg = "list alert rules command called"
		}

		// print debug messages if debug mode is turned on
		if config.IsDebugMode() {
			fmt.Println(debugMsg)
			fmt.Printf("url to call - %s\n", url)
		}

		// get alert rules
		status, body := getRequest(url)

		// print response status
		utils.ResponseMessage(status, "alert rules", "get")

		// print response body
		if status == 200 {
			outputType := config.GetOutputType()
			if outputType == "json" || outputType == "yaml" {
				utils.PrintYamlOrJson(body, string(outputType))
			} else if len(args) == 1 {
				printAlertRules(body)
			} else {
				printAlertRule(body)
			}
		}
	},
}
