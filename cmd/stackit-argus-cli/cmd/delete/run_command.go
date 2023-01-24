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
func deleteRequest(url string) (int, error) {
	authHeader := config.GetAuthHeader()
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return 0, err
	}

	req.Header.Set("Authorization", authHeader)

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
func runCommand(url, resource string) error {
	// print debug messages if debug mode is turned on
	if config.IsDebugMode() {
		fmt.Printf("delete %s command called\n", resource)
		fmt.Printf("url to call - %s\n", url)
	}

	// create the alert group
	status, err := deleteRequest(url)
	if err != nil {
		return err
	}

	// print response status
	if err := utils.ResponseMessage(status, resource, "delete"); err != nil {
		return err
	}

	return nil
}
