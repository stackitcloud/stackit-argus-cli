/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package update

import (
	"fmt"
	"github.com/spf13/cobra"
)

// AlertGroupsCmd represents the alertGroups command
var AlertGroupsCmd = &cobra.Command{
	Use:   "alertGroups <groupName>",
	Short: "Update alert groups.",
	Long:  "Patch alert groups if group name was not specified, otherwise update alert group config.",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("patch alert groups")
		} else if len(args) == 1 {
			fmt.Println("update alert group config")
		}
	},
}

func init() {
}
