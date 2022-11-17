package utils

/*
 * Implementation of some additional functions.
 */

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"sigs.k8s.io/yaml"
)

// CloseBody closes response body
func CloseBody(body io.ReadCloser) {
	err := body.Close()
	if err != nil {
		panic(err.Error())
	}
}

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
func ResponseMessage(status int, resource, action string) {
	m := map[string][]string{
		"delete": {"deleted", "deleting"},
		"get":    {"got", "getting"},
		"create": {"created", "creating"},
		"update": {"updated", "updating"},
	}

	switch status {
	case 202, 200:
		fmt.Printf("\033[32msuccessfully %s %s\n\033[0m", m[action][0], resource)
	case 400, 500:
		fmt.Printf("\033[31msomething went wrong %s the %s\n\033[0m", m[action][1], resource)
	case 403, 401:
		fmt.Printf("\033[31mYou are not authorized\n\033[0m")
	case 404:
		fmt.Printf("\033[31m%s not found\n\033[0m", resource)
	case 502:
		fmt.Println("\033[31mconnection to object storage could not be established.", "\033[0m")
	default:
		fmt.Println("something went wrong. status code -", status, "\033[0m")
	}
}
