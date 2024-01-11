package create

/*
 * Runs create commands.
 */

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/stackitcloud/stackit-argus-cli/internal/config"
	"github.com/stackitcloud/stackit-argus-cli/internal/services/yaml_to_json"
	"github.com/stackitcloud/stackit-argus-cli/internal/utils"
)

// convertToJson converts yaml request body to json
func convertToJson(body []byte, file, resource string) ([]byte, error) {
	var err error

	if path.Ext(file) == ".yaml" {
		switch resource {
		case "scrape configs":
			body, err = yaml_to_json.ScrapeConfig(body)
		case "alert config receiver":
			body, err = yaml_to_json.Receivers(body)
		case "alert config route":
			body, err = yaml_to_json.Routes(body)
		default:
			body, err = yaml_to_json.Convert(body)
		}

		if err != nil {
			return nil, err
		}
	}

	return body, nil
}

// postRequest implements post request and returns a status code
func postRequest(url, keyTarget string, resource string, targets []string, body []byte) error {
	authHeader := config.GetAuthHeader()
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	bodyReader := bytes.NewReader(body)
	req, err := http.NewRequest(http.MethodPost, url, bodyReader)
	if err != nil {
		return err
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
		return err
	}

	return utils.ResponseMessage(res.StatusCode, resource, req.Method, res.Body)
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

	body, err = convertToJson(body, file, resource)
	if err != nil {
		return err
	}

	// create the alert group

	return postRequest(url, keyTarget, resource, targets, body)
}
