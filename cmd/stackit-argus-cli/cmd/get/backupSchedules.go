/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package get

import (
	"fmt"

	"github.com/spf13/cobra"
)

// BackupSchedulesCmd represents the backupSchedules command
var BackupSchedulesCmd = &cobra.Command{
	Use:   "backupSchedules",
	Short: "Get backup schedules.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("get backup schedules called")
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// backupSchedulesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// backupSchedulesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
