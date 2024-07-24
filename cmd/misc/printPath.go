package misc

import (
	printpath "github.com/Ajlow2000/toolbox/app/misc/printPath"
	"github.com/spf13/cobra"
)

var printPathReadme = "app/misc/printPath/README.md"

var printPathCmd = &cobra.Command{
	Use: "", 
	Short: "",
    Long: "",
	Run: func(cmd *cobra.Command, args []string) {
        printpath.Main()
	},
}

