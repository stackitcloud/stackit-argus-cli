/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package update

import (
	"fmt"

	"github.com/spf13/cobra"
)

// LogsConfigsCmd represents the logsConfigs command
var LogsConfigsCmd = &cobra.Command{
	Use:   "logsConfigs",
	Short: "Update logs config.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("update logs config")
	},
}

func init() {

}
