/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package create

import (
	"fmt"

	"github.com/spf13/cobra"
)

// BackupRestoresCmd represents the backupRestores command
var BackupRestoresCmd = &cobra.Command{
	Use:   "backupRestores",
	Short: "Restore backup.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("restore backup")
	},
}

func init() {
}
