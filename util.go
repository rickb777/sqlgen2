package sqlgen2

import (
	"log"
	"strings"
)

// LogQuery writes query info to the logger, if it is not nil.
func LogQuery(logger *log.Logger, query string, args ...interface{}) {
	if logger != nil {
		query = strings.TrimSpace(query)
		if len(args) > 0 {
			logger.Printf(query + " %v\n", args)
		} else {
			logger.Println(query)
		}
	}
}
