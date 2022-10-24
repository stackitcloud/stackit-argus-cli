/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package delete

import (
	"fmt"
	"github.com/spf13/cobra"
)

// AlertRulesCmd represents the alertRules command
var AlertRulesCmd = &cobra.Command{
	Use:   "alertRules <groupName> <alertName>",
	Short: "Delete alert rules.",
	Long:  "Delete alert rules if alert name was not specified, otherwise delete an alert rule",
	Args:  cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 2 {
			fmt.Println("delete alert rule")
		} else if len(args) == 1 {
			fmt.Println("delete alert rules")
		}
	},
}

func init() {
}
