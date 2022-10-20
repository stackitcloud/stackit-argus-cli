/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package delete

import (
	"fmt"
	"github.com/spf13/cobra"
)

var remoteWriteLimits bool

// CredentialsCmd represents the credentials command
var CredentialsCmd = &cobra.Command{
	Use:   "credentials <username>",
	Short: "Delete credentials.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if remoteWriteLimits == true {
			fmt.Println("delete remote write config for credentials")
		} else {
			fmt.Println("delete technical credentials")
		}
	},
}

func init() {
	CredentialsCmd.Flags().BoolVarP(&remoteWriteLimits, "remote-write-limits", "r", false, "delete remote write config for credentials")
}
