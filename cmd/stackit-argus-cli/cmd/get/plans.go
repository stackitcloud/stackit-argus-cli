package get

/*
 * Get plans.
 */

import (
	"encoding/json"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/utils"

	"github.com/spf13/cobra"
)

// acl is used to unmarshal acl response body
type plans struct {
	Plans []struct {
		Id          string `json:"id" header:"id"`
		Description string `json:"description" header:"description"`
		Name        string `json:"name" header:"name"`
		// wide table attributes
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

// printPlansTable prints plans response body as table
func printPlansTable(body []byte, outputType config.OutputType) {
	var plans plans

	// unmarshal response body
	err := json.Unmarshal(body, &plans)
	cobra.CheckErr(err)

	// print the table
	if outputType != "wide" {
		var table []interface{}

		for _, plan := range plans.Plans {
			table = append(table, utils.RemoveColumnsFromTable(plan,
				[]string{"BucketSize", "AlertRules", "SamplesPerScrape", "TargetNumber", "Amount", "LogsAlert"}))
		}
		utils.PrintTable(table)
	} else {
		utils.PrintTable(plans.Plans)
	}
}

// PlansCmd represents the plans command
var PlansCmd = &cobra.Command{
	Use:   "plans",
	Short: "Get all plans.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// generate an url
		url := config.GetProjectUrl() + "/plans"

		// get output flag
		outputType := config.GetOutputType()

		// call the command
		body := runCommand(url, "plans", outputType)

		// print table output
		if body != nil && (outputType == "" || outputType == "wide") {
			printPlansTable(body, outputType)
		}
	},
}
