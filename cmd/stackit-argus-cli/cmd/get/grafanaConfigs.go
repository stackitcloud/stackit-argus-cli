package get

/*
 * Get grafana configs.
 */

import (
	"encoding/json"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	output_table "github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/outputTable"
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

// grafanaConfigsTable holds structure of grafana configs outputTable
type grafanaConfigsTable struct {
	OauthClientId       string `header:"oauth client id"`
	PublicReadAccess    bool   `header:"public read access"`
	Enabled             bool   `header:"generic auth enabled"`
	RoleAttributeStrict bool   `header:"role attribute strict"`
	// wide outputTable attributes
	Scopes             string `header:"scopes"`
	RoleAttributePath  string `header:"role attribute path"`
	EmailAttributePath string `header:"email attribute path"`
	LoginAttributePath string `header:"login attribute path"`
}

// printGrafanaConfigsTable prints grafana configs response body as a outputTable
func printGrafanaConfigsTable(body []byte, outputType config.OutputType) error {
	var grafanaConfigs grafanaConfigs

	// unmarshal response body
	if err := json.Unmarshal(body, &grafanaConfigs); err != nil {
		return err
	}

	// generate a outputTable
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

	// print the outputTable
	if outputType != "wide" {
		//remove attributes that are not needed for default outputTable
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
		url := config.GetBaseUrl() + "grafana-configs"

		// get output flag
		outputType := config.GetOutputType()

		// call the command
		body, err := runCommand(url, "grafana configs", outputType)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		// print outputTable output
		if body != nil && (outputType == "" || outputType == "wide") {
			if err := printGrafanaConfigsTable(body, outputType); err != nil {
				cmd.SilenceUsage = true
				return err
			}
		}

		return nil
	},
}
