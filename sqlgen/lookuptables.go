package main

import (
	. "github.com/rickb777/sqlapi/schema"
	"github.com/rickb777/sqlapi/types"
	"sort"
	"strings"
)

var mapTagToEncoding = map[string]SqlEncode{
	"":       ENCNONE,
	"json":   ENCJSON,
	"text":   ENCTEXT,
	"driver": ENCDRIVER,
}

var mapStringToSqlType = map[string]types.Kind{
	// Go-flavour names
	"bool":    types.Bool,
	"int":     types.Int,
	"int8":    types.Int8,
	"int16":   types.Int16,
	"int32":   types.Int32,
	"int64":   types.Int64,
	"uint":    types.Uint,
	"uint8":   types.Uint8,
	"uint16":  types.Uint16,
	"uint32":  types.Uint32,
	"uint64":  types.Uint64,
	"float32": types.Float32,
	"float64": types.Float64,
	"string":  types.String,

	// SQL-flavour names
	"text":     types.String,
	"json":     types.String,
	"varchar":  types.String,
	"varchar2": types.String,
	"number":   types.Int,
	"tinyint":  types.Int8,
	"smallint": types.Int16,
	"integer":  types.Int,
	"bigint":   types.Int64,
	"blob":     types.Struct,
	"bytea":    types.Struct,
}

func allowedSqlTypeStrings() string {
	var s []string
	for k, _ := range mapStringToSqlType {
		s = append(s, k)
	}
	sort.Strings(s)
	return strings.Join(s, ", ")
}
