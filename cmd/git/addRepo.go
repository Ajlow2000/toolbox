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
	Long: "",
	Run: func(cmd *cobra.Command, args []string) {
		git.AddRepo(url, path)
	},
}

func init() {
    addRepoCmd.Flags().StringVar(&url, "url", "", "The url pointing at a git repository")
    addRepoCmd.Flags().StringVar(&path, "path", "", "The path to clone the specified url into")
}
