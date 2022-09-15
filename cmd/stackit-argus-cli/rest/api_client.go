package rest

import (
	"encoding/json"
	"io"
	"net/http"
)

type serviceAccount struct {
	Id string `json:"id"`
}

type idToken struct {
	Token      string `json:"token"`
	ValidUntil string `json:"validUntil"`
}

func Authorize(projectId string) string {
	client := &http.Client{}

	// Create a service account via API. Use your own bearer token for creating it
	req, err := http.NewRequest("POST",
		"https://api-dev.stackit.cloud/service-account/v2/projects/"+projectId+"/service-accounts",
		nil)
	if err != nil {
		print("error is - ", err.Error())
		return ""
	}
	req.Header.Set("Authorization", "") //TODO: read Token for Authorization Header from Config file

	res, err := client.Do(req)
	if err != nil {
		print("error is - ", err.Error())
		return ""
	}

	if res.StatusCode != http.StatusCreated {
		res.Body.Close()
		println("Status is: ", res.Status)
		return ""
	}
	serviceAccount := &serviceAccount{}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		print("error is - ", err.Error())
		return ""
	}
	res.Body.Close()
	if err = json.Unmarshal(body, &serviceAccount); err != nil {
		println("error is - ", err.Error())
	}

	// Add the service account to the project with a role
	req, err = http.NewRequest("POST",
		"https://api-dev.stackit.cloud/membership/v1/projects/"+projectId+"/roles/project.member/service-accounts",
		nil)
	if err != nil {
		print("error is - ", err.Error())
		return ""
	}
	req.Header.Set("Authorization", "")
	req.Form.Set("serviceAccountId", serviceAccount.Id)
	res, err = client.Do(req)
	if err != nil {
		print("error is - ", err.Error())
		return ""
	}
	res.Body.Close()

	// Create an id token
	req, err = http.NewRequest("POST",
		"https://api-dev.stackit.cloud/service-account/v1/projects/"+projectId+"/service-accounts/"+serviceAccount.Id+"/access-tokens",
		nil)
	if err != nil {
		print("error is - ", err.Error())
		return ""
	}
	req.Header.Set("Authorization", "")
	req.Form.Set("ttlDays", "180")
	res, err = client.Do(req)
	if err != nil {
		print("error is - ", err.Error())
		return ""
	}
	if res.StatusCode != http.StatusCreated {
		res.Body.Close()
		println("Status is: ", res.Status)
		return ""
	}

	idToken := &idToken{}
	body, err = io.ReadAll(res.Body)
	if err != nil {
		print("error is - ", err.Error())
		return ""
	}
	res.Body.Close()
	if err = json.Unmarshal(body, &idToken); err != nil {
		println("error is - ", err.Error())
	}

	println("Here is your auth token:\n", idToken.Token)
	println("It will expire on: ", idToken.ValidUntil)
	return idToken.Token
}
