package misc

import (
	printenviron "github.com/Ajlow2000/toolbox/app/misc/printEnviron"
	"github.com/Ajlow2000/toolbox/lib"
	"github.com/spf13/cobra"
)

var printEnvironReadme = "./app/misc/printEnviron/README.md"

var printEnvironCmd = &cobra.Command{
	Use: lib.GetToolName(printEnvironReadme), 
	Short: lib.GetShortDesc(printEnvironReadme),
    Long: lib.GetLongDesc(printEnvironReadme),
	Run: func(cmd *cobra.Command, args []string) {
        printenviron.Main()
	},
}

func init() {
}
