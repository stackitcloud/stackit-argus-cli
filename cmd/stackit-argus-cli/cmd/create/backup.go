/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package create

import (
	"fmt"

	"github.com/spf13/cobra"
)

// BackupCmd represents the backup command
var BackupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Create a backup.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create a backup")
	},
}

func init() {

}
