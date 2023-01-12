package get

/*
 * Get cert check.
 */

import (
	"encoding/json"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	output_table "github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/outputTable"
)

// certCheck struct is used to unmarshal cert check response body
type certCheck struct {
	CertChecks []struct {
		Source string `json:"source" header:"source"`
	} `json:"certChecks"`
}

// printCertCheckTable prints cert checks as a outputTable
func printCertCheckTable(body []byte) error {
	var certCheck certCheck

	// unmarshal response body
	if err := json.Unmarshal(body, &certCheck); err != nil {
		return err
	}

	// print the outputTable
	output_table.PrintTable(certCheck.CertChecks)

	return nil
}

// CertCheckCmd represents the CertCheck command
var CertCheckCmd = &cobra.Command{
	Use:   "certCheck",
	Short: "Get all cert checks configured.",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		// generate an url
		url := config.GetBaseUrl() + "cert-checks"

		// get output flag
		outputType := config.GetOutputType()

		// call the command
		body, err := runCommand(url, "cert check", outputType)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		// print outputTable output
		if body != nil && (outputType == "" || outputType == "wide") {
			if err := printCertCheckTable(body); err != nil {
				cmd.SilenceUsage = true
				return err
			}
		}

		return nil
	},
}
