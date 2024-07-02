package misc

import (
	"github.com/Ajlow2000/toolbox/lib"
	"github.com/spf13/cobra"
)

var miscReadme = "./app/misc/README.md"

var MiscCmd = &cobra.Command{
	Use: lib.GetToolName(miscReadme), 
	Short: lib.GetShortDesc(miscReadme),
    Long: lib.GetLongDesc(miscReadme),
	Run: func(cmd *cobra.Command, args []string) {
        cmd.Help()
	},
}

func init() {
    MiscCmd.AddCommand(printPathCmd)
    MiscCmd.AddCommand(printEnvironCmd)
}
