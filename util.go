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
			logger.Printf(query+" %v\n", args)
		} else {
			logger.Println(query)
		}
	}
}

// LogIfError writes error info to the logger, if the logger and the error are not nil.
func LogIfError(logger *log.Logger, err error) error {
	if logger != nil && err != nil {
		logger.Printf("Error: %s\n", err)
	}
	return err
}

// LogError writes error info to the logger, if the logger is not nil.
func LogError(logger *log.Logger, err error) error {
	if logger != nil {
		logger.Printf("Error: %s\n", err)
	}
	return err
}
