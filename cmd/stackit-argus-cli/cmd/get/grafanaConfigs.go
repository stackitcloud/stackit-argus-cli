package get

/*
 * Get grafana configs.
 */

import (
	"encoding/json"
	"fmt"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/utils"

	"github.com/spf13/cobra"
)

// grafanaConfigs is used to unmarshal grafana configs response body
type grafanaConfigs struct {
	PublicReadAccess bool `json:"publicReadAccess"`
	GenericOauth     struct {
		Enabled             bool   `json:"enabled"`
		ApiUrl              string `json:"apiUrl"`
		AuthUrl             string `json:"authUrl"`
		Scopes              string `json:"scopes"`
		TokenUrl            string `json:"tokenUrl"`
		OauthClientId       string `json:"oauthClientId"`
		OauthClientSecret   string `json:"oauthClientSecret"`
		RoleAttributeStrict bool   `json:"roleAttributeStrict"`
		RoleAttributePath   string `json:"roleAttributePath"`
		LoginAttributePath  string `json:"loginAttributePath"`
	} `json:"genericOauth"`
}

// grafanaConfigsTable holds structure of grafana configs table
type grafanaConfigsTable struct {
	PublicReadAccess    bool   `header:"public read access"`
	Enabled             bool   `header:"generic auth enabled"`
	Scopes              string `header:"scopes"`
	RoleAttributeStrict bool   `header:"role attribute strict"`
}

// printGrafanaConfigsResponse prints grafana configs response body as a table
func printGrafanaConfigsResponse(body []byte) {
	var grafanaConfigs grafanaConfigs

	// unmarshal response body
	err := json.Unmarshal(body, &grafanaConfigs)
	cobra.CheckErr(err)

	// print the table
	utils.PrintTable(grafanaConfigsTable{
		PublicReadAccess:    grafanaConfigs.PublicReadAccess,
		Enabled:             grafanaConfigs.GenericOauth.Enabled,
		Scopes:              grafanaConfigs.GenericOauth.Scopes,
		RoleAttributeStrict: grafanaConfigs.GenericOauth.RoleAttributeStrict,
	})
}

// GrafanaConfigsCmd represents the grafanaConfigs command
var GrafanaConfigsCmd = &cobra.Command{
	Use:   "grafanaConfigs",
	Short: "Get grafana config.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// generate an url
		url := config.GetBaseUrl() + "grafana-configs"

		// print debug messages if debug mode is turned on
		if config.IsDebugMode() {
			fmt.Println("get grafana configs command called")
			fmt.Printf("url to call - %s\n", url)
		}

		// get grafana configs
		status, body := getRequest(url)

		// print response status
		utils.ResponseMessage(status, "grafana configs", "get")

		// print response body
		if status == 200 {
			outputType := config.GetOutputType()
			if outputType == "json" || outputType == "yaml" {
				utils.PrintYamlOrJson(body, string(outputType))
			} else {
				printGrafanaConfigsResponse(body)
			}
		}
	},
}
