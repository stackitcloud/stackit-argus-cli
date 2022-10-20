/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package get

import (
	"fmt"

	"github.com/spf13/cobra"
)

// AlertRecordsCmd represents the alertRecords command
var AlertRecordsCmd = &cobra.Command{
	Use:     "alertRecords",
	Short:   "A brief description of your command",
	Long:    "",
	Example: "  get alertRecords <groupName> - to list alert records\n  get alertRecords <groupName> <alertRecord> - to get alert record",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("get alertRecords called")
		if len(args) == 2 {
			fmt.Println("get alert record")
		} else if len(args) == 1 {
			fmt.Println("get alert records")
		} else {
			fmt.Println("wrongs number of arguments")
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
