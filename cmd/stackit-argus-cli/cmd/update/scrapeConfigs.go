/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package update

import (
	"fmt"
	"github.com/spf13/cobra"
)

// ScrapeConfigsCmd represents the scrapeConfigs command
var ScrapeConfigsCmd = &cobra.Command{
	Use:   "scrapeConfigs <jobName>",
	Short: "Update scrape config.",
	Long:  "Patch scrape config if job name was nit specified, otherwise update scrape config.",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("patch scrape config")
		} else if len(args) == 1 {
			fmt.Println("update scrape config")
		}
	},
}

func init() {
}
