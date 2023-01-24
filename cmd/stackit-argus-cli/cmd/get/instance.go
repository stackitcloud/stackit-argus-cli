package get

/*
 * Get instances.
 */

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/output"
	"github.com/stackitcloud/stackit-argus-cli/internal/config"
)

// instance is used to unmarshal instance response body
type instance struct {
	Id       string `json:"id" header:"id"`
	Status   string `json:"status" header:"status"`
	PlanName string `json:"planName" header:"plan name"`
	// wide output attributes
	PlanId string `json:"planId" header:"plan id"`
	Error  string `json:"error" header:"error"`

	IsUpdatable bool   `json:"isUpdatable"`
	ServiceName string `json:"serviceName"`
}

// instances is used to unmarshal instances list response body
type instances struct {
	Instances []struct {
		Id       string `json:"id" header:"id"`
		PlanName string `json:"planName" header:"plan name"`
		Instance string `json:"instance"`
		Name     string `json:"name" header:"name"`
		Status   string `json:"status" header:"status"`
		// wide output attributes
		Error       string `json:"error" header:"error"`
		ServiceName string `json:"serviceName"`
	} `json:"instances"`
}

// printInstanceTable prints instance as a output
func printInstanceTable(body []byte, outputType config.OutputType) error {
	var i instance

	// unmarshal response body
	if err := json.Unmarshal(body, &i); err != nil {
		return err
	}

	// print the output
	if outputType != "wide" && outputType != "wide-table" {
		table := output.RemoveColumnsFromTable(i, []string{"PlanId", "Error"})
		output.PrintTable(table, string(outputType))
	} else {
		output.PrintTable(i, string(outputType))
	}

	return nil
}

// printListInstancesListTable prints instances list as a output
func printListInstancesListTable(body []byte, outputType config.OutputType) error {
	var is instances

	// unmarshal response body
	if err := json.Unmarshal(body, &is); err != nil {
		return err
	}

	// print the output
	if outputType != "wide" && outputType != "wide-table" {
		var table []interface{}

		for _, i := range is.Instances {
			table = append(table, output.RemoveColumnsFromTable(i, []string{"Error"}))
		}
		output.PrintTable(table, string(outputType))
	} else {
		output.PrintTable(is.Instances, string(outputType))
	}

	return nil
}

// InstanceCmd represents the instance command
var InstanceCmd = &cobra.Command{
	Use:   "instances <instance-id>",
	Short: "Get project instance.",
	Long:  "Get list of all project's instances if instance id was not specified, otherwise get instance.",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
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
		body, err := runCommand(url, resource, outputType)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		// print output output
		if body != nil && outputType != "yaml" && outputType != "json" {
			if len(args) == 0 {
				if err := printListInstancesListTable(body, outputType); err != nil {
					cmd.SilenceUsage = true
					return err
				}
			} else {
				if err := printInstanceTable(body, outputType); err != nil {
					cmd.SilenceUsage = true
					return err
				}
			}
		}

		return nil
	},
}
