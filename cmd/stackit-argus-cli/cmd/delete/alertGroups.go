/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package delete

import (
	"fmt"
	"github.com/spf13/cobra"
)

// AlertGroupsCmd represents the alertGroups command
var AlertGroupsCmd = &cobra.Command{
	Use:   "alertGroups <groupName>",
	Short: "Delete alert groups.",
	Long:  "Delete alert groups if group name was not specified, otherwise delete alert group config.",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("delete alert groups")
		} else if len(args) == 1 {
			fmt.Println("delete alert group config")
		}
	},
}

func init() {
}
