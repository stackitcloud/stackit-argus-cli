package get

/*
 * Get an acl.
 */

import (
	"encoding/json"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/output"
	"github.com/stackitcloud/stackit-argus-cli/internal/config"
)

// acl is used to unmarshal acl response body
type acl struct {
	Acl []string `json:"acl"`
}

// aclTable holds structure of acl output
type aclTable struct {
	Acl string `header:"acl"`
}

// printAclTable prints acl response body as output
func printAclTable(body []byte, outputType config.OutputType) error {
	var a acl
	var table []aclTable

	// unmarshal response body
	if err := json.Unmarshal(body, &a); err != nil {
		return err
	}

	// fill output with values
	for _, i := range a.Acl {
		table = append(table, aclTable{i})
	}

	// print the output
	output.PrintTable(table, string(outputType))

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

		// print output output
		if body != nil && outputType != "yaml" && outputType != "json" {
			if err := printAclTable(body, outputType); err != nil {
				cmd.SilenceUsage = true
				return err
			}
		}

		return nil
	},
}
