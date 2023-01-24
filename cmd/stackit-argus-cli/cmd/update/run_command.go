package update

/*
 * Runs update commands.
 */

import (
	"bytes"
	"fmt"
	"github.com/stackitcloud/stackit-argus-cli/internal/config"
	"github.com/stackitcloud/stackit-argus-cli/internal/services/yaml_to_json"
	"github.com/stackitcloud/stackit-argus-cli/internal/utils"
	"net/http"
	"os"
	"path"
	"time"
)

// convertToJson converts yaml request body to json
func convertToJson(body []byte, file, resource string) ([]byte, error) {
	if path.Ext(file) == ".yaml" {
		switch resource {
		case "alert config":
			return yaml_to_json.AlertConfig(body)
		case "scrape configs":
			return yaml_to_json.ScrapeConfigs(body)
		case "scrape config":
			return yaml_to_json.ScrapeConfig(body)
		case "alert config receiver":
			return yaml_to_json.Receivers(body)
		case "alert config route":
			return yaml_to_json.Routes(body)
		default:
			return yaml_to_json.Convert(body)
		}
	}

	return body, nil
}

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

	body, err = convertToJson(body, file, resource)
	if err != nil {
		return err
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
