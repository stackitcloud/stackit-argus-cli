/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/viper"
	"os"

	"github.com/spf13/cobra"
)

var projectName string

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("login called")
		// TODO: Set ProjectID as env variable
		// viper.SetConfigName("config")
		// viper.SetConfigType("yaml")
		// viper.AddConfigPath(".")
		//if err := viper.ReadInConfig(); err != nil {
		//	if _, ok := err.(viper.ConfigFileNotFoundError); ok {
		//		// Config file not found; ignore error if desired
		//		fmt.Println("Config file not found:", err, viper.ConfigFileUsed())
		//		return
		//	} else {
		// Config file was found but another error was produced
		//		fmt.Println("Config file found but something else went wrong:", err, viper.ConfigFileUsed())
		//		return
		//	}
		//}
		fmt.Println(viper.ConfigFileUsed(), "Project:", projectName)
		err := os.Setenv(ProjectId, viper.GetString(projectName))
		if err != nil {
			cobra.CheckErr(err)
		} else {
			fmt.Println(viper.GetString(projectName))
			fmt.Println("Logged into Project: " + os.Getenv(ProjectId))
		}
	},
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

	loginCmd.Flags().StringVarP(&projectName, "projectName", "n", "", "specify the name of the project to log in to")
}
