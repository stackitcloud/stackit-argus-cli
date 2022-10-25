package get

/*
 * Implementation of get request.
 */

import (
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/utils"
	"io"
	"net/http"
	"time"
)

// getRequest implements get request and returns a status code with response body
func getRequest(url string) (int, string) {
	authHeader := config.GetAuthHeader()
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	cobra.CheckErr(err)

	req.Header.Set("Authorization", authHeader)

	res, err := client.Do(req)
	defer utils.CloseBody(res.Body)

	body, err := io.ReadAll(res.Body)
	cobra.CheckErr(err)

	return res.StatusCode, string(body)
}
