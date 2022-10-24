/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package delete

import (
	"fmt"

	"github.com/spf13/cobra"
)

// RoutesCmd represents the routes command
var RoutesCmd = &cobra.Command{
	Use:   "routes <receiver>",
	Short: "Delete alert receiver for route.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("delete alert receiver for route")
	},
}

func init() {
}
