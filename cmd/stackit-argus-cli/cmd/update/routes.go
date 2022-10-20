/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package update

import (
	"fmt"

	"github.com/spf13/cobra"
)

// RoutesCmd represents the routes command
var RoutesCmd = &cobra.Command{
	Use:   "routes <receiver>",
	Short: "Update alert receiver for route.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("update alert receiver for route")
	},
}

func init() {
}
