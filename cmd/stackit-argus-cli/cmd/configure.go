package cmd

/*
 * Create and write configurations.
 */

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/term"
	"os"
)

var baseUrl = "https://argus.api.eu01.stackit.cloud/v1"

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
		conf, err = readLongString()
		if err != nil {
			return err
		}
		viper.Set("token", conf)
		viper.Set("base_url", baseUrl)

		if err := viper.WriteConfig(); err != nil {
			return err
		}

		return nil
	},
}

// readLongString provides a way to read more than 1024 characters from the terminal by switching the terminal into
// raw mode. Otherwise, long strings like the authentication token would be truncated to 1024 characters because of
// canonical input mode for terminals.
func readLongString() (string, error) {
	termStateBackup, _ := term.MakeRaw(int(os.Stdin.Fd()))
	defer term.Restore(int(os.Stdin.Fd()), termStateBackup) // nolint:errcheck

	result := ""
	characters := make([]byte, 1024)
	for {
		n, err := os.Stdin.Read(characters)
		if err != nil {
			return result, err
		}
		for i := 0; i < n; i++ {
			if characters[i] == 0x03 || // Ctrl+C
				characters[i] == 0x0d { // Enter
				fmt.Print("\r\n")
				return result, nil
			}
			fmt.Printf("%s", string(characters[i]))
			result = result + string(characters[i])
		}
	}
}
