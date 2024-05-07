package misc

import (
	"github.com/Ajlow2000/toolbox/internal/misc"
	"github.com/spf13/cobra"
)

var printPathCmd = &cobra.Command{
	Use:   "print-path",
	Short: "Prints entries in $PATH on newlines",
	Long: "",
	Run: func(cmd *cobra.Command, args []string) {
        misc.PrintPath()
	},
}

func init() {

}
