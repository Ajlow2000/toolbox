package git

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
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

func Main(url string, path string) {
    path = os.Expand(path, os.Getenv)

	localName := buildLocalDirName(url)

    cmd := exec.Command("git", "-C", path, "clone", url, localName)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    err := cmd.Run()
    if err != nil {
        fmt.Fprintln(os.Stderr, err.Error())
    }
}
