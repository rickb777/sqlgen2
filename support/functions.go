package support

import (
	"database/sql"
	"fmt"
	"strings"
	"github.com/rickb777/sqlgen2"
	"github.com/rickb777/sqlgen2/require"
	"github.com/rickb777/sqlgen2/where"
	"context"
)

// ReplaceTableName replaces all occurrences of "{TABLE}" with the table's name.
func ReplaceTableName(tbl sqlgen2.Table, query string) string {
	return strings.Replace(query, "{TABLE}", tbl.Name().String(), -1)
}

// QueryOneNullThing queries for one cell of one record. Normally, the holder will be sql.NullString or similar.
func QueryOneNullThing(tbl sqlgen2.Table, req require.Requirement, holder interface{}, query string, args ...interface{}) error {
	var n int64 = 0
	query = ReplaceTableName(tbl, query)
	database := tbl.Database()

	rows, err := tbl.Query(query, args...)
	if err != nil {
		return err
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(holder)

		if err == sql.ErrNoRows {
			return database.LogIfError(require.ErrorIfQueryNotSatisfiedBy(req, 0))
		} else {
			n++
		}

		if rows.Next() {
			n++ // not singular
		}
	}

	return database.LogIfError(require.ChainErrorIfQueryNotSatisfiedBy(rows.Err(), req, n))
}

//-------------------------------------------------------------------------------------------------

func GetInt64List(tbl sqlgen2.Table, req require.Requirement, sqlname string, wh where.Expression, qc where.QueryConstraint) ([]int64, error) {
	dialect := tbl.Dialect()
	database := tbl.Database()

	whs, args := where.BuildExpression(wh, dialect)
	orderBy := where.BuildQueryConstraint(qc, dialect)
	query := fmt.Sprintf("SELECT %s FROM %s %s %s", dialect.Quote(sqlname), tbl.Name(), whs, orderBy)

	rows, err := tbl.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var v int64
	list := make([]int64, 0, 10)

	for rows.Next() {
		err = rows.Scan(&v)
		if err == sql.ErrNoRows {
			return list, database.LogIfError(require.ErrorIfQueryNotSatisfiedBy(req, int64(len(list))))
		} else {
			list = append(list, v)
		}
	}
	return list, database.LogIfError(require.ChainErrorIfQueryNotSatisfiedBy(rows.Err(), req, int64(len(list))))
}

//-------------------------------------------------------------------------------------------------

// Query is the low-level request method for this table.
//
// The query is logged using whatever logger is configured. If an error arises, this too is logged.
//
// The args are for any placeholder parameters in the query.
//
// The caller must call rows.Close() on the result.
func Query(ctx context.Context, tbl sqlgen2.Table, query string, args ...interface{}) (*sql.Rows, error) {
	database := tbl.Database()
	database.LogQuery(query, args...)
	rows, err := tbl.Execer().QueryContext(ctx, query, args...)
	return rows, database.LogIfError(err)
}


// Exec executes a modification query (insert, update, delete, etc) and returns the number of items affected.
//
// The query is logged using whatever logger is configured. If an error arises, this too is logged.
func Exec(ctx context.Context, tbl sqlgen2.Table, req require.Requirement, query string, args ...interface{}) (int64, error) {
	database := tbl.Database()
	database.LogQuery(query, args...)
	res, err := tbl.Execer().ExecContext(ctx, query, args...)
	if err != nil {
		return 0, database.LogError(err)
	}
	n, err := res.RowsAffected()
	return n, database.LogIfError(require.ChainErrorIfExecNotSatisfiedBy(err, req, n))
}


// UpdateFields writes certain fields of all the records matching a 'where' expression.
func UpdateFields(ctx context.Context, tbl sqlgen2.Table, req require.Requirement, wh where.Expression, fields ...sql.NamedArg) (int64, error) {
	list := sqlgen2.NamedArgList(fields)
	assignments := strings.Join(list.Assignments(tbl.Dialect(), 1), ", ")
	whs, wargs := where.BuildExpression(wh, tbl.Dialect())
	query := fmt.Sprintf("UPDATE %s SET %s %s", tbl.Name(), assignments, whs)
	args := append(list.Values(), wargs...)
	return Exec(ctx, tbl, req, query, args...)
}
