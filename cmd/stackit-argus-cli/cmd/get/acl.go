package get

/*
 * Get an acl.
 */

import (
	"encoding/json"
	"fmt"
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

// printAclResponse prints acl response body like as table
func printAclResponse(body []byte) {
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

		// print debug messages if debug mode is turned on
		if config.IsDebugMode() {
			fmt.Println("get acl command called")
			fmt.Printf("url to call - %s\n", url)
		}

		// get acl
		status, body := getRequest(url)

		// print response status
		utils.ResponseMessage(status, "acl", "get")

		// print response body
		if status == 200 {
			outputType := config.GetOutputType()
			if outputType == "json" || outputType == "yaml" {
				utils.PrintYamlOrJson(body, string(outputType))
			} else {
				printAclResponse(body)
			}
		}
	},
}
