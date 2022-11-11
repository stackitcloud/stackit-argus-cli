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
		PlanId                  string  `json:"planId"`
		Id                      string  `json:"id" header:"id"`
		Description             string  `json:"description" header:"description"`
		Name                    string  `json:"name" header:"name"`
		BucketSize              int     `json:"bucketSize" header:"bucketSize"`
		GrafanaGlobalUsers      int     `json:"grafanaGlobalUsers"`
		GrafanaGlobalOrgs       int     `json:"grafanaGlobalOrgs"`
		GrafanaGlobalDashboards int     `json:"grafanaGlobalDashboards"`
		AlertRules              int     `json:"alertRules" header:"alert rules"`
		TargetNumber            int     `json:"targetNumber" header:"target number"`
		SamplesPerScrape        int     `json:"samplesPerScrape" header:"samples per scrape"`
		GrafanaGlobalSessions   int     `json:"grafanaGlobalSessions"`
		Amount                  float64 `json:"amount" header:"amount"`
		AlertReceivers          int     `json:"alertReceivers"`
		AlertMatchers           int     `json:"alertMatchers"`
		TracesStorage           int     `json:"tracesStorage"`
		LogsStorage             int     `json:"logsStorage"`
		LogsAlert               int     `json:"logsAlert"`
		IsFree                  bool    `json:"isFree" header:"free"`
		IsPublic                bool    `json:"isPublic" header:"public"`
	} `json:"plans"`
}

// printPlansTable prints plans response body as table
func printPlansTable(body []byte) {
	var plans plans

	// unmarshal response body
	err := json.Unmarshal(body, &plans)
	cobra.CheckErr(err)

	// print the table
	utils.PrintTable(plans.Plans)
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
			printPlansTable(body)
		}
	},
}
