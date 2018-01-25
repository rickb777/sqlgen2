package main

import (
	. "github.com/rickb777/sqlgen2/schema"
	"github.com/rickb777/sqlgen2/sqlgen/parse"
	"sort"
	"strings"
)

var mapTagToEncoding = map[string]SqlEncode{
	"":       ENCNONE,
	"json":   ENCJSON,
	"text":   ENCTEXT,
	"driver": ENCDRIVER,
}

var mapStringToSqlType = map[string]parse.Kind{
	// Go-flavour names
	"bool":    parse.Bool,
	"int":     parse.Int,
	"int8":    parse.Int8,
	"int16":   parse.Int16,
	"int32":   parse.Int32,
	"int64":   parse.Int64,
	"uint":    parse.Uint,
	"uint8":   parse.Uint8,
	"uint16":  parse.Uint16,
	"uint32":  parse.Uint32,
	"uint64":  parse.Uint64,
	"float32": parse.Float32,
	"float64": parse.Float64,
	"string":  parse.String,

	// SQL-flavour names
	"text":     parse.String,
	"json":     parse.String,
	"varchar":  parse.String,
	"varchar2": parse.String,
	"number":   parse.Int,
	"tinyint":  parse.Int8,
	"smallint": parse.Int16,
	"integer":  parse.Int,
	"bigint":   parse.Int64,
	"blob":     parse.Struct,
	"bytea":    parse.Struct,
}

func allowedSqlTypeStrings() string {
	var s []string
	for k, _ := range mapStringToSqlType {
		s = append(s, k)
	}
	sort.Strings(s)
	return strings.Join(s, ", ")
}
