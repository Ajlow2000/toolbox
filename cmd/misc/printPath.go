package misc

import (
	printpath "github.com/Ajlow2000/toolbox/app/misc/printPath"
	"github.com/Ajlow2000/toolbox/lib"
	"github.com/spf13/cobra"
)

var printPathReadme = "./app/misc/printPath/README.md"

var printPathCmd = &cobra.Command{
	Use: lib.GetToolName(printPathReadme), 
	Short: lib.GetShortDesc(printPathReadme),
    Long: lib.GetLongDesc(printPathReadme),
	Run: func(cmd *cobra.Command, args []string) {
        printpath.Main()
	},
}

func init() {

}
