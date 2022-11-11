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
		Record      string            `json:"record" header:"alert"`
		Alert       string            `json:"alert" header:"for"`
		Expr        string            `json:"expr" header:"expr"`
		For         string            `json:"for" header:"labels"`
		Labels      map[string]string `json:"labels" header:"annotations"`
		Annotations map[string]string `json:"annotations" header:"annotations"`
	} `json:"data"`
}

// printAlertRulesTable prints alert rules table
func printAlertRulesTable(body []byte) {
	var alertRules alertRules

	// unmarshal response body
	err := json.Unmarshal(body, &alertRules)
	cobra.CheckErr(err)

	// print the table
	utils.PrintTable(alertRules.Data)
}

// alertRule is used to unmarshal alert rule response body
type alertRule struct {
	Data struct {
		Record      string            `json:"record" header:"alert"`
		Alert       string            `json:"alert" header:"for"`
		Expr        string            `json:"expr" header:"expr"`
		For         string            `json:"for" header:"labels"`
		Labels      map[string]string `json:"labels" header:"annotations"`
		Annotations map[string]string `json:"annotations" header:"annotations"`
	} `json:"data"`
}

// printAlertRuleTable prints alert rule table
func printAlertRuleTable(body []byte) {
	var alertRule alertRule

	// unmarshal response body
	err := json.Unmarshal(body, &alertRule)
	cobra.CheckErr(err)

	// print the table
	utils.PrintTable(alertRule.Data)
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
				printAlertRulesTable(body)
			} else {
				printAlertRuleTable(body)
			}
		}
	},
}
