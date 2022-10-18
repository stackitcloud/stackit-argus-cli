/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"errors"
	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/viper"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/models"
	logging "github.com/stackitcloud/stackit-argus-cli/internal/log"
	"io"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

var listFlag bool
var rawFlag bool

func closeResponseBody(body io.ReadCloser) {
	err := body.Close()

	if err != nil {
		logger.Fatal("cannot close request body", logging.String("err", err.Error()))
	}
}

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:     "get",
	Short:   "Retrieve information about an ARGUS resource",
	Long:    "Retrieve information about an ARGUS resource by passing its project id and optionally its instance id",
	Example: "  stackit-argus-cli get instance <instanceId>",
	Run: func(cmd *cobra.Command, args []string) {
		token := "Bearer " + viper.GetString("token")

		url := getBaseUrl()

		if listFlag == false {
			if len(args) > 0 {
				url = url + "/" + args[0]
			} else {
				logger.Info("Please provide an instance ID or set the list flag")

				return
			}
		}

		logger.Info("The request url was successfully created.", logging.String("url", url))

		instanceList := &models.InstanceList{}
		projectInstance := &models.ProjectInstance{}

		if listFlag == true {
			if err := getJsonWithToken(url, token, instanceList); err != nil {
				return
			}

			spew.Dump(instanceList)
		} else {
			if err := getJsonWithToken(url, token, projectInstance); err != nil {
				return
			}

			spew.Dump(projectInstance)
		}
	},
}

var myClient = &http.Client{
	Timeout: time.Second * 10,
}

func getJsonWithToken(url string, token string, target interface{}) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		logger.Fatal("cannot create a request", logging.String("err", err.Error()))
	}

	req.Header.Set("Authorization", token)

	r, err := myClient.Do(req)
	if err != nil {
		logger.Fatal("cannot make the request", logging.String("err", err.Error()))
	}

	if r.StatusCode != 200 {
		logger.Error("request failed", logging.String("status", r.Status))

		return errors.New("request failed")
	}

	defer closeResponseBody(r.Body)

	if err := json.NewDecoder(r.Body).Decode(target); err != nil {
		logger.Fatal("cannot decode a response body", logging.String("err", err.Error()))
	}

	return nil
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
