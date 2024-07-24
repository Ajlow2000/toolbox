package misc

import (
	"github.com/Ajlow2000/toolbox/lib"
	"github.com/spf13/cobra"
)

var miscReadme = "app/misc/README.md"
var name string

var MiscCmd = &cobra.Command{
	Use: name, 
	Short: "bar",
    Long:  "baz",
	Run: func(cmd *cobra.Command, args []string) {
        cmd.Help()
	},
}

func Initialize() {
    MiscCmd.AddCommand(printPathCmd)
    MiscCmd.AddCommand(printEnvironCmd)
}
