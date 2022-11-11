package delete

/*
 * Runs delete commands.
 */

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/utils"
	"net/http"
	"time"
)

// deleteRequest implements delete request and returns a status code
func deleteRequest(url string) int {
	authHeader := config.GetAuthHeader()
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest(http.MethodDelete, url, nil)
	cobra.CheckErr(err)

	req.Header.Set("Authorization", authHeader)

	res, err := client.Do(req)
	cobra.CheckErr(err)
	defer utils.CloseBody(res.Body)

	return res.StatusCode
}

func runCommand(url, resource string) {
	// print debug messages if debug mode is turned on
	if config.IsDebugMode() {
		fmt.Printf("delete %s command called", resource)
		fmt.Printf("url to call - %s\n", url)
	}

	// create the alert group
	status := deleteRequest(url)

	// print response status
	utils.ResponseMessage(status, resource, "delete")
}
