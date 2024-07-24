package misc

import (
	printenviron "github.com/Ajlow2000/toolbox/app/misc/printEnviron"
	"github.com/spf13/cobra"
)

var printEnvironReadme = "app/misc/printEnviron/README.md"

var printEnvironCmd = &cobra.Command{
	Use: "", 
	Short: "",
    Long: "",
	Run: func(cmd *cobra.Command, args []string) {
        printenviron.Main()
	},
}

func init() {
}
