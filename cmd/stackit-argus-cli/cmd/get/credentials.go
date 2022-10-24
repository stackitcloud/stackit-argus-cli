/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package get

import (
	"fmt"

	"github.com/spf13/cobra"
)

var remoteWriteLimits bool

// CredentialsCmd represents the credentials command
var CredentialsCmd = &cobra.Command{
	Use:   "credentials <username>",
	Short: "Get technical user credentials.",
	Long:  "Get list of all credentials if username was not specified, otherwise get technical credentials.",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if remoteWriteLimits == true {
			if len(args) == 1 {
				fmt.Println("get remote write limits credentials called")
			}
		} else {
			if len(args) == 1 {
				fmt.Println("get credential called")
			} else if len(args) == 0 {
				fmt.Println("get list of credentials called")
			}
		}
	},
}

func init() {
	CredentialsCmd.Flags().BoolVarP(&remoteWriteLimits, "remote-write-limits", "r", false, "get remote write limits credentials")
}
