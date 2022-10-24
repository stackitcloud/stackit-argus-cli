/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package delete

import (
	"fmt"

	"github.com/spf13/cobra"
)

// InstanceCmd represents the instance command
var InstanceCmd = &cobra.Command{
	Use:   "instance <instanceId>",
	Short: "delete instance",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("delete instance")
	},
}

func init() {
}
