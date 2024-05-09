package git

import (
	"strings"

	"github.com/Ajlow2000/toolbox/internal/git"
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
        git.AuditDir(target, logFile, ignoreList)
	},
}

func init() {
    auditDirCmd.Flags().StringVarP(&target, "target-dir", "t", "", "Target directory to search for git repos within")
    auditDirCmd.Flags().StringVarP(&logFile, "log-file", "f", "", "Path to log file. Passing 'stderr' prints logs to stderr")
    auditDirCmd.Flags().StringVarP(&ignore, "ignore-dirs", "i", "", "Provide a ':' deliminated list of paths to ignore")
}
