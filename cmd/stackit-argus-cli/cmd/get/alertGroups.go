/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package get

import (
	"fmt"

	"github.com/spf13/cobra"
)

// AlertGroupsCmd represents the alertGroups command
var AlertGroupsCmd = &cobra.Command{
	Use:     "alertGroups",
	Short:   "A brief description of your command",
	Long:    "",
	Example: "  get alertGroups - to list alert groups\n  get alertGroups <groupName> - to get an alert group",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("get alertGroups called")
		if len(args) == 1 {
			fmt.Println("get alert group")
		} else if len(args) == 0 {
			fmt.Println("list alert groups")
		} else {
			fmt.Println("wrong number of arguments")
		}
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// alertGroupsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// alertGroupsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
