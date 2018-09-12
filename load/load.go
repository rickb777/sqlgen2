package load

import (
	"sort"
	"github.com/rickb777/sqlapi/schema"
	"github.com/rickb777/sqlgen2/parse"
	stypes "github.com/rickb777/sqlapi/types"
	"github.com/acsellers/inflections"
	"github.com/rickb777/sqlgen2/parse/exit"
	"github.com/kortschak/utter"
)

type context struct {
	pkgStore parse.PackageStore
	indices  map[string]*schema.Index
	table    *schema.TableDescription
	mainPkg  string
	fileTags stypes.Tags
}

func Load(pkgStore parse.PackageStore, name parse.LType, mainPkg string, fileTags stypes.Tags) (*schema.TableDescription, error) {
	table := new(schema.TableDescription)

	// local map of indexes, used for quick lookups and de-duping.
	indices := map[string]*schema.Index{}

	table.Type = name.Name
	table.Name = inflections.Pluralize(inflections.Underscore(table.Type))

	nm := pkgStore.FindNamed(name)
	tags := pkgStore.FindTags(name)

	ctx := &context{pkgStore, indices, table, mainPkg, fileTags}
	ctx.examineStruct(nm, name, tags, nil)

	for _, idx := range ctx.indices {
		table.Index = append(table.Index, idx)
	}

	sort.Slice(table.Index, func(i, j int) bool {
		return table.Index[i].Name < table.Index[j].Name
	})

	checkNoConflictingNames(name, table)

	return table, nil
}

func mergeTags(structTags, fileTags stypes.Tags) stypes.Tags {
	merged := make(stypes.Tags)
	for n, t := range structTags {
		merged[n] = t
	}
	for n, t := range fileTags {
		merged[n] = t
	}
	parse.DevInfo("merged tags\n-----------\n%s\n", merged.String())
	return merged
}

func checkNoConflictingNames(name parse.LType, table *schema.TableDescription) {
	names := make(map[string]struct{})
	var duplicates schema.Identifiers

	for _, name := range table.Fields.SqlNames() {
		_, exists := names[name]
		if exists {
			duplicates = append(duplicates, name)
		}
		names[name] = struct{}{}
	}

	if len(duplicates) > 0 {
		parse.DevInfo("checkNoConflictingNames %s %+v\n", name, utter.Sdump(table))
		exit.Fail(1, "%s: found conflicting SQL column names: %s.\nPlease set the names on these fields explicitly using tags.\n",
			name, duplicates.MkString(", "))
	}
}


