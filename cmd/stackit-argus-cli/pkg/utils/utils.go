package utils

/*
 * Implementation of some additional functions.
 */

import (
	"fmt"
	"github.com/spf13/cobra"
	"io"
)

// CloseBody closes response body
func CloseBody(body io.ReadCloser) {
	err := body.Close()
	cobra.CheckErr(err)
}

func ResponseMessage(status int, resource, action string) {
	m := map[string][]string{
		"delete": {"deleted", "deleting"},
		"get":    {"got", "getting"},
		"create": {"created", "creating"},
		"update": {"updated", "updating"},
	}

	switch status {
	case 202:
		fmt.Printf("%s was %s successfully\n", resource, m[action][0])
	case 400, 500:
		fmt.Printf("Something went wrong %s the %s\n", m[action][1], resource)
	case 403, 401:
		fmt.Printf("You are not authorized\n")
	case 404:
		fmt.Printf("%s not found\n", resource)
	default:
		fmt.Println("something went wrong. status code -", status)
	}
}
