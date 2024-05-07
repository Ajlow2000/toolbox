package misc

import (
	"github.com/spf13/cobra"
)

var MiscCmd = &cobra.Command{
	Use:   "misc",
	Short: "Miscellaneous Tools",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
        cmd.Help()
	},
}

func init() {
    MiscCmd.AddCommand(printPathCmd)
    MiscCmd.AddCommand(printEnvironCmd)
}
