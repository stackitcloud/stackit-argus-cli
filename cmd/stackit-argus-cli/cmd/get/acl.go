package get

/*
 * Get an acl.
 */

import (
	"encoding/json"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/utils"
)

// acl is used to unmarshal acl response body
type acl struct {
	Acl []string `json:"acl"`
}

// aclTable holds structure of acl table
type aclTable struct {
	Acl string `header:"acl"`
}

// printAclTable prints acl response body as table
func printAclTable(body []byte) {
	var acl acl
	var table []aclTable

	// unmarshal response body
	err := json.Unmarshal(body, &acl)
	cobra.CheckErr(err)

	// fill table with values
	for _, a := range acl.Acl {
		table = append(table, aclTable{a})
	}

	// print the table
	utils.PrintTable(table)
}

// AclCmd represents the acl command
var AclCmd = &cobra.Command{
	Use:   "acl",
	Short: "Get acl for the instance.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// generate an url
		url := config.GetBaseUrl() + "acl"

		// get output flag
		outputType := config.GetOutputType()

		// call the command
		body := runCommand(url, "acl", outputType)

		// print table output
		if body != nil && (outputType == "" || outputType == "wide") {
			printAclTable(body)
		}
	},
}
