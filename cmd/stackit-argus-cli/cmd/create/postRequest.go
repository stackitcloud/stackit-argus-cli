package create

/*
 * Implementation of post request.
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

// postRequest implements post request and returns a status code
func postRequest(url string, targets []string) int {
	authHeader := config.GetAuthHeader()
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	bodyFile := config.GetBodyFile()
	body, err := os.ReadFile(bodyFile)
	cobra.CheckErr(err)
	bodyReader := bytes.NewReader(body)

	req, err := http.NewRequest(http.MethodPost, url, bodyReader)
	cobra.CheckErr(err)

	// set auth header
	req.Header.Set("Authorization", authHeader)

	// set query parameters
	if len(targets) > 0 {
		for _, target := range targets {
			req.URL.Query().Add("backupTarget", target)
		}
	}

	res, err := client.Do(req)
	cobra.CheckErr(err)
	defer utils.CloseBody(res.Body)

	return res.StatusCode
}
