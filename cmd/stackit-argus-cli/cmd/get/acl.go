/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package get

import (
	"fmt"

	"github.com/spf13/cobra"
)

// AclCmd represents the acl command
var AclCmd = &cobra.Command{
	Use:     "acl",
	Short:   "A brief description of your command",
	Long:    "",
	Example: "  get acl - to get acl",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("get acl called")
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// aclCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// aclCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
