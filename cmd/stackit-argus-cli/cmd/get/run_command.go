package get

/*
 * Runs get commands.
 */

import (
	"fmt"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/output"
	"github.com/stackitcloud/stackit-argus-cli/internal/config"
	"github.com/stackitcloud/stackit-argus-cli/internal/utils"
	"io"
	"net/http"
	"time"
)

var status = 0
var responseBody []byte

// GetRequest implements get request and returns a status code with response body
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

	return utils.ResponseMessageNew(res.StatusCode, resource, req.Method, res.Body)
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

	// print response status
	if err := utils.ResponseMessage(status, resource, "get"); err != nil {
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
