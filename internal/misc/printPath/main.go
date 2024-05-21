package printpath

import (
	"os"
	"strings"
)

// prints entries in $PATH on newlines
func Main()  {
    var path = strings.Split(os.Getenv("PATH"), ":")
    for _, entry := range path {
        println(entry)
    }
}
