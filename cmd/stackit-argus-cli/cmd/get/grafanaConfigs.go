package get

/*
 * Get grafana configs.
 */

import (
	"encoding/json"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/utils"

	"github.com/spf13/cobra"
)

// grafanaConfigs is used to unmarshal grafana configs response body
type grafanaConfigs struct {
	PublicReadAccess bool `json:"publicReadAccess"`
	GenericOauth     struct {
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

// grafanaConfigsTable holds structure of grafana configs table
type grafanaConfigsTable struct {
	OauthClientId       string `header:"oauth client id"`
	PublicReadAccess    bool   `header:"public read access"`
	Enabled             bool   `header:"generic auth enabled"`
	RoleAttributeStrict bool   `header:"role attribute strict"`
	// wide table attributes
	Scopes             string `header:"scopes"`
	RoleAttributePath  string `header:"role attribute path"`
	EmailAttributePath string `header:"email attribute path"`
	LoginAttributePath string `header:"login attribute path"`
}

// printGrafanaConfigsTable prints grafana configs response body as a table
func printGrafanaConfigsTable(body []byte, outputType config.OutputType) {
	var grafanaConfigs grafanaConfigs

	// unmarshal response body
	err := json.Unmarshal(body, &grafanaConfigs)
	cobra.CheckErr(err)

	// generate a table
	wideTable := grafanaConfigsTable{
		OauthClientId:       grafanaConfigs.GenericOauth.OauthClientId,
		PublicReadAccess:    grafanaConfigs.PublicReadAccess,
		Enabled:             grafanaConfigs.GenericOauth.Enabled,
		RoleAttributeStrict: grafanaConfigs.GenericOauth.RoleAttributeStrict,
		Scopes:              grafanaConfigs.GenericOauth.Scopes,
		RoleAttributePath:   grafanaConfigs.GenericOauth.RoleAttributePath,
		EmailAttributePath:  grafanaConfigs.GenericOauth.EmailAttributePath,
		LoginAttributePath:  grafanaConfigs.GenericOauth.LoginAttributePath,
	}

	// print the table
	if outputType != "wide" {
		//remove attributes that are not needed for default table
		table := utils.RemoveColumnsFromTable(wideTable,
			[]string{"RoleAttributePath", "EmailAttributePath", "LoginAttributePath", "Scopes"})
		utils.PrintTable(table)
	} else {
		utils.PrintTable(wideTable)
	}
}

// GrafanaConfigsCmd represents the grafanaConfigs command
var GrafanaConfigsCmd = &cobra.Command{
	Use:   "grafanaConfigs",
	Short: "Get grafana config.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// generate an url
		url := config.GetBaseUrl() + "grafana-configs"

		// get output flag
		outputType := config.GetOutputType()

		// call the command
		body := runCommand(url, "grafana configs", outputType)

		// print table output
		if body != nil && (outputType == "" || outputType == "wide") {
			printGrafanaConfigsTable(body, outputType)
		}
	},
}
