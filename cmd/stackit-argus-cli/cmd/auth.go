/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"github.com/spf13/cobra"
	"io"
	"net/http"
)

type serviceAccount struct {
	id string `json:"id"`
}

type idToken struct {
	token      string `json:"token"`
	validUntil string `json:"validUntil"`
}

// authCmd represents the auth command
var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client := &http.Client{}

		// Create a service account via API. Use your own bearer token for creating it
		req, err := http.NewRequest("POST",
			"https://api-dev.stackit.cloud/service-account/v2/projects/"+args[0]+"/service-accounts",
			nil)
		if err != nil {
			print("error is - ", err.Error())
			return
		}
		req.Header.Set("Authorization", "")

		res, err := client.Do(req)
		if err != nil {
			print("error is - ", err.Error())
			return
		}

		if res.StatusCode != http.StatusCreated {
			res.Body.Close()
			println("Status is: ", res.Status)
			return
		}
		serviceAccount := &serviceAccount{}
		body, err := io.ReadAll(res.Body)
		if err != nil {
			print("error is - ", err.Error())
			return
		}
		res.Body.Close()
		if err = json.Unmarshal(body, &serviceAccount); err != nil {
			println("error is - ", err.Error())
		}

		// Add the service account to the project with a role
		req, err = http.NewRequest("POST",
			"https://api-dev.stackit.cloud/membership/v1/projects/"+args[0]+"/roles/project.member/service-accounts",
			nil)
		if err != nil {
			print("error is - ", err.Error())
			return
		}
		req.Header.Set("Authorization", "")
		req.Form.Set("serviceAccountId", serviceAccount.id)
		res, err = client.Do(req)
		if err != nil {
			print("error is - ", err.Error())
			return
		}
		res.Body.Close()

		// Create an id token
		req, err = http.NewRequest("POST",
			"https://api-dev.stackit.cloud/service-account/v1/projects/"+args[0]+"/service-accounts/"+serviceAccount.id+"/access-tokens",
			nil)
		if err != nil {
			print("error is - ", err.Error())
			return
		}
		req.Header.Set("Authorization", "")
		req.Form.Set("ttlDays", "180")
		res, err = client.Do(req)
		if err != nil {
			print("error is - ", err.Error())
			return
		}
		if res.StatusCode != http.StatusCreated {
			res.Body.Close()
			println("Status is: ", res.Status)
			return
		}

		idToken := &idToken{}
		body, err = io.ReadAll(res.Body)
		if err != nil {
			print("error is - ", err.Error())
			return
		}
		res.Body.Close()
		if err = json.Unmarshal(body, &idToken); err != nil {
			println("error is - ", err.Error())
		}

		println("Here is your auth token:\n", idToken.token)
		println("It will expire on: ", idToken.validUntil)
	},
}

func init() {
	rootCmd.AddCommand(authCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// authCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// authCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
