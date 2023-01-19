package config

/*
 * Init flags and configurations.
 */

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

// global variables to save configurations
var (
	token             string
	confFile          string
	projectId         string
	instanceId        string
	bodyFile          string
	baseUrl           string
	debugMode         bool
	remoteWriteLimits bool
	targets           []string
)

// InitFromConfigFile inits info from configuration file
func InitFromConfigFile(f string) error {
	viper.SetConfigFile(f)

	// ask user to set up conf file if it is not found
	if err := viper.ReadInConfig(); err != nil {
		return nil
	}

	// get instance id from a config file, if it is not set via flag
	if instanceId == "" {
		instanceId = viper.GetString("instance_id")
	}

	// get project id from a config file, if it is not set via flag
	if projectId == "" {
		projectId = viper.GetString("project_id")
		if projectId == "" {
			return errors.New("please, set a current project id in a config file or with appropriate flag. " +
				"To set project id run \"configure\" command")
		}
	}

	// get auth token from config file
	token = viper.GetString("token")
	if token == "" {
		return errors.New("please, set an auth token in a config file. " +
			"To set auth token run \"configure\" command")
	}

	// get base url from config file
	baseUrl = viper.GetString("base_url")
	if baseUrl == "" {
		return errors.New("please, set a base url in a config file. " +
			"To set bae url run \"configure\" command")
	}

	return nil
}

// Init inits global flags
func Init(cmd *cobra.Command) {
	homeDir, err := os.UserHomeDir()
	cobra.CheckErr(err)

	// init flags
	cmd.PersistentFlags().StringVarP(&confFile, "config", "c", homeDir+"/.stackit-argus-cli.yaml",
		"provide config file")
	cmd.PersistentFlags().StringVarP(&instanceId, "instance-id", "i", "",
		"provide instance id that should be used")
	cmd.PersistentFlags().StringVarP(&projectId, "project-id", "p", "",
		"provide project id that should be used")
	cmd.PersistentFlags().BoolVarP(&debugMode, "debug", "d", false,
		"turn on debug mode to see more information about the command")
	cmd.PersistentFlags().StringVarP(&bodyFile, "file", "f", "", "provide file with request body")
	cmd.PersistentFlags().StringSliceVarP(&targets, "target", "t", nil,
		"defines targets of the backup schedules, possible targets: alertConfig, alertRules, scrapeConfig or grafana")
	cmd.PersistentFlags().BoolVarP(&remoteWriteLimits, "remote-write-limits", "r", false, "delete remote write config for credentials")
	cmd.PersistentFlags().VarP(&flagOutputType, "output", "o", "defines output format: yaml, json or wide")

	// init flag possible values
	err = cmd.RegisterFlagCompletionFunc("output", outputFlagCompletion)
	cobra.CheckErr(err)

	// parse the flags passed to the command-line application
	if err := cmd.ParseFlags(os.Args[1:]); err != nil {
		cobra.CheckErr(err)
	}

	// cobra duplicates values, idk why. So need to clean it to avoid duplication
	targets = []string{}
}

// GetBaseUrl returns basic url for mostly all api calls
func GetBaseUrl() string {
	if instanceId == "" {
		cobra.CheckErr("please, set a current instance id in a config file or with appropriate flag. " +
			"To set instance id run \"configure\" command.")
	}
	return fmt.Sprintf("%s/projects/%s/instances/%s/", baseUrl, projectId, instanceId)
}

// GetInstancesUrl returns instances url
func GetInstancesUrl() string {
	return fmt.Sprintf("%s/projects/%s/instances", baseUrl, projectId)
}

// GetProjectUrl returns project url
func GetProjectUrl() string {
	return fmt.Sprintf("%s/projects/%s", baseUrl, projectId)
}

// GetAuthHeader returns auth header to make api calls
func GetAuthHeader() string {
	return fmt.Sprintf("Bearer %s", token)
}

// GetInstanceId returns instance id from flag if set, otherwise gets instance id from the config file
func GetInstanceId() string {
	return instanceId
}

// GetBodyFile returns path to file that contains body of a request
func GetBodyFile() string {
	return bodyFile
}

// GetProjectId returns project id from flag if set, otherwise gets project id from the config file
func GetProjectId() string {
	return projectId
}

// GetConfigFile returns configuration file name
func GetConfigFile() string {
	return confFile
}

// GetTargets returns backup targets slice
func GetTargets() []string {
	return targets
}

// IsDebugMode checks if debug mode is turned on
func IsDebugMode() bool {
	return debugMode
}

// IsRemoteWriteLimits checks if remote write limits is set
func IsRemoteWriteLimits() bool {
	return remoteWriteLimits
}

// ResetConfigurations reset all configurations to default value
func ResetConfigurations() {
	token, confFile, projectId, instanceId, bodyFile, token, baseUrl, targets, debugMode, remoteWriteLimits =
		"", "", "", "", "", "", "", []string{}, false, false
}
