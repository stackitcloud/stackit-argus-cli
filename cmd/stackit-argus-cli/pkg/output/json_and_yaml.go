package output

import (
	"bytes"
	"encoding/json"
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
