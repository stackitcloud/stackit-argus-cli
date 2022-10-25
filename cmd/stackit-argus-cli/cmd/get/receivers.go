/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package get

import (
	"fmt"

	"github.com/spf13/cobra"
)

// ReceiversCmd represents the receivers command
var ReceiversCmd = &cobra.Command{
	Use:   "receivers <receiver>",
	Short: "Get alert config receivers.",
	Long:  "Get list of alert config receivers if receiver was not specified, otherwise get alert config receiver.",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			fmt.Println("get alert config receiver")
		} else if len(args) == 0 {
			fmt.Println("get alert config receivers list")
		}
	},
}

func init() {

}
