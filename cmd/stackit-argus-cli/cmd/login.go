/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	logging "github.com/stackitcloud/stackit-argus-cli/internal/log"
)

var list = false
var projects *viper.Viper

const loggedIn = "LOGGED_IN"

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	// Args:  cobra.MinimumNArgs(1),
	Use:     "login",
	Short:   "Log in to your project",
	Long:    "Log in to your project by providing the name of the project as saved in the config file",
	Example: "stackit-argus-cli login <project-name>",
	Run: func(cmd *cobra.Command, args []string) {
		if viper.IsSet("projects") == true {
			projects = viper.Sub("projects")
		} else {
			logger.Info("projects configuration not found")

			if viper.IsSet(ProjectId) == true {
				logger.Info("falling back to default project",
					logging.String("projectId", viper.GetString(ProjectId)))

				viper.Set(loggedIn, viper.GetString(ProjectId))

				if err := viper.WriteConfigAs(viper.ConfigFileUsed()); err != nil {
					logger.Fatal("fatal saving project", logging.String("err", err.Error()))
				}

				return
			} else {
				logger.Info("Please either set a default project with the key \"PROJECT_ID\".\nOr register a project using the register command.")

				return
			}
		}

		if list == true {
			logger.Info("Projects list.", logging.Any("projects", projects.AllKeys()))
		} else {
			if len(args) == 0 {
				if viper.IsSet(loggedIn) == false {
					logger.Info("You are currently not logged into any project." +
						"Please specify a project name to log in")
				} else {
					logger.Info("You have been successfully logged into the project.",
						logging.String("project", viper.GetString(loggedIn)))
				}
			} else {
				login(args[0], projects)
			}
		}
	},
}

func login(projectName string, projects *viper.Viper) {
	if projects.IsSet(projectName) == false {
		logger.Info("the project with the name: \"" + projectName + "\" doesn't exist")
	} else {
		viper.Set(loggedIn, projects.GetString(projectName))

		if err := viper.WriteConfigAs(viper.ConfigFileUsed()); err != nil {
			logger.Fatal("fatal saving project", logging.String("err", err.Error()))
		}

		logger.Info("Logged in to "+projectName, logging.String("projectId", projects.GetString(projectName)))
	}
}

func init() {
	rootCmd.AddCommand(loginCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	loginCmd.Flags().BoolVarP(&list, "list", "l", false, "display list of projects to log in to")
}
