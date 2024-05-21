package misc

import (
	printenviron "github.com/Ajlow2000/toolbox/internal/misc/printEnviron"
	"github.com/spf13/cobra"
)

var printEnvironCmd = &cobra.Command{
	Use:   "print-environ",
	Short: "Prints current environment variables on newlines",
	Long: `Prints current environment variables on newlines`,
	Run: func(cmd *cobra.Command, args []string) {
        printenviron.Main()
	},
}

func init() {
}
