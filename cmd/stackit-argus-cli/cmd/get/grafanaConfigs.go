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
	var grafanaConfigs grafanaConfigs

	// unmarshal response body
	if err := json.Unmarshal(body, &grafanaConfigs); err != nil {
		return err
	}

	// generate a output_table
	wideTable := grafanaConfigsTable{}
	if grafanaConfigs.GenericOauth != nil {
		wideTable = grafanaConfigsTable{
			OauthClientId:       grafanaConfigs.GenericOauth.OauthClientId,
			PublicReadAccess:    grafanaConfigs.PublicReadAccess,
			Enabled:             grafanaConfigs.GenericOauth.Enabled,
			RoleAttributeStrict: grafanaConfigs.GenericOauth.RoleAttributeStrict,
			Scopes:              grafanaConfigs.GenericOauth.Scopes,
			RoleAttributePath:   grafanaConfigs.GenericOauth.RoleAttributePath,
			EmailAttributePath:  grafanaConfigs.GenericOauth.EmailAttributePath,
			LoginAttributePath:  grafanaConfigs.GenericOauth.LoginAttributePath,
		}
	}

	// print the output_table
	if outputType != "wide" {
		//remove attributes that are not needed for default output_table
		table := output_table.RemoveColumnsFromTable(wideTable,
			[]string{"RoleAttributePath", "EmailAttributePath", "LoginAttributePath", "Scopes"})
		output_table.PrintTable(table)
	} else {
		output_table.PrintTable(wideTable)
	}

	return nil
}

// GrafanaConfigsCmd represents the grafanaConfigs command
var GrafanaConfigsCmd = &cobra.Command{
	Use:   "grafanaConfigs",
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
		if body != nil && (outputType == "" || outputType == "wide") {
			if err := printGrafanaConfigsTable(body, outputType); err != nil {
				cmd.SilenceUsage = true
				return err
			}
		}

		return nil
	},
}
