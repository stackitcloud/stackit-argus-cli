package get

/*
 * Get alert configs routes.
 */

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/output"
	"github.com/stackitcloud/stackit-argus-cli/internal/config"
)

// route is used to unmarshal routes response body and generate a output out of it
type route struct {
	Receiver       string `json:"receiver" header:"receiver" validate:"required"`
	GroupWait      string `json:"groupWait" header:"groupWait"`
	GroupInterval  string `json:"groupInterval" header:"groupInterval"`
	RepeatInterval string `json:"repeatInterval" header:"repeatInterval"`
	Continue       bool   `json:"continue" header:"continue"`
	// wide output attributes
	GroupBy  []string          `json:"groupBy" header:"groupBy"`
	Match    map[string]string `json:"match" header:"match"`
	MatchRe  map[string]string `json:"matchRe" header:"matchRe"`
	Matchers []string          `json:"matchers" header:"matchers"`

	Routes []route `json:"routes"`
}

// routesList is used to unmarshal routes response body
type routesList struct {
	Data route `json:"data" validate:"required"`
}

// getAllRoutes transform recursion structure to a slice of structures
func getAllRoutes(routes []route, newRoutes *[]route) {
	for _, route := range routes {
		*newRoutes = append(*newRoutes, route)
		getAllRoutes(route.Routes, newRoutes)
	}
}

// printRoutesListTable prints routes response body as output
func printRoutesListTable(body []byte, outputType config.OutputType) error {
	var routes routesList
	var table []route

	// unmarshal response body
	if err := json.Unmarshal(body, &routes); err != nil {
		return err
	}

	getAllRoutes(routes.Data.Routes, &table)
	table = append(table, routes.Data)

	// print the output
	if outputType != "wide" && outputType != "wide-table" {
		var newTable []interface{}

		for _, data := range table {
			newTable = append(newTable, output.RemoveColumnsFromTable(data,
				[]string{"GroupBy", "Match", "MatchRe", "Matchers"}))
		}
		output.PrintTable(newTable, string(outputType))
	} else {
		output.PrintTable(table, string(outputType))
	}

	return nil
}

// RoutesCmd represents the routes command
var RoutesCmd = &cobra.Command{
	Use:   "routes <receiver>",
	Short: "Get alert config routes.",
	Long:  "Get list of alert config route if receiver was not specified, otherwise get alert receiver for route.",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		// generate an url
		url := config.GetBaseUrl() + "alertconfigs/routes"

		// modify url and debug message depend on arguments
		resource := "alert config routes"
		if len(args) == 1 {
			resource = "alert config route"
			url += fmt.Sprintf("/%s", args[0])
		}

		// get output flag
		outputType := config.GetOutputType()

		// call the command
		body, err := runCommand(url, resource, outputType)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		// print output output
		if body != nil && outputType != "yaml" && outputType != "json" {
			if err := printRoutesListTable(body, outputType); err != nil {
				cmd.SilenceUsage = true
				return err
			}
		}

		return nil
	},
}
