/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package get

import (
	"fmt"

	"github.com/spf13/cobra"
)

// RecordsCmd represents the alertRecords command
var RecordsCmd = &cobra.Command{
	Use:   "records <groupName> <alertRecord>",
	Short: "Get alert records.",
	Long:  "Get list of alert records if alert record was not specified, otherwise get alert record.",
	Args:  cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 2 {
			fmt.Println("get alert record")
		} else if len(args) == 1 {
			fmt.Println("get alert records")
		}
	},
}

func init() {
}
