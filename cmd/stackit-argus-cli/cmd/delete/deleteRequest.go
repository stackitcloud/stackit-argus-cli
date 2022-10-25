package delete

/*
 * Implementation if delete request.
 */

import (
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd/config"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/pkg/utils"
	"net/http"
	"time"
)

// deleteRequest implements delete request and return a status code
func deleteRequest(url string) int {
	authHeader := config.GetAuthHeader()
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest(http.MethodDelete, url, nil)
	cobra.CheckErr(err)

	req.Header.Set("Authorization", authHeader)

	res, err := client.Do(req)
	defer utils.CloseBody(res.Body)

	return res.StatusCode
}
