/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package update

import (
	"fmt"
	"github.com/spf13/cobra"
)

// AlertRecordsCmd represents the alertRecords command
var AlertRecordsCmd = &cobra.Command{
	Use:   "alertRecords <groupName> <alertRecord>",
	Short: "Update alert records.",
	Long:  "Patch alert records if alert record was not specified, otherwise update alert record.",
	Args:  cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			fmt.Println("patch alert records")
		} else if len(args) == 2 {
			fmt.Println("update alert record")
		}
	},
}

func init() {
}
