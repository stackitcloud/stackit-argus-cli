package config

/*
 * Init flags and configurations.
 */

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"time"
)

// global variables to save configurations
var (
	token      string
	confFile   string
	projectId  string
	instanceId string
	debugMode  bool
)

// initFromConfigFile inits info from configuration file
func initFromConfigFile() {
	viper.SetConfigFile(confFile)

	err := viper.ReadInConfig()
	cobra.CheckErr(err)

	// get instance id from a config file, if it is not set via flag
	if instanceId == "" {
		instanceId = viper.GetString("current_instance")
		if instanceId == "" {
			fmt.Println("Please, set a current instance id in a config file. Default config file path is" +
				"'./.stackit-argus-cli.yaml'. Current instance id key is 'current_instance'.")

			os.Exit(1)
		}
	}

	// get project id from a config file, if it is not set via flag
	if projectId == "" {
		projectId = viper.GetString("current_project")
		if projectId == "" {
			fmt.Println("Please, set a current project id in a config file. Default config file path is" +
				"'./.stackit-argus-cli.yaml'. Current project id key is 'current_project'.")

			os.Exit(1)
		}
	}

	// get auth token from config file
	token = viper.GetString("token")
	if token == "" {
		fmt.Println("Please, set an auth token in a config file. Default config file path is" +
			"'./.stackit-argus-cli.yaml'. Token id key is 'token'.")

		os.Exit(1)
	}
}

// InitConfig inits global flags and info from configuration file
func InitConfig(cmd *cobra.Command) {
	// init global flags
	cmd.PersistentFlags().StringVarP(&confFile, "config", "c", "./.stackit-argus-cli.yaml",
		"provide config file (default is ./.stackit-argus-cli.yaml)")
	cmd.PersistentFlags().StringVarP(&instanceId, "instanceId", "i", "",
		"provide instance id that should be used")
	cmd.PersistentFlags().StringVarP(&projectId, "projectId", "p", "",
		"provide project id that should be used")
	cmd.PersistentFlags().BoolVarP(&debugMode, "debug", "d", false,
		"turn on debug mode to see more information about the command")

	// init info from configuration file
	initFromConfigFile()
}

// InitInputFile inits file content of which will be used as a body for http request
func InitInputFile(cmd *cobra.Command) {
	// init flag
	cmd.PersistentFlags().StringP("file", "f", "", "provide file with request body")
	if err := cmd.MarkPersistentFlagRequired("file"); err != nil {
		return
	}
}

// GetHttpClient get http client
func GetHttpClient() *http.Client {
	return &http.Client{
		Timeout: time.Second * 10,
	}
}

// GetBaseUrl get basic url for mostly all api calls
func GetBaseUrl() string {
	return fmt.Sprintf("https://api-dev.stackit.cloud/argus-service/v1/projects/%s/instances/%s", projectId, instanceId)
}

// GetAuthHeader returns auth header to make api calls
func GetAuthHeader() string {
	return fmt.Sprintf("Bearer %s", token)
}

// GetInstanceId gets instance id from flag if set, otherwise gets instance id from the config file
func GetInstanceId() string {
	return instanceId
}

// GetProjectId gets project id from flag if set, otherwise gets project id from the config file
func GetProjectId() string {
	return projectId
}

// IsDebugMode checks if debug mode is set
func IsDebugMode() bool {
	return debugMode
}
