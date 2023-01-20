package yaml_to_json

/*
 * YAML to JSON converter.
 */

import (
	"encoding/json"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/models"
	yaml3 "gopkg.in/yaml.v3"
	"sigs.k8s.io/yaml"
)

// Convert converts yaml files to json
func Convert(content []byte) ([]byte, error) {
	return yaml.YAMLToJSON(content)
}

// AlertConfig converts alert config yaml file to json
func AlertConfig(content []byte) ([]byte, error) {
	var res models.Config

	if err := yaml3.Unmarshal(content, &res); err != nil {
		return nil, err
	}

	return json.Marshal(res)
}

// Receivers converts receivers yaml file to json
func Receivers(content []byte) ([]byte, error) {
	var res models.Receiver

	if err := yaml3.Unmarshal(content, &res); err != nil {
		return nil, err
	}

	return json.Marshal(res)
}

// Routes converts routes yaml file to json
func Routes(content []byte) ([]byte, error) {
	var res models.Route

	if err := yaml3.Unmarshal(content, &res); err != nil {
		return nil, err
	}

	return json.Marshal(res)
}

// ScrapeConfig converts scrape config yaml file to json
func ScrapeConfig(content []byte) ([]byte, error) {
	var res models.ScrapeConfig

	if err := yaml3.Unmarshal(content, &res); err != nil {
		return nil, err
	}

	return json.Marshal(res)
}

// ScrapeConfigs converts scrape configs yaml file to json
func ScrapeConfigs(content []byte) ([]byte, error) {
	var res []models.ScrapeConfig

	if err := yaml3.Unmarshal(content, &res); err != nil {
		return nil, err
	}

	return json.Marshal(res)
}
