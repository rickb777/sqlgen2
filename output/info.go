package output

import (
	"os"
	"fmt"
)

var Verbose = false

func Info(format string, args ...interface{}) {
	if Verbose {
		fmt.Fprintf(os.Stderr, format, args...)
	}
}

func Require(predicate bool, message string, args ...interface{}) {
	if !predicate {
		fmt.Fprintf(os.Stderr, message, args...)
		os.Exit(1)
	}
}
