package get

/*
 * Get plans.
 */

import (
	"encoding/json"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/output"
	"github.com/stackitcloud/stackit-argus-cli/internal/config"
)

// acl is used to unmarshal acl response body
type plans struct {
	Plans []struct {
		Id          string `json:"id" header:"id"`
		Description string `json:"description" header:"description"`
		Name        string `json:"name" header:"name"`
		// wide output attributes
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

// printPlansTable prints plans response body as output
func printPlansTable(body []byte, outputType config.OutputType) error {
	var p plans

	// unmarshal response body
	if err := json.Unmarshal(body, &p); err != nil {
		return err
	}

	// print the output
	if outputType != "wide" {
		var table []interface{}

		for _, plan := range p.Plans {
			table = append(table, output.RemoveColumnsFromTable(plan,
				[]string{"BucketSize", "AlertRules", "SamplesPerScrape", "TargetNumber", "Amount", "LogsAlert"}))
		}
		output.PrintTable(table, string(outputType))
	} else {
		output.PrintTable(p.Plans, string(outputType))
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
		url := config.GetProjectUrl() + "/plans"

		// get output flag
		outputType := config.GetOutputType()

		// call the command
		body, err := runCommand(url, "plans", outputType)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		// print output output
		if body != nil && outputType != "yaml" && outputType != "json" {
			if err := printPlansTable(body, outputType); err != nil {
				cmd.SilenceUsage = true
				return err
			}
		}

		return nil
	},
}
