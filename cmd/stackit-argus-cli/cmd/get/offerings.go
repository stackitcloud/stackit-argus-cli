/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package get

import (
	"fmt"

	"github.com/spf13/cobra"
)

// OfferingsCmd represents the plansOfferings command
var OfferingsCmd = &cobra.Command{
	Use:   "offerings",
	Short: "Get all offerings.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("get offerings called")
	},
}

func init() {
}
