package git

import (
	"strings"

	auditdir "github.com/Ajlow2000/toolbox/internal/git/auditDir"
	"github.com/spf13/cobra"
)

var (
    target = ""
    logFile = ""
    ignore = ""
)

var auditDirCmd = &cobra.Command{
	Use:   "audit-dir",
	Short: "Generate a report of the status of all git repositories within the specified [dir]",
	Long: "Logging is disabled until a log file is provided via --log-file",
	Run: func(cmd *cobra.Command, args []string) {
        ignoreList := strings.Split(ignore, ":")
        auditdir.Main(target, ignoreList)
	},
}

func init() {
    auditDirCmd.Flags().StringVarP(&target, "target-dir", "t", "$HOME", "Target directory to search for git repos within. Defaults to $HOME")
    auditDirCmd.Flags().StringVarP(&ignore, "ignore-dirs", "i", "", "Provide a ':' deliminated list of paths to ignore")
}
