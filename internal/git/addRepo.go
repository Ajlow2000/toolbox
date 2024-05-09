package git

import (
	"fmt"
	"os"
	"strings"

	"github.com/go-git/go-git/v5"
)

func AddRepo() {
	gitUrl := "https://github.com/Ajlow2000/home-manager"
	repoHome := "foo"
	params := strings.Split(gitUrl, "/")
	
	userName := params[3]
	projectName := params[4]
	localName := userName + "_" + projectName
	fmt.Println(localName)
	cloneOptions := git.CloneOptions{
		URL: gitUrl,
		Progress: os.Stdout,
	}
	git.PlainClone(repoHome, false, &cloneOptions)
	// git -C repoHome clone gitUrl localName
	// git -C "$HOME/repos" clone $GIT_URL $PROJECT_DIR
}

// os.Getenv(string Home) for get root