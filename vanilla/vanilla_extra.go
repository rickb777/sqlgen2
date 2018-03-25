package vanilla

import (
	"github.com/rickb777/sqlgen2/require"
	"github.com/rickb777/sqlgen2/where"
	"github.com/rickb777/sqlgen2/support"
)

// SliceStringColumn gets a string column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl RecordTable) SliceStringColumn(req require.Requirement, column string, wh where.Expression, qc where.QueryConstraint) ([]string, error) {
	return support.SliceStringList(tbl, req, column, wh, qc)
}

// SliceIntColumn gets an int column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl RecordTable) SliceIntColumn(req require.Requirement, column string, wh where.Expression, qc where.QueryConstraint) ([]int, error) {
	return support.SliceIntList(tbl, req, column, wh, qc)
}

// SliceUintColumn gets a uint column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl RecordTable) SliceUintColumn(req require.Requirement, column string, wh where.Expression, qc where.QueryConstraint) ([]uint, error) {
	return support.SliceUintList(tbl, req, column, wh, qc)
}

// SliceInt64Column gets an int64 column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl RecordTable) SliceInt64Column(req require.Requirement, column string, wh where.Expression, qc where.QueryConstraint) ([]int64, error) {
	return support.SliceInt64List(tbl, req, column, wh, qc)
}

// SliceUint64Column gets a uint64 column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl RecordTable) SliceUint64Column(req require.Requirement, column string, wh where.Expression, qc where.QueryConstraint) ([]uint64, error) {
	return support.SliceUint64List(tbl, req, column, wh, qc)
}

// SliceInt32Column gets an int32 column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl RecordTable) SliceInt32Column(req require.Requirement, column string, wh where.Expression, qc where.QueryConstraint) ([]int32, error) {
	return support.SliceInt32List(tbl, req, column, wh, qc)
}

// SliceUint32Column gets a uint32 column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl RecordTable) SliceUint32Column(req require.Requirement, column string, wh where.Expression, qc where.QueryConstraint) ([]uint32, error) {
	return support.SliceUint32List(tbl, req, column, wh, qc)
}

// SliceInt16Column gets an int16 column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl RecordTable) SliceInt16Column(req require.Requirement, column string, wh where.Expression, qc where.QueryConstraint) ([]int16, error) {
	return support.SliceInt16List(tbl, req, column, wh, qc)
}

// SliceUint16Column gets a uint16 column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl RecordTable) SliceUint16Column(req require.Requirement, column string, wh where.Expression, qc where.QueryConstraint) ([]uint16, error) {
	return support.SliceUint16List(tbl, req, column, wh, qc)
}

// SliceInt8Column gets an int8 column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl RecordTable) SliceInt8Column(req require.Requirement, column string, wh where.Expression, qc where.QueryConstraint) ([]int8, error) {
	return support.SliceInt8List(tbl, req, column, wh, qc)
}

// SliceUint8Column gets a uint8 column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl RecordTable) SliceUint8Column(req require.Requirement, column string, wh where.Expression, qc where.QueryConstraint) ([]uint8, error) {
	return support.SliceUint8List(tbl, req, column, wh, qc)
}

// SliceFloat32Column gets a float32 column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl RecordTable) SliceFloat32Column(req require.Requirement, column string, wh where.Expression, qc where.QueryConstraint) ([]float32, error) {
	return support.SliceFloat32List(tbl, req, column, wh, qc)
}

// SliceFloat64Column gets a float64 column for all rows that match the 'where' condition.
// Any order, limit or offset clauses can be supplied in query constraint 'qc'.
// Use nil values for the 'wh' and/or 'qc' arguments if they are not needed.
func (tbl RecordTable) SliceFloat64Column(req require.Requirement, column string, wh where.Expression, qc where.QueryConstraint) ([]float64, error) {
	return support.SliceFloat64List(tbl, req, column, wh, qc)
}

