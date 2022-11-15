package utils

/*
 * Implementation of some additional functions.
 */

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/lensesio/tableprinter"
	"github.com/spf13/cobra"
	"io"
	"os"
	"sigs.k8s.io/yaml"
)

// CloseBody closes response body
func CloseBody(body io.ReadCloser) {
	err := body.Close()
	cobra.CheckErr(err)
}

// PrintYamlOrJson prints output in json or yaml format
func PrintYamlOrJson(body []byte, outputType string) {
	if outputType == "json" {
		var out bytes.Buffer
		err := json.Indent(&out, body, "", "  ")
		cobra.CheckErr(err)
		fmt.Println(out.String())
	} else {
		out, err := yaml.JSONToYAML(body)
		cobra.CheckErr(err)
		fmt.Print(string(out))
	}
}

func RemoveColumnsFromTable(originalTable interface{}, fieldNames []string) interface{} {
	if len(fieldNames) == 0 {
		return originalTable
	}

	newTable := tableprinter.RemoveStructHeader(originalTable, fieldNames[0])
	for i := 1; i < len(fieldNames); i++ {
		newTable = tableprinter.RemoveStructHeader(newTable, fieldNames[i])
	}

	return newTable
}

// PrintTable prints a table
func PrintTable(in interface{}) {
	// init table printer
	printer := tableprinter.New(os.Stdout)

	// customize the table
	printer.BorderLeft = true
	printer.BorderRight = true
	printer.RowLine = true
	printer.CenterSeparator = "│"
	printer.ColumnSeparator = "│"
	printer.HeaderAlignment = tableprinter.AlignCenter
	printer.DefaultAlignment = tableprinter.AlignCenter
	printer.RowCharLimit = 30
	printer.RowTextWrap = true

	printer.Print(in)
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
