package git

import (
	addrepo "github.com/Ajlow2000/toolbox/app/git/addRepo"
	"github.com/Ajlow2000/toolbox/lib"
	"github.com/spf13/cobra"
)

var (
    url = "";
    path = "";
    addRepoReadme = "./app/git/addRepo/README.md"
)

var addRepoCmd = &cobra.Command{
	Use: lib.GetToolName(addRepoReadme),
	Short: lib.GetShortDesc(addRepoReadme),
    Long: lib.GetLongDesc(addRepoReadme),
	Run: func(cmd *cobra.Command, args []string) {
        if url == "" {
            cmd.Help()
        } else {
		    addrepo.Main(url, path)
        }
	},
}

func init() {
    addRepoCmd.Flags().StringVar(&url, "url", "", "The url pointing at a git repository")
    addRepoCmd.Flags().StringVar(&path, "path", "$HOME/repos", "The path to clone the specified url into")
}
