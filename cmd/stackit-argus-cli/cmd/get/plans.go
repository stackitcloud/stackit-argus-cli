/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package get

import (
	"fmt"

	"github.com/spf13/cobra"
)

// PlansCmd represents the plans command
var PlansCmd = &cobra.Command{
	Use:   "plans",
	Short: "Get all plans.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("get plans called")
	},
}

func init() {
}
