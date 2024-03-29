package update

/*
 * Update alert groups.
 */

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/internal/config"
)

// AlertGroupsCmd represents the alertGroups command
var AlertGroupsCmd = &cobra.Command{
	Use:   "alert-groups <group-name>",
	Short: "Update alert groups.",
	Long:  "Patch alert groups if group name was not specified, otherwise update alert group config.",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if config.GetBodyFile() == "" {
			return errors.New("required flag \"--file(-f)\" not set")
		}

		// generate an url
		url := config.GetBaseUrl() + "alertgroups"

		// modify url and debug msg if an argument has been given
		method := "PATCH"
		resource := "alert groups"
		if len(args) == 1 {
			resource = "alert group"
			method = "PUT"
			url += fmt.Sprintf("/%s", args[0])
		}

		// call command
		if err := runCommand(url, resource, method); err != nil {
			cmd.SilenceUsage = true
			return err
		}

		return nil
	},
}
