package get

/*
 * Get alert configs receivers.
 */

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/output"
	"github.com/stackitcloud/stackit-argus-cli/internal/config"
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
	} `json:"data" validate:"required"`
}

// webHooksConfigTable holds structure of web hook config output
type webHooksConfigTable struct {
	Url          string `header:"url"`
	MsTeams      bool   `header:"MS teams"`
	SendResolved bool   `header:"send resolved"`
}

// emailConfigTable holds structure of email config output
type emailConfigTable struct {
	To        string `header:"to"`
	From      string `header:"from"`
	SmartHost string `header:"smart host"`
}

// opsgenieConfigTable holds structure of opsgenie config output
type opsgenieConfigTable struct {
	ApiKey string `header:"api key"`
	ApiUrl string `header:"api url"`
	Tags   string `header:"tags"`
}

// printReceiverTable prints receiver response body as output
func printReceiverTable(body []byte, outputType config.OutputType) error {
	var r receiver

	// unmarshal response body
	if err := json.Unmarshal(body, &r); err != nil {
		return err
	}

	// print configs' output
	if len(r.Data.WebHookConfigs) > 0 {
		var table []webHooksConfigTable

		for _, data := range r.Data.WebHookConfigs {
			table = append(table, webHooksConfigTable{
				Url:          data.Url,
				MsTeams:      data.MsTeams,
				SendResolved: data.SendResolved,
			})

			fmt.Println("\nWEB HOOK CONFIGS")
			output.PrintTable(table, string(outputType))
		}
	}
	if len(r.Data.EmailConfigs) > 0 {
		var table []emailConfigTable

		for _, data := range r.Data.EmailConfigs {
			table = append(table, emailConfigTable{
				To:        data.To,
				From:      data.From,
				SmartHost: data.SmartHost,
			})

			fmt.Println("\nEMAIL CONFIGS")
			output.PrintTable(table, string(outputType))
		}
	}
	if len(r.Data.OpsgenieConfigs) > 0 {
		var table []opsgenieConfigTable

		for _, data := range r.Data.OpsgenieConfigs {
			table = append(table, opsgenieConfigTable{
				ApiKey: data.ApiKey,
				ApiUrl: data.ApiUrl,
				Tags:   data.Tags,
			})

			fmt.Println("\nOPSGENIE CONFIGS")
			output.PrintTable(table, string(outputType))
		}
	}

	return nil
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

// receiversListTable holds structure of receivers list output
type receiversListTable struct {
	Name            string `header:"name"`
	EmailConfigs    int    `header:"email configs"`
	WebHookConfigs  int    `header:"web hook configs"`
	OpsgenieConfigs int    `header:"opsgenie configs"`
}

// printReceiversListTable prints receivers list response body as output
func printReceiversListTable(body []byte, outputType config.OutputType) error {
	var rl receiversList
	var table []receiversListTable

	// unmarshal response body
	if err := json.Unmarshal(body, &rl); err != nil {
		return err
	}

	// fill the output with values
	for _, data := range rl.Data {
		table = append(table, receiversListTable{
			Name:            data.Name,
			EmailConfigs:    len(data.EmailConfigs),
			WebHookConfigs:  len(data.WebHookConfigs),
			OpsgenieConfigs: len(data.OpsgenieConfigs),
		})
	}

	// print the output
	output.PrintTable(table, string(outputType))

	return nil
}

// ReceiversCmd represents the receivers command
var ReceiversCmd = &cobra.Command{
	Use:   "receivers <receiver>",
	Short: "Get alert config receivers.",
	Long:  "Get list of alert config receivers if receiver was not specified, otherwise get alert config receiver.",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
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
		body, err := runCommand(url, resource, outputType)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		// print output output
		if body != nil && outputType != "yaml" && outputType != "json" {
			if len(args) == 0 {
				if err := printReceiversListTable(body, outputType); err != nil {
					cmd.SilenceUsage = true
					return err
				}
			} else {
				if err := printReceiverTable(body, outputType); err != nil {
					cmd.SilenceUsage = true
					return err
				}
			}
		}

		return nil
	},
}
