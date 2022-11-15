package get

/*
 * Runs get commands.
 */

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/utils"
	"io"
	"net/http"
	"time"
)

// getRequest implements get request and returns a status code with response body
func getRequest(url string) (int, []byte) {
	authHeader := config.GetAuthHeader()
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	cobra.CheckErr(err)

	req.Header.Set("Authorization", authHeader)

	res, err := client.Do(req)
	cobra.CheckErr(err)
	defer utils.CloseBody(res.Body)

	body, err := io.ReadAll(res.Body)
	cobra.CheckErr(err)

	return res.StatusCode, body
}

// runCommand call the url
func runCommand(url, resource string, outputType config.OutputType) []byte {
	// print debug messages if debug mode is turned on
	if config.IsDebugMode() {
		fmt.Printf("get %s command called", resource)
		fmt.Printf("url to call - %s\n", url)
	}

	// get response
	status, body := getRequest(url)

	// print response status
	utils.ResponseMessage(status, resource, "get")

	// print response body
	if status == 200 {
		if outputType == "json" || outputType == "yaml" {
			utils.PrintYamlOrJson(body, string(outputType))
		}

		return body
	}

	return nil
}
