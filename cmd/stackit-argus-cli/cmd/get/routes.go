package get

/*
 * Get alert configs routes.
 */

import (
	"encoding/json"
	"fmt"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/utils"

	"github.com/spf13/cobra"
)

// route is used to unmarshal routes response body and generate a table out of it
type route struct {
	Receiver       string `json:"receiver" header:"receiver"`
	GroupWait      string `json:"groupWait" header:"groupWait"`
	GroupInterval  string `json:"groupInterval" header:"groupInterval"`
	RepeatInterval string `json:"repeatInterval" header:"repeatInterval"`
	Continue       bool   `json:"continue" header:"continue"`
	// wide table attributes
	GroupBy  []string          `json:"groupBy" header:"groupBy"`
	Match    map[string]string `json:"match" header:"match"`
	MatchRe  map[string]string `json:"matchRe" header:"matchRe"`
	Matchers []string          `json:"matchers" header:"matchers"`

	Routes []route `json:"routes"`
}

// routesList is used to unmarshal routes response body
type routesList struct {
	Data route `json:"data"`
}

// getAllRoutes transform recursion structure to a slice of structures
func getAllRoutes(routes []route, newRoutes *[]route) {
	for _, route := range routes {
		*newRoutes = append(*newRoutes, route)
		getAllRoutes(route.Routes, newRoutes)
	}
}

// printRoutesListTable prints routes response body as table
func printRoutesListTable(body []byte, outputType config.OutputType) {
	var routes routesList
	var table []route

	// unmarshal response body
	err := json.Unmarshal(body, &routes)
	cobra.CheckErr(err)

	getAllRoutes(routes.Data.Routes, &table)
	table = append(table, routes.Data)

	// print the table
	if outputType != "wide" {
		var newTable []interface{}

		for _, data := range table {
			newTable = append(newTable, utils.RemoveColumnsFromTable(data,
				[]string{"GroupBy", "Match", "MatchRe", "Matchers"}))
		}
		utils.PrintTable(newTable)
	} else {
		utils.PrintTable(table)
	}
}

// RoutesCmd represents the routes command
var RoutesCmd = &cobra.Command{
	Use:   "routes <receiver>",
	Short: "Get alert config routes.",
	Long:  "Get list of alert config route if receiver was not specified, otherwise get alert receiver for route.",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
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
		body := runCommand(url, resource, outputType)

		// print table output
		if body != nil && (outputType == "" || outputType == "wide") {
			printRoutesListTable(body, outputType)
		}
	},
}
