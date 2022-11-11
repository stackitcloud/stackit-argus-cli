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

// printRemoteWriteLimitsTable prints remote write limits response body as table
func printRemoteWriteLimitsTable(body []byte) {
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

// printCredentialTable prints credential's info
func printCredentialTable(body []byte) {
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

// printCredentialsListTable prints credentials
func printCredentialsListTable(body []byte) {
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
		// generate an url
		url := config.GetBaseUrl() + "credentials"

		// modify url and debug message depend on argument and flag
		resource := "credentials"
		if remoteWriteLimits == true {
			if len(args) == 1 {
				resource = "write limits credentials"
				url += fmt.Sprintf("/%s/remote-write-limits", args[0])
			} else {
				cobra.CheckErr("get credentials with remote write limits must have one argument")
			}
		} else {
			if len(args) == 1 {
				resource = "credential"
				url += fmt.Sprintf("/%s", args[0])
			}
		}

		// get output flag
		outputType := config.GetOutputType()

		// call the command
		body := runCommand(url, resource, outputType)

		// print table output
		if body != nil && (outputType == "" || outputType == "wide") {
			if remoteWriteLimits == true {
				printRemoteWriteLimitsTable(body)
			} else if len(args) == 1 {
				printCredentialsListTable(body)
			} else {
				printCredentialTable(body)
			}
		}
	},
}

// init flags for the command
func init() {
	CredentialsCmd.Flags().BoolVarP(&remoteWriteLimits, "remote-write-limits", "r", false, "get remote write limits credentials")
}
