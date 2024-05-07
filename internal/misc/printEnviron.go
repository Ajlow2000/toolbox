package misc

import "os"

// PrintEnviron prints the current environment variables
// on newlines with entries in the format of 'key=val'
func PrintEnviron()  {
    for _, env := range os.Environ() {
        println(env)
    }
}
