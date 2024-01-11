package get

/*
 * Runs get commands.
 */

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/output"
	"github.com/stackitcloud/stackit-argus-cli/internal/config"
	"github.com/stackitcloud/stackit-argus-cli/internal/utils"
)

var status = 0
var responseBody []byte

// GetRequest implements get request and returns error
func GetRequest(url string, resource string) error {
	authHeader := config.GetAuthHeader()
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", authHeader)

	res, err := client.Do(req)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	status = res.StatusCode
	responseBody = body

	if err := utils.ResponseMessage(res.StatusCode, resource, req.Method, res.Body); err != nil {
		return err
	}

	if err := res.Body.Close(); err != nil {
		return err
	}

	return nil
}

// err := runCommand call the url
func runCommand(url, resource string, outputType config.OutputType) ([]byte, error) {
	// print debug messages if debug mode is turned on
	if config.IsDebugMode() {
		fmt.Printf("get %s command called\n", resource)
		fmt.Printf("url to call - %s\n", url)
	}

	// get response
	if err := GetRequest(url, resource); err != nil {
		return nil, err
	}

	// print response body
	if status == 200 {
		if outputType == "json" || outputType == "yaml" {
			if err := output.PrintYamlOrJson(responseBody, string(outputType)); err != nil {
				return nil, err
			}
		}

		return responseBody, nil
	}

	return nil, nil
}
