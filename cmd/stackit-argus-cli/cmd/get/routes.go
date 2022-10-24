/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package get

import (
	"fmt"

	"github.com/spf13/cobra"
)

// RoutesCmd represents the routes command
var RoutesCmd = &cobra.Command{
	Use:   "routes <receiver>",
	Short: "Get alert config routes.",
	Long:  "Get list of alert config route if receiver was not specified, otherwise get alert receiver for route.",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("get alert config route")
		} else if len(args) == 1 {
			fmt.Println("get alert receiver for route")
		}
	},
}

func init() {
}
