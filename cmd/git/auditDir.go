package git

import (
	"strings"

	auditdir "github.com/Ajlow2000/toolbox/app/git/auditDir"
	"github.com/Ajlow2000/toolbox/lib"
	"github.com/spf13/cobra"
)

var (
    target = ""
    logFile = ""
    ignore = ""
    auditDirReadme = "./app/git/auditDir/README.md"
)

var auditDirCmd = &cobra.Command{
	Use: lib.GetToolName(auditDirReadme), 
	Short: lib.GetShortDesc(auditDirReadme),
    Long: lib.GetLongDesc(auditDirReadme),
	Run: func(cmd *cobra.Command, args []string) {
        ignoreList := strings.Split(ignore, ":")
        auditdir.Main(target, ignoreList)
	},
}

func init() {
    auditDirCmd.Flags().StringVarP(&target, "target-dir", "t", "$HOME", "Target directory to search for git repos within. Defaults to $HOME")
    auditDirCmd.Flags().StringVarP(&ignore, "ignore-dirs", "i", "", "Provide a ':' deliminated list of paths to ignore")
}
