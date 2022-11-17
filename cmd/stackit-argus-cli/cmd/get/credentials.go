package get

/*
 * Get credentials.
 */

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	output_table "github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/outputTable"
)

var remoteWriteLimits string

// show only username and only credentials list

// credential structure is used to unmarshal credential response body
type remoteWriteLimit struct {
	MaxLimit            string `json:"maxLimit" header:"max limit"`
	CredentialsMaxLimit string `json:"credentialsMaxLimit" header:"credentials max limit"`
}

// credentials structure is used to unmarshal credentials response body
type credentials struct {
	Credentials []struct {
		Name            string `json:"name"`
		Id              string `json:"id"`
		CredentialsInfo *struct {
			Username string `json:"username"`
		} `json:"credentialsInfo"`
	} `json:"credentials" validate:"required"`
}

// credentialsTable holds structure of credentials outputTable
type credentialsTable struct {
	UserName string `header:"username"`
}

// printRemoteWriteLimitsTable prints remote write limits response body as outputTable
func printRemoteWriteLimitsTable(body []byte) {
	var remoteWriteLimit remoteWriteLimit

	// unmarshal response body
	err := json.Unmarshal(body, &remoteWriteLimit)
	cobra.CheckErr(err)

	// print the outputTable
	output_table.PrintTable(remoteWriteLimits)
}

// printCredentialsListTable prints credentials
func printCredentialsListTable(body []byte) {
	var credentials credentials
	var table []credentialsTable

	// unmarshal response body
	err := json.Unmarshal(body, &credentials)
	cobra.CheckErr(err)

	// fill outputTable with values
	for _, data := range credentials.Credentials {
		if data.CredentialsInfo != nil {
			table = append(table, credentialsTable{
				UserName: data.CredentialsInfo.Username,
			})
		} else {
			table = append(table, credentialsTable{})
		}
	}

	// print the outputTable
	output_table.PrintTable(table)
}

// CredentialsCmd represents the credentials command
var CredentialsCmd = &cobra.Command{
	Use:   "credentials <username>",
	Short: "Get technical user credentials.",
	Long:  "Get list of all credentials if username was not specified, otherwise get technical credentials.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// generate an url
		url := config.GetBaseUrl() + "credentials"

		// modify url and debug message depend on argument and flag
		resource := "credentials"
		if remoteWriteLimits != "" {
			resource = "write limits credentials"
			url += fmt.Sprintf("/%s/remote-write-limits", remoteWriteLimits)
		}

		// get output flag
		outputType := config.GetOutputType()

		// call the command
		body, err := runCommand(url, resource, outputType)
		cobra.CheckErr(err)

		// print outputTable output
		if body != nil && (outputType == "" || outputType == "wide") {
			if remoteWriteLimits != "" {
				printRemoteWriteLimitsTable(body)
			} else {
				printCredentialsListTable(body)
			}
		}
	},
}

// init flags for the command
func init() {
	CredentialsCmd.Flags().StringVarP(&remoteWriteLimits, "remote-write-limits", "r", "",
		"get remote write limits credentials, username must be provided")
}
