/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/get"
)

type outputType string

const (
	outputJson outputType = "json"
	outputYaml outputType = "yaml"
	outputWide outputType = "wide"
)

func (o *outputType) String() string {
	return string(*o)
}

func (o *outputType) Set(s string) error {
	switch s {
	case "json", "yaml", "wide":
		*o = outputType(s)

		return nil
	default:
		return errors.New("output type should be one of these: \"json\", \"yaml\" or \"wide\"")
	}
}

func (o *outputType) Type() string {
	return "string"
}

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve information about an ARGUS resource.",
	Args:  cobra.NoArgs,
}

func init() {
	// add subcommands
	getCmd.AddCommand(get.InstanceCmd)
	getCmd.AddCommand(get.CredentialsCmd)
	getCmd.AddCommand(get.AlertConfigsCmd)
	getCmd.AddCommand(get.ReceiversCmd)
	getCmd.AddCommand(get.RoutesCmd)
	getCmd.AddCommand(get.AlertGroupsCmd)
	getCmd.AddCommand(get.AlertRulesCmd)
	getCmd.AddCommand(get.RecordsCmd)
	getCmd.AddCommand(get.BackupCmd)
	getCmd.AddCommand(get.BackupRetentionsCmd)
	getCmd.AddCommand(get.BackupSchedulesCmd)
	getCmd.AddCommand(get.GrafanaConfigsCmd)
	getCmd.AddCommand(get.LogsAlertGroupsCmd)
	getCmd.AddCommand(get.LogsConfigsCmd)
	getCmd.AddCommand(get.MetricsStorageRetentionCmd)
	getCmd.AddCommand(get.ScrapeConfigsCmd)
	getCmd.AddCommand(get.TracesConfigsCmd)
	getCmd.AddCommand(get.AclCmd)
	getCmd.AddCommand(get.PlansCmd)
	getCmd.AddCommand(get.OfferingsCmd)

	// define flags
	var flagOutputType = outputJson

	getCmd.PersistentFlags().VarP(&flagOutputType, "output", "o", "defines output format: yaml, json or wide")
	if err := getCmd.RegisterFlagCompletionFunc("output", outputFlagCompletion); err != nil {
		logger.Error("output flag autocompletion failed")

		return
	}
}

func outputFlagCompletion(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return []string{
		"json",
		"yaml",
		"wide",
	}, cobra.ShellCompDirectiveDefault
}
