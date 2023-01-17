package update

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/config"
)

type UpdateSubCmd struct {
	Use      string
	Short    string
	Long     string
	Args     cobra.PositionalArgs
	MaxArgs  int
	Resource string
	BodyFile bool
	Path     []string
}

func GenerateUpdateSubCmd(u UpdateSubCmd) *cobra.Command {
	return &cobra.Command{
		Use:   u.Use,
		Short: u.Short,
		Long:  u.Long,
		Args:  u.Args,
		RunE: func(cmd *cobra.Command, args []string) error {
			if u.BodyFile && config.GetBodyFile() == "" {
				return errors.New("required flag \"--file(-f)\" not set")
			}

			method := "PUT"
			url := config.GetBaseUrl() + u.Path[0]
			if u.MaxArgs == 1 {
				if len(args) == 1 {
					url += fmt.Sprintf("/%s", args[0])
				} else {
					method = "PATCH"
					u.Resource += "s"
				}
			} else if u.MaxArgs == 2 {
				url += fmt.Sprintf("/%s/%s", args[0], u.Path[1])
				if len(args) == 2 {
					url += fmt.Sprintf("/%s", args[1])
				} else {
					method = "PATCH"
					u.Resource += "s"
				}
			}

			if err := runCommand(url, u.Resource, method); err != nil {
				cmd.SilenceUsage = true
				return err
			}

			return nil
		},
	}
}
