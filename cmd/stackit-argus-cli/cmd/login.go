/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var list = false
var projects *viper.Viper

const loggedIn = "LOGGED_IN"

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	// Args:  cobra.MinimumNArgs(1),
	Use:   "login",
	Short: "Log in to your project",
	Long: `Log in to your project by providing the name of the project as saved in the config file:

stackit-argus-cli login <project-name>`,
	Run: func(cmd *cobra.Command, args []string) {
		//viper.SetConfigName(".stackit-argus-cli")
		//viper.SetConfigType("yaml")
		//viper.AddConfigPath(".")
		err := viper.ReadInConfig()
		if err != nil {
			panic(fmt.Errorf("fatal error config file: %w", err))
		}
		if viper.IsSet("projects") == true {
			projects = viper.Sub("projects")
		} else {
			fmt.Println("projects configuration not found")
			if viper.IsSet(ProjectId) == true {
				fmt.Println("Falling back to default Project: \"PROJECT_ID\"")
				viper.Set(loggedIn, viper.GetString(ProjectId))
				err = viper.WriteConfigAs(viper.ConfigFileUsed())
				if err != nil {
					panic(fmt.Errorf("fatal saving project: %w", err))
				}
				return
			} else {
				fmt.Println("Please either set a default project with the key \"PROJECT_ID\".\nOr register a project using the register command.")
				return
			}
		}
		if list == true {
			fmt.Println(projects.AllKeys())
		} else {
			if len(args) == 0 {
				if viper.IsSet(loggedIn) == false {
					fmt.Println("You are currently not logged into any project")
					fmt.Println("Please specify a project name to log in to")
				} else {
					fmt.Println("You are logged into project: " + viper.GetString(loggedIn))
				}
			} else {
				login(args[0], projects)
			}
		}
	},
}

func login(projectName string, projects *viper.Viper) {

	if projects.IsSet(projectName) == false {
		fmt.Println("the project with the name: \"" + projectName + "\" doesn't exist")
	} else {
		viper.Set(loggedIn, projects.GetString(projectName))
		err := viper.WriteConfigAs(viper.ConfigFileUsed())
		if err != nil {
			panic(fmt.Errorf("fatal saving project: %w", err))
		}
		fmt.Println("Logged in to " + projectName + ": (ProjectId: " + projects.GetString(projectName))
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
