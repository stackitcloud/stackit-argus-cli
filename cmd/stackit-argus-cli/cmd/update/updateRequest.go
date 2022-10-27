package update

/*
 * Implementation of put and patch requests.
 */

import (
	"bytes"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/utils"
	"net/http"
	"os"
	"time"
)

// updateRequest implements put and patch request and returns a status code
func updateRequest(url string, method string) int {
	authHeader := config.GetAuthHeader()
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	bodyFile := config.GetBodyFile()
	body, err := os.ReadFile(bodyFile)
	cobra.CheckErr(err)
	bodyReader := bytes.NewReader(body)

	req, err := http.NewRequest(method, url, bodyReader)
	cobra.CheckErr(err)

	// set auth header
	req.Header.Set("Authorization", authHeader)

	res, err := client.Do(req)
	defer utils.CloseBody(res.Body)

	return res.StatusCode
}
