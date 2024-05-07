/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package misc

import (
	"os"

	"github.com/spf13/cobra"
)

var printEnvironCmd = &cobra.Command{
	Use:   "print-environ",
	Short: "Prints current environment variables on newlines",
	Long: `Prints current environment variables on newlines`,
	Run: func(cmd *cobra.Command, args []string) {
        for _, env := range os.Environ() {
            println(env)
        }

	},
}

func init() {
}
