package schema

import (
	"fmt"
	"os"
	"strings"

	. "github.com/acsellers/inflections"
	"github.com/rickb777/sqlgen/parse"
)

func Load(tree *parse.Node) *Table {
	table := new(Table)

	// local map of indexes, used for quick
	// lookups and de-duping.
	indices := map[string]*Index{}

	table.Type = tree.Type
	table.Name = Pluralize(Underscore(tree.Type))

	// each edge node in the tree is a column
	// in the table. Convert each edge node to
	// a Field structure.
	for _, node := range tree.Leaves() {
		field, ok := loadNode(node, indices, table)
		if ok {
			table.Fields = append(table.Fields, field)
		}
	}

	return table
}

func loadNode(node *parse.Node, indices map[string]*Index, table *Table) (*Field, bool) {
	field := new(Field)
	field.Name = node.Name

	// Lookup the SQL column type
	field.Type = BLOB
	if t, ok := parse.Types[node.Type]; ok {
		if tt, ok := types[t]; ok {
			field.Type = tt
		}
	}

	// substitute tag variables
	if node.Tags != nil {

		if node.Tags.Skip {
			return nil, false
		}

		field.Auto = node.Tags.Auto
		field.Primary = node.Tags.Primary
		field.Size = node.Tags.Size

			if node.Tags.Primary {
				if table.Primary != nil {
					fmt.Fprintf(os.Stderr, "%s, %s: compound primary keys are not supported.\n",
						table.Primary.Name, field.Name)
					os.Exit(1)
				}
				table.Primary = field
			}

		if node.Tags.Index != "" {
			index, ok := indices[node.Tags.Index]
			if !ok {
				index = new(Index)
				index.Name = node.Tags.Index
				indices[index.Name] = index
				table.Index = append(table.Index, index)
			}
			index.Fields = append(index.Fields, field)
		}

		if node.Tags.Unique != "" {
			index, ok := indices[node.Tags.Index]
			if !ok {
				index = new(Index)
				index.Name = node.Tags.Unique
				index.Unique = true
				indices[index.Name] = index
				table.Index = append(table.Index, index)
			}
			index.Fields = append(index.Fields, field)
		}

		if node.Tags.Type != "" {
			t, ok := sqlTypes[node.Tags.Type]
			if ok {
				field.Type = t
			}
		}
	}

	// get the full path name
	// omit table name
	path := node.Path()[1:]
	var parts []string
	for _, part := range path {
		if part.Tags != nil && part.Tags.Name != "" {
			parts = append(parts, part.Tags.Name)
			return nil, false
		}

		parts = append(parts, part.Name)
	}
	field.SqlName = strings.Join(parts, "_")
	field.SqlName = Underscore(field.SqlName)

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
