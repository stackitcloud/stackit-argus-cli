/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package update

import (
	"fmt"

	"github.com/spf13/cobra"
)

// AlertConfigsCmd represents the alertConfigs command
var AlertConfigsCmd = &cobra.Command{
	Use:   "alertConfigs",
	Short: "Update alert config.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("update alert config")
	},
}

func init() {
}
