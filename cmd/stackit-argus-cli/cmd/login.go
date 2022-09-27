/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var list = false

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
		projects := viper.Sub("projects")
		if projects == nil {
			panic("projects configuration not found")
		}
		if list == true {
			fmt.Println(projects.AllKeys())
		} else {
			if len(args) == 0 {
				if os.Getenv(ProjectId) == "" {
					fmt.Println("You are currently not logged into any project")
					fmt.Println("Please specify a project name to log in to")
				} else {
					fmt.Println("You are logged into project: " + os.Getenv(ProjectId))
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
		err := os.Setenv(ProjectId, projects.GetString(projectName))
		if err != nil {
			panic(fmt.Errorf("fatal error logging in: %w", err))
		} else {
			fmt.Println("Logged in to " + projectName + ": (ProjectId: " + os.Getenv(ProjectId) + ")")
		}
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
