package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func ResponseMessage(responseStatusCode int, resource string, action string, body io.ReadCloser) error {
	fmt.Printf("%s: %s \n", action, resource)
	if responseStatusCode != http.StatusUnauthorized {
		fmt.Printf("status code: %v\n", responseStatusCode)
	}

	if responseStatusCode >= http.StatusBadRequest {
		bodyByte, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		prettyBody, err := prettyJSONString(string(bodyByte))
		if err != nil {
			return err
		}
		if err := body.Close(); err != nil {
			return err
		}
		return fmt.Errorf("\n\033[31m%s\033[0m", prettyBody)
	}
	return nil
}

func prettyJSONString(str string) (string, error) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(str), "", "    "); err != nil {
		return "", err
	}
	return prettyJSON.String(), nil
}
