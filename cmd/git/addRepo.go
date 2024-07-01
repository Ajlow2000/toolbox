package git

import (
	"github.com/Ajlow2000/toolbox/internal/git"
	"github.com/spf13/cobra"
)

var (
    url = "";
    path = "";
)

var addRepoCmd = &cobra.Command{
	Use:   "add-repo",
	Short: "Utility for cloning repo's with standard name",
    Long: `
    Utility for cloning repo's with standard name

    By default, this utility expects a url to be specified with 
    '--url' and defaults to a clone location of $HOME/repos.
    This location can be overridden with '--path'. The repo will
    cloned into a directory of the custom name <username>_<repo>.
    `,
	Run: func(cmd *cobra.Command, args []string) {
        if url == "" {
            cmd.Help()
        } else {
		    git.AddRepo(url, path)
        }
	},
}

func init() {
    addRepoCmd.Flags().StringVar(&url, "url", "", "The url pointing at a git repository")
    addRepoCmd.Flags().StringVar(&path, "path", "$HOME/repos", "The path to clone the specified url into")
}
