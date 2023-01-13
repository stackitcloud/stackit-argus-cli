package update

/*
 * Runs update commands.
 */

import (
	"bytes"
	"fmt"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/utils"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/yamltojson"
	"net/http"
	"os"
	"path"
	"time"
)

// updateRequest implements put and patch request and returns a status code
func updateRequest(url string, method string, body []byte) (int, error) {
	authHeader := config.GetAuthHeader()
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	bodyReader := bytes.NewReader(body)
	req, err := http.NewRequest(method, url, bodyReader)
	if err != nil {
		return 0, err
	}

	// set header
	req.Header.Set("Authorization", authHeader)
	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	if err := res.Body.Close(); err != nil {
		return 0, err
	}

	if config.IsDebugMode() {
		fmt.Println("response status: ", res.Status)
	}

	return res.StatusCode, nil
}

// runCommand call the url
func runCommand(url, resource, method string) error {
	var body []byte
	var err error

	// print debug messages if debug mode is turned on
	if config.IsDebugMode() {
		fmt.Printf("update %s command called\n", resource)
		fmt.Printf("url to call - %s\n", url)
	}

	// get file content
	file := config.GetBodyFile()
	if file != "" {
		body, err = os.ReadFile(file)
		if err != nil {
			return err
		}
	} else {
		body = nil
	}

	// convert body to json if yaml file has been given
	if path.Ext(file) == ".yaml" {
		switch resource {
		case "alert config":
			body, err = yamltojson.AlertConfig(body)
		case "scrape configs":
			body, err = yamltojson.ScrapeConfigs(body)
		case "scrape config":
			body, err = yamltojson.ScrapeConfig(body)
		case "alert config receiver":
			body, err = yamltojson.Receivers(body)
		case "alert config route":
			body, err = yamltojson.Routes(body)
		default:
			body, err = yamltojson.Convert(body)
		}

		if err != nil {
			return err
		}
	}

	// create the alert group
	status, err := updateRequest(url, method, body)
	if err != nil {
		return err
	}

	// print response status
	if err := utils.ResponseMessage(status, resource, "update"); err != nil {
		return err
	}

	return err
}
