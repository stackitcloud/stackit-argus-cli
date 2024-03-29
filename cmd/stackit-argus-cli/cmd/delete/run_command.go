package delete

/*
 * Runs delete commands.
 */

import (
	"fmt"
	"net/http"
	"time"

	"github.com/stackitcloud/stackit-argus-cli/internal/config"
	"github.com/stackitcloud/stackit-argus-cli/internal/utils"
)

// deleteRequest implements delete request and returns error
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

	if err := utils.ResponseMessage(res.StatusCode, resource, req.Method, res.Body); err != nil {
		return err
	}

	if err := res.Body.Close(); err != nil {
		return err
	}

	return nil
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
