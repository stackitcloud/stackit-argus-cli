package config

/*
 * Init flags and configurations.
 */

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

// global variables to save configurations
var (
	token             string
	confFile          string
	projectId         string
	instanceId        string
	grafanaUsername   string
	grafanaPassword   string
	grafanaUrl        string
	bodyFile          string
	baseUrl           string
	debugMode         bool
	remoteWriteLimits bool
	targets           []string
	timeStamp         string
	endTimestamp      string
	tags              []string
	panelId           int8
)

// instance is a struct to unmarshal grafana credentials and url
type instance struct {
	Instance struct {
		GrafanaUrl           string `json:"grafanaUrl"`
		GrafanaAdminUser     string `json:"grafanaAdminUser"`
		GrafanaAdminPassword string `json:"grafanaAdminPassword"`
	} `json:"instance"`
}

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
	}
	token = viper.GetString("token")
	baseUrl = viper.GetString("base_url")
	grafanaUsername = viper.GetString("grafana_username")
	grafanaPassword = viper.GetString("grafana_password")

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
	cmd.PersistentFlags().BoolVarP(&remoteWriteLimits, "remote-write-limits", "r", false,
		"delete remote write config for credentials")
	cmd.PersistentFlags().VarP(&flagOutputType, "output", "o",
		"defines output format: yaml, json, table, wide or wide-table")
	cmd.PersistentFlags().StringVar(&timeStamp, "time", "", "set an annotation time stamp")
	cmd.PersistentFlags().StringVar(&endTimestamp, "time-end", "", "set an annotation time end stamp")
	cmd.PersistentFlags().StringSliceVar(&tags, "tags", nil, "set tags for an annotation")
	cmd.PersistentFlags().Int8Var(&panelId, "panel-id", -1, "set a panel id")

	// init flag possible values
	err = cmd.RegisterFlagCompletionFunc("output", outputFlagCompletion)
	cobra.CheckErr(err)

	// parse the flags passed to the command-line application
	if err := cmd.ParseFlags(os.Args[1:]); err != nil {
		cobra.CheckErr(err)
	}

	// cobra duplicates values, idk why. So need to clean it to avoid duplication
	targets = []string{}
	tags = []string{}
}

func setGrafanaInfo() error {
	authHeader := GetAuthHeader()
	req, err := http.NewRequest(http.MethodGet, strings.TrimSuffix(GetBaseUrl(), "/"), nil)
	if err != nil {
		println("!!!!!")
		return err
	}
	req.Header.Set("Authorization", authHeader)

	client := &http.Client{
		Timeout: time.Second * 10,
	}
	res, err := client.Do(req)
	if err != nil {
		println("?????")
		return err
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	if res.StatusCode != 200 {
		return errors.New("cannot get instance info, response: " + res.Status)
	}
	var i instance
	if err := json.Unmarshal(body, &i); err != nil {
		return err
	}
	if err := res.Body.Close(); err != nil {
		return err
	}

	grafanaUrl = i.Instance.GrafanaUrl
	if grafanaUsername == "" {
		grafanaUsername = i.Instance.GrafanaAdminUser
	}
	if grafanaPassword == "" {
		grafanaPassword = i.Instance.GrafanaAdminPassword
	}

	return nil
}

func GetGrafanaUrl() string {
	if grafanaUrl == "" {
		if err := setGrafanaInfo(); err != nil {
			cobra.CheckErr(err.Error())
		}
	}
	return grafanaUrl
}

func GetGrafanaUsername() string {
	if grafanaUsername == "" {
		if err := setGrafanaInfo(); err != nil {
			cobra.CheckErr(err.Error())
		}
	}
	return grafanaUsername
}

func GetGrafanaPassword() string {
	if grafanaPassword == "" {
		if err := setGrafanaInfo(); err != nil {
			cobra.CheckErr(err.Error())
		}
	}
	return grafanaPassword
}

// GetBaseUrl returns basic url for most of api calls
func GetBaseUrl() string {
	if baseUrl == "" {
		cobra.CheckErr("please, set a base url in a config file. " +
			"To set bae url run \"configure\" command")
	}
	if instanceId == "" {
		cobra.CheckErr("please, set a current instance id in a config file or with appropriate flag. " +
			"To set instance id run \"configure\" command.")
	}
	return fmt.Sprintf("%s/projects/%s/instances/%s/", baseUrl, GetProjectId(), instanceId)
}

func GetInstancesUrl() string {
	if baseUrl == "" {
		cobra.CheckErr("please, set a base url in a config file. " +
			"To set bae url run \"configure\" command")
	}
	return fmt.Sprintf("%s/projects/%s/instances", baseUrl, GetProjectId())
}

func GetProjectUrl() string {
	if baseUrl == "" {
		cobra.CheckErr("please, set a base url in a config file. " +
			"To set bae url run \"configure\" command")
	}
	return fmt.Sprintf("%s/projects/%s", baseUrl, GetProjectId())
}

func GetAuthHeader() string {
	if token == "" {
		cobra.CheckErr("please, set an auth token in a config file. " +
			"To set auth token run \"configure\" command")
	}
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
	if projectId == "" {
		cobra.CheckErr("please, set a current project id in a config file or with appropriate flag. " +
			"To set project id run \"configure\" command")
	}
	return projectId
}

func GetConfigFile() string {
	return confFile
}

// GetTargets returns backup targets slice
func GetTargets() []string {
	return targets
}

func convertTimestampToUnix(timestamp string) (int64, error) {
	if timestamp == "now" || timestamp == "" {
		return time.Now().UnixMilli(), nil
	} else {
		t, err := time.Parse("2006-01-02T15:04:05Z07:00", timestamp)
		if err != nil {
			return 0, err
		}
		return t.UnixMilli(), nil
	}
}

func GetTimestamp() (int64, error) {
	return convertTimestampToUnix(timeStamp)
}

// GetEndTimestamp returns end timestamp for an annotation
func GetEndTimestamp() (int64, error) {
	if endTimestamp == "" {
		return -1, nil
	}
	return convertTimestampToUnix(endTimestamp)
}

// GetTags returns tags for an annotation
func GetTags() []string {
	return tags
}

// GetPanelId returns panel id for an annotation
func GetPanelId() int8 {
	return panelId
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
	confFile, projectId, instanceId, bodyFile, token, baseUrl, targets, debugMode, remoteWriteLimits, timeStamp, endTimestamp, tags, panelId =
		"", "", "", "", "", "", []string{}, false, false, "", "", []string{}, 0
}
