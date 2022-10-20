/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package update

import (
	"fmt"

	"github.com/spf13/cobra"
)

// AclCmd represents the acl command
var AclCmd = &cobra.Command{
	Use:   "acl",
	Short: "Update acl config.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("update acl config")
	},
}

func init() {
}
