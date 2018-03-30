package exit

import (
	"os"
	"fmt"
)

// seam for testing
var Fail = func(code int, message string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "Error: " + message, args...)
	os.Exit(code)
}

func TestableExit() {
	Fail = func(code int, message string, args ...interface{}) {
		panic(fmt.Sprintf(message, args...))
	}
}