package output

import (
	"fmt"
	"github.com/rickb777/sqlgen2/sqlgen/parse/exit"
	"os"
)

var Verbose = false

func Info(format string, args ...interface{}) {
	if Verbose {
		fmt.Fprintf(os.Stderr, format, args...)
	}
}

func Require(predicate bool, message string, args ...interface{}) {
	if !predicate {
		exit.Fail(1, message, args...)
	}
}
