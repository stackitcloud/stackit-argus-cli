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
	Long:  "Get list of alert records if no alert record was specified or specific alert record.",
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
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// alertRecordsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// alertRecordsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
