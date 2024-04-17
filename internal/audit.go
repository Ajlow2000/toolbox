package internal

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

type gitInfo struct {
    url string
    isOutOfSync bool
    dirtyFiles []string
}


func getStatus(dir string) {
    if (path.IsDir()) {
    } else {
        // log error, files cant be git repos
    }
}

func findGitRepos(targetDir string, repos *[]gitInfo) {
    files, err := os.ReadDir(targetDir)
    if err != nil {
        log.Fatal(err)
    }

    for _, file := range files {
        if (file.IsDir() && file.Name() == ".git" ) {
            findGitRepos(targetDir, repos)
        }
    }
}

func AuditDir(targetDir string) {
    localRepos := make([]gitInfo, 20)
    findGitRepos(targetDir, &localRepos)
    return
}
