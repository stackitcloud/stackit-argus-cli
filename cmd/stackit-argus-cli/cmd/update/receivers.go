/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package update

import (
	"fmt"

	"github.com/spf13/cobra"
)

// ReceiversCmd represents the receivers command
var ReceiversCmd = &cobra.Command{
	Use:   "receivers <receiver>",
	Short: "Update alert config receiver.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("update alert config receiver")
	},
}

func init() {
}
