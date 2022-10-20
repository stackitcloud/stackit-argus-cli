/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package delete

import (
	"fmt"
	"github.com/spf13/cobra"
)

// ScrapeConfigsCmd represents the scrapeConfigs command
var ScrapeConfigsCmd = &cobra.Command{
	Use:   "scrapeConfigs",
	Short: "Delete scrape configs.",
	Long:  "Delete scrape configs if job name was not specified, otherwise delete scrape config.",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("delete scrape configs")
		} else if len(args) == 1 {
			fmt.Println("delete scrape config")
		}
	},
}

func init() {
}
