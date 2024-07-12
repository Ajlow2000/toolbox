package printenviron

import (
	"fmt"
)

// #cgo CFLAGS: -I.
// #cgo LDFLAGS: -L. -lzgo
// #include <./main.h>
import "C"

// prints the current environment variables
// on newlines with entries in the format of 'key=val'
func Main()  {
    fmt.Printf("Invoking zig library!\n")
    fmt.Println("Done ", C.x(10))
}
