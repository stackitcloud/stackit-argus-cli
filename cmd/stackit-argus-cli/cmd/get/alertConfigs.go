/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package get

import (
	"fmt"

	"github.com/spf13/cobra"
)

// AlertConfigsCmd represents the alertConfigs command
var AlertConfigsCmd = &cobra.Command{
	Use:   "alertConfigs",
	Short: "Get alert configs.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("get alert configs called")
	},
}

func init() {

}
