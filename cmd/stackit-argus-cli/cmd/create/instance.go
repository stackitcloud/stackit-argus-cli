/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package create

import (
	"fmt"

	"github.com/spf13/cobra"
)

// InstanceCmd represents the instance command
var InstanceCmd = &cobra.Command{
	Use:   "instance",
	Short: "Create new instance.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create new instance")
	},
}

func init() {
}
