package git

import (
	"fmt"
	"os"
	// "os"
	"strings"
	// "github.com/go-git/go-git/v5"
)

// Build a name for the local checkout of the specified url.
// Namespaces repos under the owner's username (in lower case)
func buildLocalDirName(url string) string {
    userName := ""
    projectName := ""

    if strings.HasPrefix(url, "git@") {     // ssh
        url = strings.Split(url, ":")[1]
        params := strings.Split(url, "/")
        userName = params[0]
        projectName = params[1]
    } else {                                // https
        params := strings.Split(url, "/")
        userName = params[3]
        projectName = params[4]
    }

    userName = strings.ToLower(userName)
    projectName, _ = strings.CutSuffix(projectName, ".git")

    return userName + "_" + projectName;
}

func AddRepo(url string, path string) {
	localName := buildLocalDirName(url)
    fmt.Fprintln(os.Stderr, localName)

	// cloneOptions := git.CloneOptions{
	// 	URL: url,
	// 	Progress: os.Stdout,
	// }

    // repo, err := git.PlainClone(path, false, &cloneOptions)
    // if err != nil {
    //     fmt.Fprintln(os.Stderr, "An error occured when cloning '" + url + "' into '" + path + "'")
    // }
	// git -C repoHome clone gitUrl localName
	// git -C "$HOME/repos" clone $GIT_URL $PROJECT_DIR
}

// os.Getenv(string Home) for get root
