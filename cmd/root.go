package cmd

import (
	"os"

	"github.com/spf13/cobra"
)



var rootCmd = &cobra.Command{
	Use:   "weaver",
	Short: "weaves up a basic web application for ya",
	Long: `weaver is a tool to improve your dev experience.
It removes the constant rewrite of boiler plate code used in every project
	it includes support for many frameworks including:

	-Fiber
	-Standard Http
	-echo
	-chi
	-gin`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


