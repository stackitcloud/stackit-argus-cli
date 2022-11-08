package get

/*
 * Get instances.
 */

import (
	"encoding/json"
	"fmt"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/utils"

	"github.com/spf13/cobra"
)

// alertRule is used to unmarshal instances list response body
type instances struct {
	Instances []struct {
		Id          string `json:"id"`
		PlanName    string `json:"planName"`
		Instance    string `json:"instance"`
		Name        string `json:"name"`
		Status      string `json:"status"`
		Error       string `json:"error"`
		ServiceName string `json:"serviceName"`
	} `json:"instances"`
}

// instancesTable holds structure of instances list table
type instancesTable struct {
	Id       string `header:"id"`
	PlanName string `header:"plan name"`
	Name     string `header:"name"`
	Status   string `header:"status"`
	Error    string `header:"error"`
}

// plan, id, error

// printListInstancesResponse prints instances list as a table
func printListInstancesResponse(body []byte) {
	var instances instances
	var table []instancesTable

	// unmarshal response body
	err := json.Unmarshal(body, &instances)
	cobra.CheckErr(err)

	// fill table with values
	for _, data := range instances.Instances {
		table = append(table, instancesTable{
			Id:       data.Id,
			PlanName: data.PlanName,
			Name:     data.Name,
			Status:   data.Status,
			Error:    data.Error,
		})
	}

	// print the table
	utils.PrintTable(table)
}

// InstanceCmd represents the instance command
var InstanceCmd = &cobra.Command{
	Use:   "instances <instanceId>",
	Short: "Get project instance.",
	Long:  "Get list of all project's instances if instance id was not specified, otherwise get instance.",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var debugMsg string

		// generate an url
		url := config.GetInstancesUrl()

		// modify url and debug message depend on arguments
		if len(args) == 0 {
			debugMsg = "list instances command called"
		} else if len(args) == 1 {
			debugMsg = "get instance command called"
			url += fmt.Sprintf("/%s", args[0])
		}

		// print debug messages if debug mode is turned on
		if config.IsDebugMode() {
			fmt.Println(debugMsg)
			fmt.Printf("url to call - %s\n", url)
		}

		// get instances
		status, body := getRequest(url)

		// print response status
		utils.ResponseMessage(status, "instances", "get")

		// print response body
		if status == 200 {
			outputType := config.GetOutputType()
			if outputType == "json" || outputType == "yaml" {
				utils.PrintYamlOrJson(body, string(outputType))
			} else if len(args) == 0 {
				printListInstancesResponse(body)
			} else {
				fmt.Println(body)
			}
		}
	},
}
