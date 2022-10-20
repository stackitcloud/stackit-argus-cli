/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package get

import (
	"fmt"

	"github.com/spf13/cobra"
)

// LogsConfigsCmd represents the logsConfigs command
var LogsConfigsCmd = &cobra.Command{
	Use:   "logsConfigs",
	Short: "Get logs configuration.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("get logs configs called")
	},
}

func init() {
}
