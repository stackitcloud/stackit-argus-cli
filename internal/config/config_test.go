package config

import (
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfigFileInit(t *testing.T) {
	testFile := "../../cmd/stackit-argus-cli/integration_test/.test.yaml"

	// set config file to change configs
	viper.SetConfigFile(testFile)

	// test if file doesn't exist
	err := InitFromConfigFile("")
	assert.Error(t, err, "file doesn't exist")

	// test if some config is missing
	err = InitFromConfigFile(testFile)
	assert.Error(t, err, "instance_id config is missing")

	// set the config, so config file should be valid
	viper.Set("base_url", "url")

	// test working case
	err = InitFromConfigFile(testFile)
	assert.NoError(t, err, "test working case")
}
