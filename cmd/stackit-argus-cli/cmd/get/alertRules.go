package get

/*
 * Get alert rules.
 */

import (
	"encoding/json"
	"fmt"
	"github.com/lensesio/tableprinter"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/utils"

	"github.com/spf13/cobra"
)

// data stores alert rule properties
type data struct {
	Alert       string            `json:"alert" header:"alert"`
	Expr        string            `json:"expr" header:"expr"`
	For         string            `json:"for" header:"for"`
	Labels      map[string]string `json:"labels" header:"labels"`
	Annotations map[string]string `json:"annotations" header:"annotations"`
}

// alertRules is used to unmarshal alert rules response body
type alertRules struct {
	Data []data `json:"data"`
}

// alertRule is used to unmarshal alert rule response body
type alertRule struct {
	Data data `json:"data"`
}

// printAlertRulesTable prints alert rules table
func printAlertRulesTable(body []byte, outputType config.OutputType) {
	var alertRules alertRules

	// unmarshal response body
	err := json.Unmarshal(body, &alertRules)
	cobra.CheckErr(err)

	// print the table
	if outputType == "wide" {
		utils.PrintTable(alertRules.Data)
	} else {
		var table []interface{}
		for _, data := range alertRules.Data {
			newTable := tableprinter.RemoveStructHeader(data, "Labels")
			newTable = tableprinter.RemoveStructHeader(newTable, "Annotations")
			table = append(table, newTable)
		}
		utils.PrintTable(table)
	}
}

// printAlertRuleTable prints alert rule table
func printAlertRuleTable(body []byte, outputType config.OutputType) {
	var alertRule alertRule

	// unmarshal response body
	err := json.Unmarshal(body, &alertRule)
	cobra.CheckErr(err)

	// print the table
	if outputType == "wide" {
		utils.PrintTable(alertRule.Data)
	} else {
		table := tableprinter.RemoveStructHeader(alertRule.Data, "Labels")
		table = tableprinter.RemoveStructHeader(table, "Annotations")
		utils.PrintTable(table)
	}
}

// AlertRulesCmd represents the alertRules command
var AlertRulesCmd = &cobra.Command{
	Use:   "alertRules <groupName> <alertName>",
	Short: "Get alert rules.",
	Long:  "Get list of alert rules if alert name was not specified, otherwise get alert rule.",
	Args:  cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		// generate an url
		url := config.GetBaseUrl() + fmt.Sprintf("alertgroups/%s/alertrules", args[0])

		// modify url and debug message depend on arguments
		resource := "alert rules"
		if len(args) == 2 {
			resource = "alert rule"
			url += fmt.Sprintf("/%s", args[1])
		}

		// get output flag
		outputType := config.GetOutputType()

		// call the command
		body := runCommand(url, resource, outputType)

		// print table output
		if body != nil && (outputType == "" || outputType == "wide") {
			if len(args) == 1 {
				printAlertRulesTable(body, outputType)
			} else {
				printAlertRuleTable(body, outputType)
			}
		}
	},
}
