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

// getRequest implements get request and returns a status code with response body
func getRequest(url string) (int, []byte, error) {
	authHeader := config.GetAuthHeader()
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return 0, nil, err
	}

	req.Header.Set("Authorization", authHeader)

	res, err := client.Do(req)
	if err != nil {
		return 0, nil, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return 0, nil, err
	}

	if err := res.Body.Close(); err != nil {
		return 0, nil, err
	}

	if config.IsDebugMode() {
		fmt.Println("response status: ", res.Status)
	}

	return res.StatusCode, body, nil
}

// err := runCommand call the url
func runCommand(url, resource string, outputType config.OutputType) ([]byte, error) {
	// print debug messages if debug mode is turned on
	if config.IsDebugMode() {
		fmt.Printf("get %s command called\n", resource)
		fmt.Printf("url to call - %s\n", url)
	}

	// get response
	status, body, err := getRequest(url)
	if err != nil {
		return nil, err
	}

	// print response status
	if err := utils.ResponseMessage(status, resource, "get"); err != nil {
		return nil, err
	}

	// print response body
	if status == 200 {
		if outputType == "json" || outputType == "yaml" {
			if err := output.PrintYamlOrJson(body, string(outputType)); err != nil {
				return nil, err
			}
		}

		return body, nil
	}

	return nil, nil
}