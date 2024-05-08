package git

import (
	"github.com/Ajlow2000/toolbox/internal/git"
	"github.com/spf13/cobra"
)

var (
    target = ""
    logFile = ""
)

var auditDirCmd = &cobra.Command{
	Use:   "audit-dir",
	Short: "Generate a report of the status of all git repositories within the specified [dir]",
	Long: "Logging is disabled until a log file is provided via --log-file",
	Run: func(cmd *cobra.Command, args []string) {
        git.AuditDir(target, logFile)
	},
}

func init() {
    auditDirCmd.Flags().StringVarP(&target, "target-dir", "t", "", "Target directory to search for git repos within")
    auditDirCmd.Flags().StringVarP(&logFile, "log-file", "f", "", "Path to log file. Passing 'stderr' prints logs to stderr")
}
