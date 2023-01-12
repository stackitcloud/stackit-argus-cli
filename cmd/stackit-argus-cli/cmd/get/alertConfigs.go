package get

/*
 * Get alert configs.
 */

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	output_table "github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/outputTable"
)

// alertConfigs is used to unmarshal alert configs response body
type alertConfigs struct {
	Data struct {
		Global *struct {
			SmtpSmarthost    string `json:"smtpSmarthost"`
			SmtpFrom         string `json:"smtpFrom"`
			SmtpAuthUsername string `json:"smtpAuthUsername"`
			SmtpAuthPassword string `json:"smtpAuthPassword"`
		} `json:"global" validate:"required"`
		Route        route
		InhibitRules []struct {
			SourceMatch *struct {
				Severity string `json:"severity"`
			} `json:"sourceMatch"`
			TargetMatch *struct {
				Severity string `json:"severity"`
			} `json:"targetMatch"`
			Equal []string `json:"equal"`
		}
		Receivers []struct{} `json:"receivers" validate:"required"`
	} `json:"data" validate:"required"`
}

// alertConfigsTable holds structure of alert configs outputTable
type alertConfigsTable struct {
	SmtpSmarthost string `header:"smtp smarthost"`
	SmtpFrom      string `header:"smtp from"`
	Receivers     int    `header:"receivers"`
	Routes        int    `header:"routes"`
}

// alertConfigsWideTable holds structure of alert configs wide outputTable
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

// printAlertConfigsTable prints alert configs response body as outputTable
func printAlertConfigsTable(body []byte, outputType config.OutputType) error {
	var alertConfigs alertConfigs
	var routes int

	// unmarshal response body
	if err := json.Unmarshal(body, &alertConfigs); err != nil {
		return err
	}

	countRoutes(alertConfigs.Data.Route.Routes, &routes)

	// print the outputTable
	table := alertConfigsTable{
		Receivers: len(alertConfigs.Data.Receivers),
		Routes:    routes + 1,
	}
	if alertConfigs.Data.Global != nil {
		table.SmtpFrom = alertConfigs.Data.Global.SmtpFrom
		table.SmtpSmarthost = alertConfigs.Data.Global.SmtpSmarthost
	}
	output_table.PrintTable(table)

	// print wide outputTable
	if outputType == "wide" {
		var table []alertConfigsWideTable

		for _, data := range alertConfigs.Data.InhibitRules {
			sourceMatchSeverity := ""
			targetMatchSeverity := ""
			if data.SourceMatch != nil {
				sourceMatchSeverity = data.SourceMatch.Severity
			}
			if data.TargetMatch != nil {
				targetMatchSeverity = data.TargetMatch.Severity
			}
			table = append(table, alertConfigsWideTable{
				SourceMatchSeverity: sourceMatchSeverity,
				TargetMatchSeverity: targetMatchSeverity,
				Equal:               data.Equal,
			})
		}

		if len(table) > 0 {
			fmt.Println("\nINHIBIT RULES")
			output_table.PrintTable(table)
		}
	}

	return nil
}

// AlertConfigsCmd represents the alertConfigs command
var AlertConfigsCmd = &cobra.Command{
	Use:   "alertConfigs",
	Short: "Get alert configs.",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		// generate an url
		url := config.GetBaseUrl() + "alertconfigs"

		// get output flag
		outputType := config.GetOutputType()

		// call the command
		body, err := runCommand(url, "alert configs", outputType)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		// print outputTable output
		if body != nil && (outputType == "" || outputType == "wide") {
			if err := printAlertConfigsTable(body, outputType); err != nil {
				cmd.SilenceUsage = true
				return err
			}
		}

		return nil
	},
}
