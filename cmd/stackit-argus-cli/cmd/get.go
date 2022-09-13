/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve information about an ARGUS resource",
	Long: `Retrieve information about an ARGUS resource by passing its project id and optionally its instance id

Examples:
- stackit-argus-cli get
- stackit-argus-cli get instance -instanceId`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("get called")
		//TODO: Call HandlerFunction which calls the API via the API Client

		resp, err := http.Get("https://api-dev.stackit.cloud/argus-service/v1/projects/" + os.Getenv(ProjectId) + "/instances")
		if err != nil {
			print("error is - ", err.Error())
			return
		}

		_, err = io.ReadAll(resp.Body) // _ should be body this is WIP
		if err != nil {
			print("error is - ", err.Error())
			return
		}

		// instanceList := &models.InstanceList{}
		// instanceListItem := &models.InstanceListItem{}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	getCmd.Flags().String("projectId", "", "your project ID")
}
