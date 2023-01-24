package utils

import (
	"errors"
	"fmt"
)

// ResponseMessage generates a response message depending on response status code
func ResponseMessage(status int, resource, action string) error {
	m := map[string][]string{
		"delete": {"deleted", "deleting"},
		"get":    {"got", "getting"},
		"create": {"created", "creating"},
		"update": {"updated", "updating"},
	}

	switch status {
	case 202, 200:
		return nil
	case 400, 500:
		return fmt.Errorf("\033[31msomething went wrong %s the %s\033[0m", m[action][1], resource)
	case 403, 401:
		return errors.New("\033[31mYou are not authorized\033[0m")
	case 404:
		return fmt.Errorf("\033[31m%s not found\033[0m", resource)
	case 502:
		return errors.New("\033[31mconnection to object storage could not be established.\033[0m")
	default:
		return fmt.Errorf("something went wrong. status code - %d\033[0m", status)
	}
}
