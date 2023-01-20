package get

/*
 * Get alert rules.
 */

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	config2 "github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/output_table"
)

// data stores alert rule properties
type data struct {
	Alert string `json:"alert" header:"alert"`
	Expr  string `json:"expr" header:"expr"`
	For   string `json:"for" header:"for"`
	// wide output_table attributes
	Labels      map[string]string `json:"labels" header:"labels"`
	Annotations map[string]string `json:"annotations" header:"annotations"`
}

// alertRules is used to unmarshal alert rules response body
type alertRules struct {
	Data []data `json:"data" validate:"required"`
}

// alertRule is used to unmarshal alert rule response body
type alertRule struct {
	Data data `json:"data" validate:"required"`
}

// printAlertRulesTable prints alert rules output_table
func printAlertRulesTable(body []byte, outputType config2.OutputType) error {
	var alertRules alertRules

	// unmarshal response body
	if err := json.Unmarshal(body, &alertRules); err != nil {
		return err
	}

	// print the output_table
	if outputType == "wide" {
		output_table.PrintTable(alertRules.Data)
	} else {
		var table []interface{}
		for _, data := range alertRules.Data {
			table = append(table, output_table.RemoveColumnsFromTable(data, []string{"Labels", "Annotations"}))
		}
		output_table.PrintTable(table)
	}

	return nil
}

// printAlertRuleTable prints alert rule output_table
func printAlertRuleTable(body []byte, outputType config2.OutputType) error {
	var alertRule alertRule

	// unmarshal response body
	if err := json.Unmarshal(body, &alertRule); err != nil {
		return err
	}

	// print the output_table
	if outputType == "wide" {
		output_table.PrintTable(alertRule.Data)
	} else {
		table := output_table.RemoveColumnsFromTable(alertRule.Data, []string{"Labels", "Annotations"})
		output_table.PrintTable(table)
	}

	return nil
}

// AlertRulesCmd represents the alertRules command
var AlertRulesCmd = &cobra.Command{
	Use:   "alert-rules <group-name> <alert-name>",
	Short: "Get alert rules.",
	Long:  "Get list of alert rules if alert name was not specified, otherwise get alert rule.",
	Args:  cobra.RangeArgs(1, 2),
	RunE: func(cmd *cobra.Command, args []string) error {
		// generate an url
		url := config2.GetBaseUrl() + fmt.Sprintf("alertgroups/%s/alertrules", args[0])

		// modify url and debug message depend on arguments
		resource := "alert rules"
		if len(args) == 2 {
			resource = "alert rule"
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

		// print output_table output
		if body != nil && (outputType == "" || outputType == "wide") {
			if len(args) == 1 {
				if err := printAlertRulesTable(body, outputType); err != nil {
					cmd.SilenceUsage = true
					return err
				}
			} else {
				if err := printAlertRuleTable(body, outputType); err != nil {
					cmd.SilenceUsage = true
					return err
				}
			}
		}

		return nil
	},
}
