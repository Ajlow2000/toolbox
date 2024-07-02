package printenviron

import (
	"fmt"
	"os"
)

// prints the current environment variables
// on newlines with entries in the format of 'key=val'
func Main()  {
    for _, env := range os.Environ() {
        fmt.Println(env)
    }
}
