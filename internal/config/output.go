package config

/*
 * Init output type flag.
 */

import (
	"errors"
	"github.com/spf13/cobra"
)

type OutputType string

var flagOutputType OutputType

// String is used both by fmt.Print and by Cobra in help text
func (o *OutputType) String() string {
	return string(*o)
}

// Set must have pointer receiver, so it doesn't change the value of a copy
func (o *OutputType) Set(s string) error {
	switch s {
	case "json", "yaml", "wide", "table", "wide-table":
		*o = OutputType(s)
		return nil
	default:
		return errors.New("output type should be one of these: \"json\", \"yaml\", \"table\", \"wide\" or \"wide-table\"")
	}
}

// Type is only used in help text
func (o *OutputType) Type() string {
	return "string"
}

// outputFlagCompletion should probably live next to the OutputType definition
func outputFlagCompletion(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return []string{
		"json",
		"yaml",
		"wide",
		"table",
		"wide-table",
	}, cobra.ShellCompDirectiveDefault
}

// GetOutputType returns output type
func GetOutputType() OutputType {
	return flagOutputType
}
