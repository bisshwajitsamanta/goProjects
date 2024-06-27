// Package cmd - Root.go file contains root command that will run if no sub command exist/*
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var cfgFile string

var RootCmd = &cobra.Command{
	Use:   "connect",
	Short: "Utility Tool",
	Long:  `Give Devops Guys capability to switch between sessions`,
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	/*
			Here you will define your flags and configuration settings Cobra supports persistent flags,
		    which, if defined here will be global for your application.
	*/

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (Default is $HOME/.switch.yaml")

	/*
	 Cobra also supports local flags, which will only run when this action is called directly.
	*/
}
