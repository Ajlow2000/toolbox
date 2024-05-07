/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package misc

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// printPathCmd represents the printPath command
var printPathCmd = &cobra.Command{
	Use:   "print-path",
	Short: "Prints entries in $PATH on newlines",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
        var path = strings.Split(os.Getenv("PATH"), ":")
        for _, entry := range path {
            println(entry)
        }
	},
}

func init() {

}
