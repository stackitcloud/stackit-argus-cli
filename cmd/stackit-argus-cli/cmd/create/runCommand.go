package create

/*
 * Runs create commands.
 */

import (
	"bytes"
	"fmt"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/YAMLToJSON"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/utils"
	"net/http"
	"os"
	"path"
	"time"
)

// postRequest implements post request and returns a status code
func postRequest(url, keyTarget string, targets []string, body []byte) (int, error) {
	authHeader := config.GetAuthHeader()
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	bodyReader := bytes.NewReader(body)
	req, err := http.NewRequest(http.MethodPost, url, bodyReader)
	if err != nil {
		return 0, err
	}

	// set auth header
	req.Header.Set("Authorization", authHeader)

	// set query parameters
	if len(targets) > 0 {
		values := req.URL.Query()
		for _, target := range targets {
			values.Add(keyTarget, target)
		}
		req.URL.RawQuery = values.Encode()
	}

	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	if err := res.Body.Close(); err != nil {
		return 0, err
	}

	if config.IsDebugMode() {
		println("response status: ", res.Status)
	}

	return res.StatusCode, nil
}

// runCommand call the url
func runCommand(url, resource, keyTarget string, targets []string) error {
	var body []byte
	var err error

	// print debug messages if debug mode is turned on
	if config.IsDebugMode() {
		fmt.Printf("create %s command called\n", resource)
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
		case "scrape configs":
			body, err = YAMLToJSON.ScrapeConfig(body)
		case "alert config receiver":
			body, err = YAMLToJSON.Receivers(body)
		case "alert config route":
			body, err = YAMLToJSON.Routes(body)
		default:
			body, err = YAMLToJSON.Convert(body)
		}

		if err != nil {
			return err
		}
	}

	// create the alert group
	status, err := postRequest(url, keyTarget, targets, body)
	if err != nil {
		return err
	}

	// print response status
	if err := utils.ResponseMessage(status, resource, "create"); err != nil {
		return err
	}

	return nil
}
