package misc

import (
	"os"
	"strings"
)

// PrintPath prints entries in $PATH on newlines
func PrintPath()  {
    var path = strings.Split(os.Getenv("PATH"), ":")
    for _, entry := range path {
        println(entry)
    }
}
