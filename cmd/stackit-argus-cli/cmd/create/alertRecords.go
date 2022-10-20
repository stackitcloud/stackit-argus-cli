/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package create

import (
	"fmt"

	"github.com/spf13/cobra"
)

// AlertRecordsCmd represents the alertRecords command
var AlertRecordsCmd = &cobra.Command{
	Use:   "alertRecords <groupName>",
	Short: "Create alert record.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create alert record")
	},
}

func init() {

}
