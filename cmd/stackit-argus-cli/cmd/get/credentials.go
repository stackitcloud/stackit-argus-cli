package get

/*
 * Get credentials.
 */

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/utils"
)

var remoteWriteLimits bool

// credential structure is used to unmarshal credential response body
type remoteWriteLimit struct {
	MaxLimit            string `json:"maxLimit" header:"max limit"`
	CredentialsMaxLimit string `json:"credentialsMaxLimit" header:"credentials max limit"`
}

func printRemoteWriteLimits(body []byte) {
	var remoteWriteLimit remoteWriteLimit

	// unmarshal response body
	err := json.Unmarshal(body, &remoteWriteLimit)
	cobra.CheckErr(err)

	// print the table
	utils.PrintTable(remoteWriteLimits)
}

// credential structure is used to unmarshal credential response body
type credential struct {
	Name string `json:"name" header:"name"`
	Id   string `json:"id" header:"id"`
}

// printCredential prints credential's info
func printCredential(body []byte) {
	var credential credential

	// unmarshal response body
	err := json.Unmarshal(body, &credential)
	cobra.CheckErr(err)

	// print the table
	utils.PrintTable(credential)
}

// credentials structure is used to unmarshal credentials response body
type credentials struct {
	Credentials []struct {
		Name            string `json:"name"`
		Id              string `json:"id"`
		CredentialsInfo struct {
			Username string `json:"username"`
		} `json:"credentialsInfo"`
	} `json:"credentials"`
}

// credentialsTable holds structure of credentials table
type credentialsTable struct {
	Name     string `header:"name"`
	Id       string `header:"id"`
	UserName string `header:"username"`
}

// printCredentials prints credentials
func printCredentials(body []byte) {
	var credentials credentials
	var table []credentialsTable

	// unmarshal response body
	err := json.Unmarshal(body, &credentials)
	cobra.CheckErr(err)

	// fill table with values
	for _, data := range credentials.Credentials {
		table = append(table, credentialsTable{
			Name:     data.Name,
			Id:       data.Id,
			UserName: data.CredentialsInfo.Username,
		})
	}

	// print the table
	utils.PrintTable(table)
}

// CredentialsCmd represents the credentials command
var CredentialsCmd = &cobra.Command{
	Use:   "credentials <username>",
	Short: "Get technical user credentials.",
	Long:  "Get list of all credentials if username was not specified, otherwise get technical credentials.",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var debugMsg string

		// generate an url
		url := config.GetBaseUrl() + "credentials"

		// modify url and debug message depend on argument and flag
		if remoteWriteLimits == true {
			if len(args) == 1 {
				debugMsg = "get remote write limits credentials command called"
				url += fmt.Sprintf("/%s/remote-write-limits", args[0])
			} else {
				cobra.CheckErr("get credentials with remote write limits must have one argument")
			}
		} else {
			if len(args) == 1 {
				debugMsg = "get credential command called"
				url += fmt.Sprintf("/%s", args[0])
			} else if len(args) == 0 {
				debugMsg = "list credentials command called"
			}
		}

		// print debug messages if debug mode is turned on
		if config.IsDebugMode() {
			fmt.Println(debugMsg)
			fmt.Printf("url to call - %s\n", url)
		}

		// get credentials
		status, body := getRequest(url)

		// print response status
		utils.ResponseMessage(status, "credentials", "get")

		// print response body
		if status == 200 {
			outputType := config.GetOutputType()
			if outputType == "json" || outputType == "yaml" {
				utils.PrintYamlOrJson(body, string(outputType))
			} else if len(args) == 0 {
				printCredentials(body)
			} else if remoteWriteLimits == true {
				printRemoteWriteLimits(body)
			} else {
				printCredential(body)
			}
		}
	},
}

// init flags for the command
func init() {
	CredentialsCmd.Flags().BoolVarP(&remoteWriteLimits, "remote-write-limits", "r", false, "get remote write limits credentials")
}
