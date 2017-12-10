package schema

import "go/types"

var mapTagToEncoding = map[string]SqlEncode{
	"":     ENCNONE,
	"json": ENCJSON,
	"text": ENCTEXT,
}

// convert Go types to SQL types.
var mapKindToSqlType = map[types.BasicKind]SqlType{
	types.Bool:       BOOLEAN,
	types.Int:        INTEGER,
	types.Int8:       INTEGER,
	types.Int16:      INTEGER,
	types.Int32:      INTEGER,
	types.Int64:      INTEGER,
	types.Uint:       INTEGER,
	types.Uint8:      INTEGER,
	types.Uint16:     INTEGER,
	types.Uint32:     INTEGER,
	types.Uint64:     INTEGER,
	types.Float32:    REAL,
	types.Float64:    REAL,
	types.Complex64:  BLOB,
	types.Complex128: BLOB,
	//types.Interface:  BLOB,
	//types.Bytes:      BLOB,
	types.String: VARCHAR,
	//types.Map:        BLOB,
	//types.Slice:      BLOB,
}

var mapStringToSqlType = map[string]SqlType{
	"text":     VARCHAR,
	"varchar":  VARCHAR,
	"varchar2": VARCHAR,
	"number":   INTEGER,
	"integer":  INTEGER,
	"int":      INTEGER,
	"blob":     BLOB,
	"bytea":    BLOB,
	"json":     JSON,
}
