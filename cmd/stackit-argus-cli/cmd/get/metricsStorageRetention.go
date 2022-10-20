/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package get

import (
	"fmt"

	"github.com/spf13/cobra"
)

// MetricsStorageRetentionCmd represents the metricsStorageRetention command
var MetricsStorageRetentionCmd = &cobra.Command{
	Use:   "metricsStorageRetention",
	Short: "Get metric storage retention time.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("get metrics storage retention called")
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// metricsStorageRetentionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// metricsStorageRetentionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
