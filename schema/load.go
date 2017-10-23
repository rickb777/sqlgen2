package schema

import (
	"strings"

	. "github.com/acsellers/inflections"
	"github.com/rickb777/sqlgen2/sqlgen/parse"
	"github.com/rickb777/sqlgen2/sqlgen/parse/exit"
)

func Load(tree *parse.Node) *Table {
	table := new(Table)

	// local map of indexes, used for quick
	// lookups and de-duping.
	indices := map[string]*Index{}

	table.Type = tree.Type.Name
	table.Name = Pluralize(Underscore(tree.Type.Name))

	// Each leaf node in the tree is a column in the table.
	// Convert each leaf node to a Field structure.
	for _, node := range tree.Leaves() {
		field, ok := convertLeafNodeToField(node, indices, table)
		if ok {
			table.Fields = append(table.Fields, field)
		}
	}

	return table
}

func convertLeafNodeToField(leaf *parse.Node, indices map[string]*Index, table *Table) (*Field, bool) {
	field := &Field{Name: leaf.Name, Type: leaf.Type}

	// Lookup the SQL column type
	field.SqlType = BLOB
	if leaf.Type.Base.IsSimpleType() {
		field.SqlType = types[leaf.Type.Base]
	}

	// substitute tag variables
	if leaf.Tags != nil {

		if leaf.Tags.Skip {
			return nil, false
		}

		field.Auto = leaf.Tags.Auto
		field.Primary = leaf.Tags.Primary
		field.Size = leaf.Tags.Size

		if leaf.Tags.Primary {
			if table.Primary != nil {
				exit.Fail(1, "%s, %s: compound primary keys are not supported.\n",
					table.Primary.Type.Name, field.Type.Name)
			}
			table.Primary = field
		}

		if leaf.Tags.Index != "" {
			index, ok := indices[leaf.Tags.Index]
			if !ok {
				index = &Index{
					Name: leaf.Tags.Index,
				}
				indices[index.Name] = index
				table.Index = append(table.Index, index)
			}
			index.Fields = append(index.Fields, field)
		}

		if leaf.Tags.Unique != "" {
			index, ok := indices[leaf.Tags.Index]
			if !ok {
				index = &Index{
					Name:   leaf.Tags.Unique,
					Unique: true,
				}
				indices[index.Name] = index
				table.Index = append(table.Index, index)
			}
			index.Fields = append(index.Fields, field)
		}

		if leaf.Tags.Type != "" {
			t, ok := sqlTypes[leaf.Tags.Type]
			if ok {
				field.SqlType = t
			}
		}

		switch leaf.Tags.Encode {
		case "json":
			field.Encode = ENCJSON
			// case "gzip":
			// case "snappy":
		}
	}

	// get the full path name
	// omit table name
	path := leaf.Path()[1:]
	var parts []string
	for _, part := range path {
		if part.Tags != nil && part.Tags.Name != "" {
			parts = append(parts, part.Tags.Name)
			return nil, false
		}

		parts = append(parts, part.Name)
	}

	field.Path = parts
	field.SqlName = Underscore(strings.Join(parts, "_"))

	return field, true
}

// convert Go types to SQL types.
var types = map[parse.Kind]SqlType{
	parse.Bool:       BOOLEAN,
	parse.Int:        INTEGER,
	parse.Int8:       INTEGER,
	parse.Int16:      INTEGER,
	parse.Int32:      INTEGER,
	parse.Int64:      INTEGER,
	parse.Uint:       INTEGER,
	parse.Uint8:      INTEGER,
	parse.Uint16:     INTEGER,
	parse.Uint32:     INTEGER,
	parse.Uint64:     INTEGER,
	parse.Float32:    INTEGER,
	parse.Float64:    INTEGER,
	parse.Complex64:  INTEGER,
	parse.Complex128: INTEGER,
	parse.Interface:  BLOB,
	parse.Bytes:      BLOB,
	parse.String:     VARCHAR,
	parse.Map:        BLOB,
	parse.Slice:      BLOB,
}

var sqlTypes = map[string]SqlType{
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
