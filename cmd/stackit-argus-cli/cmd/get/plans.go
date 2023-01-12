package get

/*
 * Get plans.
 */

import (
	"encoding/json"
	"github.com/spf13/cobra"
	config2 "github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/config"
	output_table "github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/outputTable"
)

// acl is used to unmarshal acl response body
type plans struct {
	Plans []struct {
		Id          string `json:"id" header:"id"`
		Description string `json:"description" header:"description"`
		Name        string `json:"name" header:"name"`
		// wide outputTable attributes
		BucketSize    int `json:"bucketSize" header:"metrics storage(GB)"`
		AlertRules    int `json:"alertRules" header:"alert rules"`
		TargetNumber  int `json:"targetNumber" header:"target number"`
		LogsStorage   int `json:"logsStorage" header:"logs storage(GB)"`
		TracesStorage int `json:"tracesStorage" header:"traces storage(GB)"`

		SamplesPerScrape        int     `json:"samplesPerScrape"`
		GrafanaGlobalUsers      int     `json:"grafanaGlobalUsers"`
		GrafanaGlobalOrgs       int     `json:"grafanaGlobalOrgs"`
		GrafanaGlobalDashboards int     `json:"grafanaGlobalDashboards"`
		GrafanaGlobalSessions   int     `json:"grafanaGlobalSessions"`
		AlertReceivers          int     `json:"alertReceivers"`
		AlertMatchers           int     `json:"alertMatchers"`
		LogsAlert               int     `json:"logsAlert"`
		IsFree                  bool    `json:"isFree"`
		IsPublic                bool    `json:"isPublic"`
		Amount                  float64 `json:"amount"`
	} `json:"plans"`
}

// printPlansTable prints plans response body as outputTable
func printPlansTable(body []byte, outputType config2.OutputType) error {
	var plans plans

	// unmarshal response body
	if err := json.Unmarshal(body, &plans); err != nil {
		return err
	}

	// print the outputTable
	if outputType != "wide" {
		var table []interface{}

		for _, plan := range plans.Plans {
			table = append(table, output_table.RemoveColumnsFromTable(plan,
				[]string{"BucketSize", "AlertRules", "SamplesPerScrape", "TargetNumber", "Amount", "LogsAlert"}))
		}
		output_table.PrintTable(table)
	} else {
		output_table.PrintTable(plans.Plans)
	}

	return nil
}

// PlansCmd represents the plans command
var PlansCmd = &cobra.Command{
	Use:   "plans",
	Short: "Get all plans.",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		// generate an url
		url := config2.GetProjectUrl() + "/plans"

		// get output flag
		outputType := config2.GetOutputType()

		// call the command
		body, err := runCommand(url, "plans", outputType)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		// print outputTable output
		if body != nil && (outputType == "" || outputType == "wide") {
			if err := printPlansTable(body, outputType); err != nil {
				cmd.SilenceUsage = true
				return err
			}
		}

		return nil
	},
}
