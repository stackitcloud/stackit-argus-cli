package get

/*
 * Get an acl.
 */

import (
	"encoding/json"
	"github.com/spf13/cobra"
	config2 "github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/output_table"
)

// acl is used to unmarshal acl response body
type acl struct {
	Acl []string `json:"acl"`
}

// aclTable holds structure of acl output_table
type aclTable struct {
	Acl string `header:"acl"`
}

// printAclTable prints acl response body as output_table
func printAclTable(body []byte) error {
	var acl acl
	var table []aclTable

	// unmarshal response body
	if err := json.Unmarshal(body, &acl); err != nil {
		return err
	}

	// fill output_table with values
	for _, a := range acl.Acl {
		table = append(table, aclTable{a})
	}

	// print the output_table
	output_table.PrintTable(table)

	return nil
}

// AclCmd represents the acl command
var AclCmd = &cobra.Command{
	Use:   "acl",
	Short: "Get acl for the instance.",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		// generate an url
		url := config2.GetBaseUrl() + "acl"

		// get output flag
		outputType := config2.GetOutputType()

		// call the command
		body, err := runCommand(url, "acl", outputType)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		// print output_table output
		if body != nil && (outputType == "" || outputType == "wide") {
			if err := printAclTable(body); err != nil {
				cmd.SilenceUsage = true
				return err
			}
		}

		return nil
	},
}
