package get

/*
 * Get credentials.
 */

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	config2 "github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/output_table"
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

// credentialsTable holds structure of credentials output_table
type credentialsTable struct {
	UserName string `header:"username"`
}

// printRemoteWriteLimitsTable prints remote write limits response body as output_table
func printRemoteWriteLimitsTable(body []byte, outputType config2.OutputType) error {
	var rwl remoteWriteLimit

	// unmarshal response body
	if err := json.Unmarshal(body, &rwl); err != nil {
		return err
	}

	// print the output_table
	output_table.PrintTable(remoteWriteLimits, outputType)

	return nil
}

// printCredentialsListTable prints credentials
func printCredentialsListTable(body []byte, outputType config2.OutputType) error {
	var c credentials
	var table []credentialsTable

	// unmarshal response body
	if err := json.Unmarshal(body, &c); err != nil {
		return err
	}

	// fill output_table with values
	for _, data := range c.Credentials {
		if data.CredentialsInfo != nil {
			table = append(table, credentialsTable{
				UserName: data.CredentialsInfo.Username,
			})
		} else {
			table = append(table, credentialsTable{})
		}
	}

	// print the output_table
	output_table.PrintTable(table, outputType)

	return nil
}

// CredentialsCmd represents the credentials command
var CredentialsCmd = &cobra.Command{
	Use:   "credentials",
	Short: "Get technical user credentials.",
	Long:  "Get technical credentials.",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		// generate an url
		url := config2.GetBaseUrl() + "credentials"

		// modify url and debug message depend on argument and flag
		resource := "credentials"
		if remoteWriteLimits != "" {
			resource = "write limits credentials"
			url += fmt.Sprintf("/%s/remote-write-limits", remoteWriteLimits)
		}

		// get output flag
		outputType := config2.GetOutputType()

		// call the command
		body, err := runCommand(url, resource, outputType)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		// print output_table output
		if body != nil && outputType != "yaml" && outputType != "json" {
			if remoteWriteLimits != "" {
				if err := printRemoteWriteLimitsTable(body, outputType); err != nil {
					cmd.SilenceUsage = true
					return err
				}
			} else {
				if err := printCredentialsListTable(body, outputType); err != nil {
					cmd.SilenceUsage = true
					return err
				}
			}
		}

		return nil
	},
}

// init flags for the command
func init() {
	CredentialsCmd.Flags().StringVarP(&remoteWriteLimits, "remote-write-limits", "r", "",
		"get remote write limits credentials, username must be provided")
}
