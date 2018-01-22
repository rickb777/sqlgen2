package support

import (
	"database/sql"
	"strings"
	"github.com/rickb777/sqlgen2"
	"github.com/rickb777/sqlgen2/require"
	"log"
)

// ReplaceTableName replaces all occurrences of "{TABLE}" with the table's name.
func ReplaceTableName(tbl sqlgen2.Table, query string) string {
	return strings.Replace(query, "{TABLE}", tbl.Name().String(), -1)
}

func QueryOneNullThing(tbl sqlgen2.Table, req require.Requirement, holder interface{}, query string, args ...interface{}) error {
	var n int64 = 0
	query = ReplaceTableName(tbl, query)
	LogQuery(tbl.Logger(), query, args...)

	rows, err := tbl.DB().QueryContext(tbl.Ctx(), query, args...)
	if err != nil {
		return LogError(tbl.Logger(), err)
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(holder)

		if err == sql.ErrNoRows {
			return LogIfError(tbl.Logger(), require.ErrorIfQueryNotSatisfiedBy(req, 0))
		} else {
			n++
		}

		if rows.Next() {
			n++ // not singular
		}
	}

	return LogIfError(tbl.Logger(), require.ChainErrorIfQueryNotSatisfiedBy(rows.Err(), req, n))
}

//-------------------------------------------------------------------------------------------------

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
