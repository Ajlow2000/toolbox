package git

import (
	"github.com/Ajlow2000/toolbox/internal/git"
	"github.com/spf13/cobra"
)

// addRepoCmd represents the addRepo command
var addRepoCmd = &cobra.Command{
	Use:   "add-repo",
	Short: "Utility for cloning repo's with standard name",
	Long: "",
	Run: func(cmd *cobra.Command, args []string) {
		git.AddRepo()
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addRepoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addRepoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
