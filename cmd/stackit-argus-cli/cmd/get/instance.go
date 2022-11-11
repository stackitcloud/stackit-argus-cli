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

// instance is used to unmarshal instance response body
type instance struct {
	Id       string `json:"id" header:"id"`
	PlanId   string `json:"planId" header:"plan id"`
	PlanName string `json:"planName" header:"plan name"`
	Status   string `json:"status" header:"status"`
}

// instances is used to unmarshal instances list response body
type instances struct {
	Instances []struct {
		Id          string `json:"id" header:"id"`
		PlanName    string `json:"planName" header:"plan name"`
		Instance    string `json:"instance"`
		Name        string `json:"name" header:"name"`
		Status      string `json:"status" header:"status"`
		Error       string `json:"error" header:"error"`
		ServiceName string `json:"serviceName"`
	} `json:"instances"`
}

// printInstanceTable prints instance as a table
func printInstanceTable(body []byte) {
	var instance instance

	// unmarshal response body
	err := json.Unmarshal(body, &instance)
	cobra.CheckErr(err)

	// print the table
	utils.PrintTable(instance)
}

// printListInstancesListTable prints instances list as a table
func printListInstancesListTable(body []byte) {
	var instances instances

	// unmarshal response body
	err := json.Unmarshal(body, &instances)
	cobra.CheckErr(err)

	// print the table
	utils.PrintTable(instances.Instances)
}

// InstanceCmd represents the instance command
var InstanceCmd = &cobra.Command{
	Use:   "instances <instanceId>",
	Short: "Get project instance.",
	Long:  "Get list of all project's instances if instance id was not specified, otherwise get instance.",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// generate an url
		url := config.GetInstancesUrl()

		// modify url and debug message depend on arguments
		resource := "instances"
		if len(args) == 1 {
			resource = "instance"
			url += fmt.Sprintf("/%s", args[0])
		}

		// get output flag
		outputType := config.GetOutputType()

		// call the command
		body := runCommand(url, resource, outputType)

		// print table output
		if body != nil && (outputType == "" || outputType == "wide") {
			if len(args) == 0 {
				printListInstancesListTable(body)
			} else {
				printInstanceTable(body)
			}
		}
	},
}
