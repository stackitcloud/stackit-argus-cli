package get

/*
 * Get alert configs.
 */

import (
	"encoding/json"
	"fmt"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/utils"

	"github.com/spf13/cobra"
)

// alertConfigs is used to unmarshal alert configs response body
type alertConfigs struct {
	Data struct {
		Global struct {
			SmtpSmarthost    string `json:"smtpSmarthost"`
			SmtpFrom         string `json:"smtpFrom"`
			SmtpAuthUsername string `json:"smtpAuthUsername"`
			SmtpAuthPassword string `json:"smtpAuthPassword"`
		} `json:"global"`
		Route        route
		InhibitRules []struct {
			SourceMatch struct {
				Severity string `json:"severity"`
			} `json:"sourceMatch"`
			TargetMatch struct {
				Severity string `json:"severity"`
			} `json:"targetMatch"`
			Equal []string `json:"equal"`
		}
		Receivers []struct{} `json:"receivers"`
	} `json:"data"`
}

// alertConfigsTable holds structure of alert configs table
type alertConfigsTable struct {
	SmtpSmarthost    string `header:"smtp smarthost"`
	SmtpFrom         string `header:"smtp from"`
	SmtpAuthUsername string `header:"smtp auth username"`
	SmtpAuthPassword string `header:"smtp auth password"`
	Receivers        int    `header:"receivers"`
	Routes           int    `header:"routes"`
	RouteReceiver    string `header:"route receiver"`
}

// alertConfigsWideTable holds structure of alert configs wide table
type alertConfigsWideTable struct {
	SourceMatchSeverity string   `header:"source match severity"`
	TargetMatchSeverity string   `header:"target match severity"`
	Equal               []string `header:"equal"`
}

// countRoutes counts number of routes
func countRoutes(routes []route, res *int) {
	for _, route := range routes {
		*res++
		countRoutes(route.Routes, res)
	}
}

// printAlertConfigsTable prints alert configs response body as table
func printAlertConfigsTable(body []byte, outputType config.OutputType) {
	var alertConfigs alertConfigs
	var routes int

	// unmarshal response body
	err := json.Unmarshal(body, &alertConfigs)
	cobra.CheckErr(err)

	countRoutes(alertConfigs.Data.Route.Routes, &routes)

	// print the table
	utils.PrintTable(alertConfigsTable{
		SmtpSmarthost:    alertConfigs.Data.Global.SmtpSmarthost,
		SmtpFrom:         alertConfigs.Data.Global.SmtpFrom,
		SmtpAuthUsername: alertConfigs.Data.Global.SmtpAuthUsername,
		SmtpAuthPassword: alertConfigs.Data.Global.SmtpAuthPassword,
		Receivers:        len(alertConfigs.Data.Receivers),
		Routes:           routes + 1,
		RouteReceiver:    alertConfigs.Data.Route.Receiver,
	})

	// print wide table
	if outputType == "wide" {
		var table []alertConfigsWideTable

		for _, data := range alertConfigs.Data.InhibitRules {
			table = append(table, alertConfigsWideTable{
				SourceMatchSeverity: data.SourceMatch.Severity,
				TargetMatchSeverity: data.TargetMatch.Severity,
				Equal:               data.Equal,
			})
		}

		if len(table) > 0 {
			fmt.Println("\nINHIBIT RULES")
			utils.PrintTable(table)
		}
	}
}

// AlertConfigsCmd represents the alertConfigs command
var AlertConfigsCmd = &cobra.Command{
	Use:   "alertConfigs",
	Short: "Get alert configs.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// generate an url
		url := config.GetBaseUrl() + "alertconfigs"

		// get output flag
		outputType := config.GetOutputType()

		// call the command
		body := runCommand(url, "alert configs", outputType)

		// print table output
		if body != nil && (outputType == "" || outputType == "wide") {
			printAlertConfigsTable(body, outputType)
		}
	},
}
