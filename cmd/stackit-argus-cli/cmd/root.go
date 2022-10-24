package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	logging "github.com/stackitcloud/stackit-argus-cli/internal/log"
	"os"
)

// variables to save flags' values
var (
	confFile   string
	instanceId string
	projectId  string
	token      string
	debugMode  bool
)

var logger = logging.New()

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "stackit-argus-cli",
	Short:   "Manage ARGUS resources",
	Long:    `Manage ARGUS resources, like instances with CRUD operations`,
	Version: "1.0.0",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		return err
	}

	return nil
}

// init commands and flags for the root command
func init() {
	rootCmd.AddCommand(getCmd)
	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(updateCmd)
	rootCmd.AddCommand(deleteCmd)

	rootCmd.PersistentFlags().StringVarP(&confFile, "config", "c", "./.stackit-argus-cli.yaml",
		"provide config file (default is ./.stackit-argus-cli.yaml)")
	rootCmd.PersistentFlags().StringVarP(&instanceId, "instanceId", "i", "",
		"provide instance id that should be used")
	rootCmd.PersistentFlags().StringVarP(&projectId, "projectId", "p", "",
		"provide project id that should be used")
	rootCmd.PersistentFlags().BoolVarP(&debugMode, "debug", "d", false,
		"turn on debug mode to see more information about the command")

	cobra.OnInitialize(initConfig)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.SetConfigFile(confFile)

	// read config file
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

// NewArgusCliCmd returns root cmd
func NewArgusCliCmd() *cobra.Command {
	return rootCmd
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
