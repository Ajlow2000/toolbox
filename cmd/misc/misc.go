/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package misc

import (
	"github.com/spf13/cobra"
)

// miscCmd represents the misc command
var MiscCmd = &cobra.Command{
	Use:   "misc",
	Short: "Misc is a palette that contains miscellaneous tools",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
        cmd.Help()
	},
}

func init() {
    MiscCmd.AddCommand(printPathCmd)
    MiscCmd.AddCommand(printEnvironCmd)
}
