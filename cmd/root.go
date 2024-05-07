/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/Ajlow2000/toolbox/cmd/misc"
	"github.com/spf13/cobra"
)

var (
    debug = false
)


var rootCmd = &cobra.Command{
	Use:   "toolbox",
	Short: "Toolbox is a collection of utlities that make my life easier.",
    Long: "Toolbox is a collection of utilities that make my life easier and designed to be explorable and aliased for easier access to high use tools.",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
    // Commands
    rootCmd.AddCommand(misc.MiscCmd)

    // Flags
    rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "enable debugging")
}


