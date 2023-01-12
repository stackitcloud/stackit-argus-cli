package get

/*
 * Get an acl.
 */

import (
	"encoding/json"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	output_table "github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/outputTable"
)

// acl is used to unmarshal acl response body
type acl struct {
	Acl []string `json:"acl"`
}

// aclTable holds structure of acl outputTable
type aclTable struct {
	Acl string `header:"acl"`
}

// printAclTable prints acl response body as outputTable
func printAclTable(body []byte) error {
	var acl acl
	var table []aclTable

	// unmarshal response body
	if err := json.Unmarshal(body, &acl); err != nil {
		return err
	}

	// fill outputTable with values
	for _, a := range acl.Acl {
		table = append(table, aclTable{a})
	}

	// print the outputTable
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
		url := config.GetBaseUrl() + "acl"

		// get output flag
		outputType := config.GetOutputType()

		// call the command
		body, err := runCommand(url, "acl", outputType)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		// print outputTable output
		if body != nil && (outputType == "" || outputType == "wide") {
			if err := printAclTable(body); err != nil {
				cmd.SilenceUsage = true
				return err
			}
		}

		return nil
	},
}
