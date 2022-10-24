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
	Use:   "alertGroups <groupName>",
	Short: "Get list of alert groups.",
	Long:  "Get list of alert groups if group name was not specified, otherwise get alert group.",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			fmt.Println("get alert group")
		} else if len(args) == 0 {
			fmt.Println("list alert groups")
		}
	},
}

func init() {

}
