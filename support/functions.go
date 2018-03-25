package support

import (
	"database/sql"
	"fmt"
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

// SliceStringList requests a columnar slice of strings from a specified column.
func SliceStringList(tbl sqlgen2.Table, req require.Requirement, sqlname string, wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	query, args := sliceSql(tbl, sqlname, wh, qc)
	rows, err := tbl.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return tbl.Database().ScanStringList(req, rows)
}

// SliceIntList requests a columnar slice of ints from a specified column.
func SliceIntList(tbl sqlgen2.Table, req require.Requirement, sqlname string, wh where.Expression, qc where.QueryConstraint) ([]int, error) {
	query, args := sliceSql(tbl, sqlname, wh, qc)
	rows, err := tbl.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return tbl.Database().ScanIntList(req, rows)
}

// SliceUintList requests a columnar slice of uints from a specified column.
func SliceUintList(tbl sqlgen2.Table, req require.Requirement, sqlname string, wh where.Expression, qc where.QueryConstraint) ([]uint, error) {
	query, args := sliceSql(tbl, sqlname, wh, qc)
	rows, err := tbl.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return tbl.Database().ScanUintList(req, rows)
}

// SliceInt64List requests a columnar slice of int64s from a specified column.
func SliceInt64List(tbl sqlgen2.Table, req require.Requirement, sqlname string, wh where.Expression, qc where.QueryConstraint) ([]int64, error) {
	query, args := sliceSql(tbl, sqlname, wh, qc)
	rows, err := tbl.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return tbl.Database().ScanInt64List(req, rows)
}

// SliceUint64List requests a columnar slice of uint64s from a specified column.
func SliceUint64List(tbl sqlgen2.Table, req require.Requirement, sqlname string, wh where.Expression, qc where.QueryConstraint) ([]uint64, error) {
	query, args := sliceSql(tbl, sqlname, wh, qc)
	rows, err := tbl.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return tbl.Database().ScanUint64List(req, rows)
}

// SliceInt32List requests a columnar slice of int32s from a specified column.
func SliceInt32List(tbl sqlgen2.Table, req require.Requirement, sqlname string, wh where.Expression, qc where.QueryConstraint) ([]int32, error) {
	query, args := sliceSql(tbl, sqlname, wh, qc)
	rows, err := tbl.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return tbl.Database().ScanInt32List(req, rows)
}

// SliceUint32List requests a columnar slice of uint32s from a specified column.
func SliceUint32List(tbl sqlgen2.Table, req require.Requirement, sqlname string, wh where.Expression, qc where.QueryConstraint) ([]uint32, error) {
	query, args := sliceSql(tbl, sqlname, wh, qc)
	rows, err := tbl.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return tbl.Database().ScanUint32List(req, rows)
}

// SliceInt16List requests a columnar slice of int16s from a specified column.
func SliceInt16List(tbl sqlgen2.Table, req require.Requirement, sqlname string, wh where.Expression, qc where.QueryConstraint) ([]int16, error) {
	query, args := sliceSql(tbl, sqlname, wh, qc)
	rows, err := tbl.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return tbl.Database().ScanInt16List(req, rows)
}

// SliceUint16List requests a columnar slice of uint16s from a specified column.
func SliceUint16List(tbl sqlgen2.Table, req require.Requirement, sqlname string, wh where.Expression, qc where.QueryConstraint) ([]uint16, error) {
	query, args := sliceSql(tbl, sqlname, wh, qc)
	rows, err := tbl.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return tbl.Database().ScanUint16List(req, rows)
}

// SliceInt8List requests a columnar slice of int8s from a specified column.
func SliceInt8List(tbl sqlgen2.Table, req require.Requirement, sqlname string, wh where.Expression, qc where.QueryConstraint) ([]int8, error) {
	query, args := sliceSql(tbl, sqlname, wh, qc)
	rows, err := tbl.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return tbl.Database().ScanInt8List(req, rows)
}

// SliceUint8List requests a columnar slice of uint8 from a specified column.
func SliceUint8List(tbl sqlgen2.Table, req require.Requirement, sqlname string, wh where.Expression, qc where.QueryConstraint) ([]uint8, error) {
	query, args := sliceSql(tbl, sqlname, wh, qc)
	rows, err := tbl.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return tbl.Database().ScanUint8List(req, rows)
}

// SliceFloat32List requests a columnar slice of float32s from a specified column.
func SliceFloat32List(tbl sqlgen2.Table, req require.Requirement, sqlname string, wh where.Expression, qc where.QueryConstraint) ([]float32, error) {
	query, args := sliceSql(tbl, sqlname, wh, qc)
	rows, err := tbl.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return tbl.Database().ScanFloat32List(req, rows)
}

// SliceFloat64List requests a columnar slice of float64s from a specified column.
func SliceFloat64List(tbl sqlgen2.Table, req require.Requirement, sqlname string, wh where.Expression, qc where.QueryConstraint) ([]float64, error) {
	query, args := sliceSql(tbl, sqlname, wh, qc)
	rows, err := tbl.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return tbl.Database().ScanFloat64List(req, rows)
}


func sliceSql(tbl sqlgen2.Table, sqlname string, wh where.Expression, qc where.QueryConstraint) (string, []interface{}) {
	dialect := tbl.Dialect()
	whs, args := where.BuildExpression(wh, dialect)
	orderBy := where.BuildQueryConstraint(qc, dialect)
	return fmt.Sprintf("SELECT %s FROM %s %s %s", dialect.Quote(sqlname), tbl.Name(), whs, orderBy), args
}

//-------------------------------------------------------------------------------------------------

// Query is the low-level request method for this table.
//
// The query is logged using whatever logger is configured. If an error arises, this too is logged.
//
// The args are for any placeholder parameters in the query.
//
// The caller must call rows.Close() on the result.
func Query(tbl sqlgen2.Table, query string, args ...interface{}) (*sql.Rows, error) {
	database := tbl.Database()
	database.LogQuery(query, args...)
	rows, err := tbl.Execer().QueryContext(tbl.Ctx(), query, args...)
	return rows, database.LogIfError(err)
}


// Exec executes a modification query (insert, update, delete, etc) and returns the number of items affected.
//
// The query is logged using whatever logger is configured. If an error arises, this too is logged.
func Exec(tbl sqlgen2.Table, req require.Requirement, query string, args ...interface{}) (int64, error) {
	database := tbl.Database()
	database.LogQuery(query, args...)
	res, err := tbl.Execer().ExecContext(tbl.Ctx(), query, args...)
	if err != nil {
		return 0, database.LogError(err)
	}
	n, err := res.RowsAffected()
	return n, database.LogIfError(require.ChainErrorIfExecNotSatisfiedBy(err, req, n))
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
