/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package get

import (
	"fmt"

	"github.com/spf13/cobra"
)

// AlertConfigsCmd represents the alertConfigs command
var AlertConfigsCmd = &cobra.Command{
	Use:     "alertConfigs",
	Short:   "A brief description of your command",
	Long:    "",
	Example: "  get alertConfigs - to list alert configs",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("get alertConfigs called")
		if len(args) != 0 {
			fmt.Println("get alertConfigs shouldn't have any arguments")
		}
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// alertConfigsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// alertConfigsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
