/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package update

import (
	"fmt"

	"github.com/spf13/cobra"
)

// LogsAlertGroupsCmd represents the logsAlertGroups command
var LogsAlertGroupsCmd = &cobra.Command{
	Use:   "logsAlertGroups <groupName>",
	Short: "Update logs alert group config.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("update logs alert group config")
	},
}

func init() {
}
