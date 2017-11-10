package schema

import (
	"fmt"
	. "github.com/acsellers/inflections"
	"github.com/rickb777/sqlgen2/sqlgen/parse"
	"github.com/rickb777/sqlgen2/sqlgen/parse/exit"
	"go/types"
	"strings"
)

func Load(pkgStore parse.PackageStore, pkg, name string) (*Table, error) {
	table := new(Table)

	// local map of indexes, used for quick lookups and de-duping.
	indices := map[string]*Index{}

	//for j := 0; j < s.NumFields(); j++ {
	//	f := s.Field(j)
	//	parse.DevInfo("    f%d: name:%-10s pkg:%s type:%-10s %v,%v\n", j,
	//		f.Name(), f.Pkg().Name(), f.Type(), f.Exported(), f.Anonymous())
	//}

	table.Type = name
	table.Name = Pluralize(Underscore(table.Type))

	str, tags, ok := pkgStore.Find(pkg, name)
	if ok {
		//err := checkAllFieldsAreUnexported(str, pkg, name)
		//if err != nil {
		//	return nil, err
		//}

		parse.DevInfo("\nFound\n")
		for j := 0; j < str.NumFields(); j++ {
			tField := str.Field(j)
			parse.DevInfo("    f%d: name:%-10s pkg:%s type:%-10s %v,%v\n", j,
				tField.Name(), tField.Pkg().Name(), tField.Type(), tField.Exported(), tField.Anonymous())

			field, ok := convertLeafNodeToField(tField, pkg, name, tags, indices, table)
			if ok {
				table.Fields = append(table.Fields, field)
			}
		}
	}

	// Each leaf node in the tree is a column in the table.
	// Convert each leaf node to a Field structure.
	//for _, node := range tree.Leaves() {
	//	field, ok := convertLeafNodeToField(node, indices, table)
	//	if ok {
	//		table.Fields = append(table.Fields, field)
	//	}
	//}

	return table, nil
}

func checkAllFieldsAreUnexported(str *types.Struct, pkg, name string) error {
	var un []string
	for j := 0; j < str.NumFields(); j++ {
		f := str.Field(j)
		if !f.Exported() {
			un = append(un, f.Name())
		}
	}
	if len(un) > 0 {
		return fmt.Errorf("%s.%s cannot be mapped because it contains unexported fields %v\n", pkg, name, un)
	}
	return nil
}

func convertLeafNodeToField(leaf *types.Var, pkg, name string, tags map[string]parse.Tag, indices map[string]*Index, table *Table) (*Field, bool) {
	tp := parse.Type{Pkg: "", Name: leaf.Type().String()}
	field := &Field{Name: leaf.Name(), Type: tp}

	// Lookup the SQL column type
	field.SqlType = BLOB
	underlying := leaf.Type().Underlying()
	switch u := underlying.(type) {
	case *types.Basic:
		field.SqlType = mapKindToSqlType[u.Kind()]
		field.Type.Base = parse.Kind(u.Kind())
	case *types.Slice:
		field.Type.Base = parse.Slice
	}

	// substitute tag variables
	if tag, exists := tags[leaf.Name()]; exists {

		if tag.Skip {
			return nil, false
		}

		field.Tags = tag

		if tag.Primary {
			if table.Primary != nil {
				exit.Fail(1, "%s, %s: compound primary keys are not supported.\n",
					table.Primary.Type.Name, field.Type.Name)
			}
			table.Primary = field
		}

		if tag.Index != "" {
			index, ok := indices[tag.Index]
			if !ok {
				index = &Index{
					Name: tag.Index,
				}
				indices[index.Name] = index
				table.Index = append(table.Index, index)
			}
			index.Fields = append(index.Fields, field)
		}

		if tag.Unique != "" {
			index, ok := indices[tag.Index]
			if !ok {
				index = &Index{
					Name:   tag.Unique,
					Unique: true,
				}
				indices[index.Name] = index
				table.Index = append(table.Index, index)
			}
			index.Fields = append(index.Fields, field)
		}

		if tag.Type != "" {
			t, ok := mapStringToSqlType[tag.Type]
			if ok {
				field.SqlType = t
			}
		}

		switch tag.Encode {
		case "json":
			field.Encode = ENCJSON
			// case "gzip":
			// case "snappy":
		}
	}

	// get the full path name
	// omit table name
	//path := leaf.Path()[1:]
	var parts []string
	//for _, part := range path {
	//	if part.Tags != nil && part.Tags.Name != "" {
	//		parts = append(parts, part.Tags.Name)
	//		return nil, false
	//	}
	//
	parts = append(parts, leaf.Name())

	field.Path = PathOf(parts...)
	field.SqlName = Underscore(strings.Join(parts, "_"))

	return field, true
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
