package foo

import "github.com/Ajlow2000/toolbox/lib"

func Baz() {
    // lib.DebugDocFiles()
    print("from foo.Baz(): " + lib.GetToolName("app/misc/README.md"))
}
