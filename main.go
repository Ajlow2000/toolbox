/*
Copyright © 2024 Alec Lowry @Ajlow2000
*/
package main

import (
	"embed"

	"github.com/Ajlow2000/toolbox/foo"
	"github.com/Ajlow2000/toolbox/lib"
)

//go:embed README.md
//go:embed VERSION
//go:embed app/git/README.md
//go:embed app/git/addRepo/README.md
//go:embed app/git/auditDir/README.md
//go:embed app/misc/README.md
//go:embed app/misc/printEnviron/README.md
//go:embed app/misc/printPath/README.md
var docsFS embed.FS

func main() {
    lib.SetDocsFS(docsFS)
    // lib.DebugDocFiles()

	// cmd.Blah()
    foo.Baz()
    // misc.Foo()
}
