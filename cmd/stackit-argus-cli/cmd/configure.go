package cmd

/*
 * Create and write configurations.
 */

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

// configureCmd represents the configure command
var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Configure cli.",
	Long: "Set instance id, project id and authorization token. Configuration file with these values" +
		" will be create in your home directory with \".stackit-argus-cli.yaml\" name.",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		var conf string

		confFilePath, err := os.UserHomeDir()
		if err != nil {
			return err
		}

		// create configuration file in the HOME directory
		if _, err := os.OpenFile(confFilePath+"/.stackit-argus-cli.yaml", os.O_CREATE|os.O_RDWR, 0644); err != nil {
			return err
		}

		viper.SetConfigFile(confFilePath + "/.stackit-argus-cli.yaml")

		// ask user for configurations and put them into the conf file
		fmt.Print("Instance id: ")
		if _, err := fmt.Scanf("%s", &conf); err != nil {
			return err
		}
		viper.Set("instance_id", conf)
		fmt.Print("Project id: ")
		if _, err := fmt.Scanf("%s", &conf); err != nil {
			return err
		}
		viper.Set("project_id", conf)
		fmt.Print("Authorization token: ")
		if _, err := fmt.Scanf("%s", &conf); err != nil {
			return err
		}
		viper.Set("token", conf)
		viper.Set("base_url", "https://argus.api.eu01.stackit.cloud/v1")

		if err := viper.WriteConfig(); err != nil {
			return err
		}

		return nil
	},
}
