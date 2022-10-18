package rest

import (
	"encoding/json"
	"github.com/spf13/viper"
	logging "github.com/stackitcloud/stackit-argus-cli/internal/log"
	"io"
	"net/http"
	"time"
)

type serviceAccount struct {
	Id string `json:"id"`
}

type idToken struct {
	Token      string `json:"token"`
	ValidUntil string `json:"validUntil"`
}

var logger = logging.New()

var client = &http.Client{
	Timeout: time.Second * 10,
}

func closeResponseBody(body io.ReadCloser) {
	err := body.Close()

	if err != nil {
		logger.Fatal("cannot close request body", logging.String("err", err.Error()))
	}
}

func getServiceAccount(projectId string) *serviceAccount {
	req, err := http.NewRequest("POST",
		"https://api-dev.stackit.cloud/service-account/v2/projects/"+projectId+"/service-accounts",
		nil)
	if err != nil {
		logger.Fatal("cannot create a request", logging.String("err", err.Error()))
	}

	req.Header.Add("Authorization", "Bearer "+viper.GetString("token")) // read Token for Authorization Header from Config file

	res, err := client.Do(req)
	if err != nil {
		logger.Fatal("cannot make a auth request", logging.String("err", err.Error()))
	}

	defer closeResponseBody(res.Body)

	if res.StatusCode != http.StatusCreated {
		logger.Info("cannot authorize", logging.String("status", res.Status))

		return nil
	}

	serviceAccount := &serviceAccount{}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		logger.Fatal("cannot read response body", logging.String("err", err.Error()))
	}

	if err = json.Unmarshal(body, serviceAccount); err != nil {
		logger.Error("cannot unmarshal service account", logging.String("err", err.Error()))
	}

	return serviceAccount
}

func addServiceAccountToProject(projectId string, serviceAccount *serviceAccount) {
	req, err := http.NewRequest("POST",
		"https://api-dev.stackit.cloud/membership/v1/projects/"+projectId+"/roles/project.member/service-accounts",
		nil)
	if err != nil {
		logger.Fatal("cannot create a request", logging.String("err", err.Error()))
	}

	req.Form.Set("serviceAccountId", serviceAccount.Id)

	res, err := client.Do(req)
	if err != nil {
		logger.Fatal("cannot make a request", logging.String("err", err.Error()))
	}

	closeResponseBody(res.Body)
}

func createIdToken(projectId string, serviceAccount *serviceAccount) *idToken {
	req, err := http.NewRequest("POST",
		"https://api-dev.stackit.cloud/service-account/v1/projects/"+projectId+"/service-accounts/"+serviceAccount.Id+"/access-tokens",
		nil)
	if err != nil {
		logger.Fatal("cannot create a request", logging.String("err", err.Error()))
	}

	req.Form.Set("ttlDays", "180")

	res, err := client.Do(req)
	if err != nil {
		logger.Fatal("cannot make a request", logging.String("err", err.Error()))
	}

	defer closeResponseBody(res.Body)

	if res.StatusCode != http.StatusCreated {
		logger.Info("cannot create id token", logging.String("status", res.Status))

		return nil
	}

	idToken := &idToken{}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		logger.Fatal("cannot read response body", logging.String("err", err.Error()))
	}

	if err = json.Unmarshal(body, idToken); err != nil {
		logger.Error("cannot unmarshal id token", logging.String("err", err.Error()))
	}

	return idToken
}

func Authorize(projectId string) string {
	serviceAccount := getServiceAccount(projectId)

	if serviceAccount == nil {
		return ""
	}

	addServiceAccountToProject(projectId, serviceAccount)

	idToken := createIdToken(projectId, serviceAccount)

	logger.Info("Auth token was successfully created.", logging.String("auth token", idToken.Token),
		logging.String("valid until", idToken.ValidUntil))

	return idToken.Token
}
