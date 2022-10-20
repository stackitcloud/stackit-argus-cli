/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package update

import (
	"fmt"
	"github.com/spf13/cobra"
)

// AlertRulesCmd represents the alertRules command
var AlertRulesCmd = &cobra.Command{
	Use:   "alertRules <groupName> <alertName>",
	Short: "Update alert rules.",
	Long:  "Patch alert rules if alert name was not specified, otherwise update an alert rule.",
	Args:  cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			fmt.Println("patch alert rules")
		} else if len(args) == 2 {
			fmt.Println("update an alert rule")
		}
	},
}

func init() {
}
