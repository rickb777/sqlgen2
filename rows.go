package sqlgen2

import (
	"database/sql"
	"github.com/rickb777/sqlgen2/util"
)

// Rows provides a tool for scanning result *sql.Rows of arbitrary or varying length.
// The internal *sql.Rows field is exported and is usable as per normal via its Next
// and Scan methods, or the Next and ScanToMap methods can be used instead.
type Rows struct {
	cols  []string
	types []*sql.ColumnType
	Rows  *sql.Rows
}

// RowData holds a single row result from the database.
type RowData struct {
	Columns     []string
	ColumnTypes []*sql.ColumnType
	Data        util.StringAnyMap
}

// WrapRows wraps a *sql.Rows result so that its data can be scanned into a series of
// maps, one for each row.
func WrapRows(rows *sql.Rows) (*Rows, error) {
	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	types, err := rows.ColumnTypes()
	if err != nil {
		return nil, err
	}

	return &Rows{cols, types, rows}, nil
}

// Next prepares the next result row for reading with the Scan method. It
// returns true on success, or false if there is no next result row or an error
// happened while preparing it. Err should be consulted to distinguish between
// the two cases.
//
// Every call to Scan, even the first one, must be preceded by a call to Next.
func (rams *Rows) Next() bool {
	return rams.Rows.Next()
}

// Scan copies the columns in the current row into the values pointed
// at by dest. The number of values in dest must be the same as the
// number of columns in the wrapped Rows.
func (rams *Rows) Scan(dest ...interface{}) error {
	return rams.Rows.Scan(dest...)
}

// ScanToMap copies all the column data of the current row into a map. The
// map is keyed by column name.
//
// The result describes a single row from the database, consisting of the
// column names, types and data.
func (rams *Rows) ScanToMap() (RowData, error) {
	values := make([]interface{}, len(rams.cols))
	valuePointers := make([]interface{}, len(rams.cols))
	for i, _ := range values {
		valuePointers[i] = &values[i]
	}

	if err := rams.Rows.Scan(valuePointers...); err != nil {
		return RowData{}, err
	}

	m := make(util.StringAnyMap)
	for i, colName := range rams.cols {
		v := valuePointers[i].(*interface{})
		m[colName] = *v
	}

	return RowData{rams.cols, rams.types, m}, nil
}

// Close closes the Rows, preventing further enumeration. If Next is called
// and returns false and there are no further result sets,
// the Rows are closed automatically and it will suffice to check the
// result of Err. Close is idempotent and does not affect the result of Err.
func (rams *Rows) Close() error {
	return rams.Rows.Close()
}

// Err returns the error, if any, that was encountered during iteration.
// Err may be called after an explicit or implicit Close.
func (rams *Rows) Err() error {
	return rams.Rows.Err()
}

// Columns returns the column names.
func (rams *Rows) Columns() ([]string, error) {
	return rams.cols, nil
}

// ColumnTypes returns column information such as column type, length,
// and nullable. Some information may not be available from some drivers.
func (rams *Rows) ColumnTypes() ([]*sql.ColumnType, error) {
	return rams.Rows.ColumnTypes()
}
