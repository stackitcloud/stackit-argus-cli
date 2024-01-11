package delete

/*
 * Runs delete commands.
 */

import (
	"fmt"
	"github.com/stackitcloud/stackit-argus-cli/internal/config"
	"github.com/stackitcloud/stackit-argus-cli/internal/utils"
	"net/http"
	"time"
)

// deleteRequest implements delete request and returns a status code
func deleteRequest(url string, resource string) error {
	authHeader := config.GetAuthHeader()
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", authHeader)

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	
	return utils.ResponseMessageNew(res.StatusCode, resource, req.Method, res.Body)
}

// runCommand call the url
func runCommand(url, resource string) error {
	// print debug messages if debug mode is turned on
	if config.IsDebugMode() {
		fmt.Printf("delete %s command called\n", resource)
		fmt.Printf("url to call - %s\n", url)
	}

	// create the alert group
	return deleteRequest(url, resource)
}
