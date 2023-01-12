package utils

/*
 * Implementation of some additional functions.
 */

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"sigs.k8s.io/yaml"
)

// PrintYamlOrJson prints output in json or yaml format
func PrintYamlOrJson(body []byte, outputType string) error {
	if outputType == "json" {
		var out bytes.Buffer
		if err := json.Indent(&out, body, "", "  "); err != nil {
			return err
		}
		fmt.Println(out.String())
	} else {
		out, err := yaml.JSONToYAML(body)
		if err != nil {
			return err
		}
		fmt.Print(string(out))
	}

	return nil
}

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
		fmt.Printf("\033[32msuccessfully %s %s\n\033[0m", m[action][0], resource)
		return nil
	case 400, 500:
		return errors.New(fmt.Sprintf("\033[31msomething went wrong %s the %s\033[0m", m[action][1], resource))
	case 403, 401:
		return errors.New(fmt.Sprintf("\033[31mYou are not authorized\033[0m"))
	case 404:
		return errors.New(fmt.Sprintf("\033[31m%s not found\033[0m", resource))
	case 502:
		return errors.New(fmt.Sprintf("\033[31mconnection to object storage could not be established.\033[0m"))
	default:
		return errors.New(fmt.Sprintf("something went wrong. status code - %d\033[0m", status))
	}
}
