package git

import (
	"github.com/Ajlow2000/toolbox/lib"
	"github.com/spf13/cobra"
)

var gitReadme = "./app/git/README.md"

var GitCmd = &cobra.Command{
	Use: lib.GetToolName(gitReadme), 
	Short: lib.GetShortDesc(gitReadme),
    Long: lib.GetLongDesc(gitReadme),
	Run: func(cmd *cobra.Command, args []string) {
        cmd.Help()
	},
}

func init() {
    GitCmd.AddCommand(auditDirCmd)
	GitCmd.AddCommand(addRepoCmd)
}
