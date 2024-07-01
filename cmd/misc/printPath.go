package misc

import (
	printpath "github.com/Ajlow2000/toolbox/app/misc/printPath"
	"github.com/spf13/cobra"
)

var printPathCmd = &cobra.Command{
	Use:   "print-path",
	Short: "Prints entries in $PATH on newlines",
	Long: "",
	Run: func(cmd *cobra.Command, args []string) {
        printpath.Main()
	},
}

func init() {

}
