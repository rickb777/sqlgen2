package support

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"github.com/rickb777/sqlgen2"
	"github.com/rickb777/sqlgen2/require"
	"github.com/rickb777/sqlgen2/where"
)

// ReplaceTableName replaces all occurrences of "{TABLE}" with the table's name.
func ReplaceTableName(tbl sqlgen2.Table, query string) string {
	return strings.Replace(query, "{TABLE}", tbl.Name().String(), -1)
}

// QueryOneNullThing queries for one cell of one record. Normally, the holder will be sql.NullString or similar.
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

// Exec executes a modification query (insert, update, delete, etc) and returns the number of items affected.
func Exec(tbl sqlgen2.Table, req require.Requirement, query string, args ...interface{}) (int64, error) {
	LogQuery(tbl.Logger(), query, args...)
	res, err := tbl.Database().DB().ExecContext(tbl.Ctx(), query, args...)
	if err != nil {
		return 0, LogError(tbl.Logger(), err)
	}
	n, err := res.RowsAffected()
	return n, LogIfError(tbl.Logger(), require.ChainErrorIfExecNotSatisfiedBy(err, req, n))
}

// UpdateFields writes certain fields of all the records matching a 'where' expression.
func UpdateFields(tbl sqlgen2.Table, req require.Requirement, wh where.Expression, fields ...sql.NamedArg) (int64, error) {
	list := sqlgen2.NamedArgList(fields)
	assignments := strings.Join(list.Assignments(tbl.Dialect(), 1), ", ")
	whs, wargs := where.BuildExpression(wh, tbl.Dialect())
	query := fmt.Sprintf("UPDATE %s SET %s %s", tbl.Name(), assignments, whs)
	args := append(list.Values(), wargs...)
	return Exec(tbl, req, query, args...)
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
