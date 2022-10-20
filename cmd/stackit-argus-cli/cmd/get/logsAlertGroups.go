/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package get

import (
	"fmt"

	"github.com/spf13/cobra"
)

// LogsAlertGroupsCmd represents the logsAlertGroups command
var LogsAlertGroupsCmd = &cobra.Command{
	Use:   "logsAlertGroups <groupName>",
	Short: "Get logs alert groups config.",
	Long:  "Get list of logs alert groups config if group name was not specified, otherwise get logs alert group config.",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("get logs alert groups config")
		} else if len(args) == 1 {
			fmt.Println("get logs alert group config")
		}
	},
}

func init() {
}
