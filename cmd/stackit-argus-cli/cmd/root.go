package cmd

/*
 * Main command implementation (stackit-argus-cli).
 */

import (
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/internal/config"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:           "stackit-argus-cli",
	Short:         "Manage ARGUS resources",
	Args:          cobra.NoArgs,
	Version:       "1.0.0",
	SilenceErrors: true,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() error {
	defer config.ResetConfigurations()

	if err := config.InitFromConfigFile(config.GetConfigFile()); err != nil {
		return err
	}

	if err := rootCmd.Execute(); err != nil {
		return err
	}

	return nil
}

// init commands and flags for the root command
func init() {
	// init subcommands
	rootCmd.AddCommand(configureCmd)
	rootCmd.AddCommand(getCmd)
	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(updateCmd)
	rootCmd.AddCommand(deleteCmd)

	// init flags
	config.Init(rootCmd)
}

// NewArgusCliCmd returns root cmd
func NewArgusCliCmd() *cobra.Command {
	return rootCmd
}
