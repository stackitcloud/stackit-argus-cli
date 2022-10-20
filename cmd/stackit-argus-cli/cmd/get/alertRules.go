/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package get

import (
	"fmt"

	"github.com/spf13/cobra"
)

// AlertRulesCmd represents the alertRules command
var AlertRulesCmd = &cobra.Command{
	Use:     "alertRules",
	Short:   "A brief description of your command",
	Long:    "",
	Example: "  get alertRules <groupName> - to list alert rules\n  get alertRules <groupName> <alertName> - to get alert rule",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("get alertRules called")
		if len(args) == 2 {
			fmt.Println("get alert rule")
		} else if len(args) == 1 {
			fmt.Println("list alert rules")
		} else {
			fmt.Println("wrong number of arguments")
		}
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// alertRulesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// alertRulesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
