/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package delete

import (
	"fmt"
	"github.com/spf13/cobra"
)

// AlertRecordsCmd represents the alertRecords command
var AlertRecordsCmd = &cobra.Command{
	Use:   "alertRecords <groupName> <alertRecord>",
	Short: "Delete alert records.",
	Long:  "Delete alert records if alert record was not specified, otherwise delete an alert record.",
	Args:  cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 2 {
			fmt.Println("delete alert records")
		} else if len(args) == 1 {
			fmt.Println("delete an alert record")
		}
	},
}

func init() {
}
