/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package update

import (
	"fmt"

	"github.com/spf13/cobra"
)

// MetricsStorageRetentionCmd represents the metricsStorageRetention command
var MetricsStorageRetentionCmd = &cobra.Command{
	Use:   "metricsStorageRetention",
	Short: "Update metric update retention time.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("update metric update retention time")
	},
}

func init() {
}
