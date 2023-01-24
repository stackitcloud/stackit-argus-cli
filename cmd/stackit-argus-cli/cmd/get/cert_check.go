package get

/*
 * Get cert check.
 */

import (
	"encoding/json"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/output"
	"github.com/stackitcloud/stackit-argus-cli/internal/config"
)

// certCheck struct is used to unmarshal cert check response body
type certCheck struct {
	CertChecks []struct {
		Source string `json:"source" header:"source"`
	} `json:"certChecks"`
}

// printCertCheckTable prints cert checks as a output
func printCertCheckTable(body []byte, outputType config.OutputType) error {
	var cc certCheck

	// unmarshal response body
	if err := json.Unmarshal(body, &cc); err != nil {
		return err
	}

	// print the output
	output.PrintTable(cc.CertChecks, string(outputType))

	return nil
}

// CertCheckCmd represents the CertCheck command
var CertCheckCmd = &cobra.Command{
	Use:   "cert-check",
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

		// print output output
		if body != nil && outputType != "yaml" && outputType != "json" {
			if err := printCertCheckTable(body, outputType); err != nil {
				cmd.SilenceUsage = true
				return err
			}
		}

		return nil
	},
}
