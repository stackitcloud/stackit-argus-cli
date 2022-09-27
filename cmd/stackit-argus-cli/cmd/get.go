/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/viper"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/models"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

var listFlag bool
var rawFlag bool

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve information about an ARGUS resource",
	Long: `Retrieve information about an ARGUS resource by passing its project id and optionally its instance id

Examples:
- stackit-argus-cli get
- stackit-argus-cli get instance -instanceId`,
	Run: func(cmd *cobra.Command, args []string) {
		//viper.SetConfigName(".stackit-argus-cli")
		//viper.SetConfigType("yaml")
		//viper.AddConfigPath(".")
		err := viper.ReadInConfig()
		if err != nil {
			panic(fmt.Errorf("fatal error config file: %w", err))
		}

		var url string

		token := "Bearer " + viper.GetString("token")

		if viper.Get(loggedIn) == "" {
			pId := viper.GetString(ProjectId)
			url = "https://api-dev.stackit.cloud/argus-service/v1/projects/" + pId + "/instances"
		} else {
			url = "https://api-dev.stackit.cloud/argus-service/v1/projects/" + viper.GetString(loggedIn) + "/instances"
		}

		if listFlag == false {
			if len(args) > 0 {
				url = url + "/" + args[0]
			} else {
				fmt.Println("Please provide an instance ID or set the list flag")
				return
			}
		}

		fmt.Println("The request url is: " + url)

		instanceList := &models.InstanceList{}
		projectInstance := &models.ProjectInstance{}

		if listFlag == true {
			err = getJsonWithToken(url, token, instanceList)
			if err != nil {
				panic(fmt.Errorf("fatal error get request: %w", err))
			}
			spew.Dump(instanceList)
		} else {
			err = getJsonWithToken(url, token, projectInstance)
			if err != nil {
				panic(fmt.Errorf("fatal error get request: %w", err))
			}
			spew.Dump(projectInstance)
		}

	},
}

var myClient = &http.Client{Timeout: 10 * time.Second}

func getJsonWithToken(url string, token string, target interface{}) error {
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", token)
	r, err := myClient.Do(req)
	//fmt.Println(r)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
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
	getCmd.Flags().BoolVarP(&listFlag, "list", "l", false, "set to retrieve list of instances")
	getCmd.Flags().BoolVarP(&rawFlag, "raw", "r", false, "set to get all fields in output")
}
