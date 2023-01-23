package get

/*
 * Get grafana configs.
 */

import (
	"encoding/json"
	"github.com/spf13/cobra"
	config2 "github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/output_table"
)

// grafanaConfigs is used to unmarshal grafana configs response body
type grafanaConfigs struct {
	PublicReadAccess bool `json:"publicReadAccess"`
	GenericOauth     *struct {
		Enabled             bool   `json:"enabled"`
		RoleAttributeStrict bool   `json:"roleAttributeStrict"`
		OauthClientId       string `json:"oauthClientId"`
		Scopes              string `json:"scopes"`
		RoleAttributePath   string `json:"roleAttributePath"`
		EmailAttributePath  string `json:"emailAttributePath"`
		LoginAttributePath  string `json:"loginAttributePath"`
		OauthClientSecret   string `json:"oauthClientSecret"`
		ApiUrl              string `json:"apiUrl"`
		AuthUrl             string `json:"authUrl"`
		TokenUrl            string `json:"tokenUrl"`
	} `json:"genericOauth"`
}

// grafanaConfigsTable holds structure of grafana configs output_table
type grafanaConfigsTable struct {
	OauthClientId       string `header:"oauth client id"`
	PublicReadAccess    bool   `header:"public read access"`
	Enabled             bool   `header:"generic auth enabled"`
	RoleAttributeStrict bool   `header:"role attribute strict"`
	// wide output_table attributes
	Scopes             string `header:"scopes"`
	RoleAttributePath  string `header:"role attribute path"`
	EmailAttributePath string `header:"email attribute path"`
	LoginAttributePath string `header:"login attribute path"`
}

// printGrafanaConfigsTable prints grafana configs response body as a output_table
func printGrafanaConfigsTable(body []byte, outputType config2.OutputType) error {
	var gc grafanaConfigs

	// unmarshal response body
	if err := json.Unmarshal(body, &gc); err != nil {
		return err
	}

	// generate a output_table
	wideTable := grafanaConfigsTable{}
	if gc.GenericOauth != nil {
		wideTable = grafanaConfigsTable{
			OauthClientId:       gc.GenericOauth.OauthClientId,
			PublicReadAccess:    gc.PublicReadAccess,
			Enabled:             gc.GenericOauth.Enabled,
			RoleAttributeStrict: gc.GenericOauth.RoleAttributeStrict,
			Scopes:              gc.GenericOauth.Scopes,
			RoleAttributePath:   gc.GenericOauth.RoleAttributePath,
			EmailAttributePath:  gc.GenericOauth.EmailAttributePath,
			LoginAttributePath:  gc.GenericOauth.LoginAttributePath,
		}
	}

	// print the output_table
	if outputType != "wide" && outputType != "wide-table" {
		//remove attributes that are not needed for default output_table
		table := output_table.RemoveColumnsFromTable(wideTable,
			[]string{"RoleAttributePath", "EmailAttributePath", "LoginAttributePath", "Scopes"})
		output_table.PrintTable(table, outputType)
	} else {
		output_table.PrintTable(wideTable, outputType)
	}

	return nil
}

// GrafanaConfigsCmd represents the grafanaConfigs command
var GrafanaConfigsCmd = &cobra.Command{
	Use:   "grafana-configs",
	Short: "Get grafana config.",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		// generate an url
		url := config2.GetBaseUrl() + "grafana-configs"

		// get output flag
		outputType := config2.GetOutputType()

		// call the command
		body, err := runCommand(url, "grafana configs", outputType)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		// print output_table output
		if body != nil && outputType != "yaml" && outputType != "json" {
			if err := printGrafanaConfigsTable(body, outputType); err != nil {
				cmd.SilenceUsage = true
				return err
			}
		}

		return nil
	},
}
