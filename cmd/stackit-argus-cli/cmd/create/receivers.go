/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package create

import (
	"fmt"

	"github.com/spf13/cobra"
)

// ReceiversCmd represents the receivers command
var ReceiversCmd = &cobra.Command{
	Use:   "receivers",
	Short: "Create alert config receiver.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create alert config receiver")
	},
}

func init() {
}
