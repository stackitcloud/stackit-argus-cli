package get

/*
 * Get alert configs receivers.
 */

import (
	"encoding/json"
	"fmt"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/utils"

	"github.com/spf13/cobra"
)

// receiver is used to unmarshal receiver response body
type receiver struct {
	Data struct {
		EmailConfigs []struct {
			To           string `json:"to"`
			From         string `json:"from"`
			SmartHost    string `json:"smartHost"`
			AuthUsername string `json:"authUsername"`
			AuthPassword string `json:"authPassword"`
			AuthIdentity string `json:"authIdentity"`
		} `json:"emailConfigs"`
		OpsgenieConfigs []struct {
			ApiKey string `json:"apiKey"`
			ApiUrl string `json:"apiUrl"`
			Tags   string `json:"tags"`
		} `json:"opsgenieConfigs"`
		WebHookConfigs []struct {
			Url          string `json:"url"`
			MsTeams      bool   `json:"msTeams"`
			SendResolved bool   `json:"send resolved"`
		} `json:"webHookConfigs"`
	} `json:"data"`
}

// webHooksConfigTable holds structure of web hook config table
type webHooksConfigTable struct {
	Url          string `header:"url"`
	MsTeams      bool   `header:"MS teams"`
	SendResolved bool   `header:"send resolved"`
}

// emailConfigTable holds structure of email config table
type emailConfigTable struct {
	To           string `header:"to"`
	From         string `header:"from"`
	SmartHost    string `header:"smart host"`
	AuthUsername string `header:"auth username"`
	AuthPassword string `header:"auth password"`
	AuthIdentity string `header:"auth identity"`
}

// opsgenieConfigTable holds structure of opsgenie config table
type opsgenieConfigTable struct {
	ApiKey string `header:"api key"`
	ApiUrl string `header:"api url"`
	Tags   string `header:"tags"`
}

// printReceiverTable prints receiver response body as table
func printReceiverTable(body []byte) {
	var receiver receiver

	// unmarshal response body
	err := json.Unmarshal(body, &receiver)
	cobra.CheckErr(err)

	// print configs' tables
	if len(receiver.Data.WebHookConfigs) > 0 {
		var table []webHooksConfigTable

		for _, data := range receiver.Data.WebHookConfigs {
			table = append(table, webHooksConfigTable{
				Url:          data.Url,
				MsTeams:      data.MsTeams,
				SendResolved: data.SendResolved,
			})

			fmt.Println("\nWEB HOOK CONFIGS")
			utils.PrintTable(table)
		}
	}
	if len(receiver.Data.EmailConfigs) > 0 {
		var table []emailConfigTable

		for _, data := range receiver.Data.EmailConfigs {
			table = append(table, emailConfigTable{
				To:           data.To,
				From:         data.From,
				SmartHost:    data.SmartHost,
				AuthUsername: data.AuthUsername,
				AuthPassword: data.AuthPassword,
				AuthIdentity: data.AuthIdentity,
			})

			fmt.Println("\nEMAIL CONFIGS")
			utils.PrintTable(table)
		}
	}
	if len(receiver.Data.OpsgenieConfigs) > 0 {
		var table []opsgenieConfigTable

		for _, data := range receiver.Data.OpsgenieConfigs {
			table = append(table, opsgenieConfigTable{
				ApiKey: data.ApiKey,
				ApiUrl: data.ApiUrl,
				Tags:   data.Tags,
			})

			fmt.Println("\nOPSGENIE CONFIGS")
			utils.PrintTable(table)
		}
	}
}

// receiversList is used to unmarshal receivers list response body
type receiversList struct {
	Data []struct {
		Name         string `json:"name"`
		EmailConfigs []struct {
		} `json:"emailConfigs"`
		OpsgenieConfigs []struct {
		} `json:"opsgenieConfigs"`
		WebHookConfigs []struct {
		} `json:"webHookConfigs"`
	} `json:"data"`
}

// receiversListTable holds structure of receivers list table
type receiversListTable struct {
	Name            string `header:"name"`
	EmailConfigs    int    `header:"email configs"`
	WebHookConfigs  int    `header:"web hook configs"`
	OpsgenieConfigs int    `header:"opsgenie configs"`
}

// printReceiversListTable prints receivers list response body as table
func printReceiversListTable(body []byte) {
	var receiversList receiversList
	var table []receiversListTable

	// unmarshal response body
	err := json.Unmarshal(body, &receiversList)
	cobra.CheckErr(err)

	// fill table with values
	for _, data := range receiversList.Data {
		table = append(table, receiversListTable{
			Name:            data.Name,
			EmailConfigs:    len(data.EmailConfigs),
			WebHookConfigs:  len(data.WebHookConfigs),
			OpsgenieConfigs: len(data.OpsgenieConfigs),
		})
	}

	// print the table
	utils.PrintTable(table)
}

// ReceiversCmd represents the receivers command
var ReceiversCmd = &cobra.Command{
	Use:   "receivers <receiver>",
	Short: "Get alert config receivers.",
	Long:  "Get list of alert config receivers if receiver was not specified, otherwise get alert config receiver.",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// generate an url
		url := config.GetBaseUrl() + "alertconfigs/receivers"

		// modify url and debug message depend on arguments
		resource := "alert config receivers"
		if len(args) == 1 {
			resource = "alert config receiver"
			url += fmt.Sprintf("/%s", args[0])
		}

		// get output flag
		outputType := config.GetOutputType()

		// call the command
		body := runCommand(url, resource, outputType)

		// print table output
		if body != nil && (outputType == "" || outputType == "wide") {
			if len(args) == 0 {
				printReceiversListTable(body)
			} else {
				printReceiverTable(body)
			}
		}
	},
}
