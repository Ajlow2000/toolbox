/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package misc

import (
    "github.com/Ajlow2000/toolbox/internal/misc/printPath"
	"github.com/spf13/cobra"
)

// printPathCmd represents the printPath command
var printPathCmd = &cobra.Command{
	Use:   "print-path",
	Short: "A brief description of your command",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
        printPath.main()
	},
}

func init() {

}
