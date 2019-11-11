package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/kortschak/utter"
	"github.com/rickb777/filemod"
	"github.com/rickb777/sqlapi/schema"
	"github.com/rickb777/sqlapi/types"
	"github.com/rickb777/sqlgen2/code"
	"github.com/rickb777/sqlgen2/load"
	"github.com/rickb777/sqlgen2/output"
	"github.com/rickb777/sqlgen2/parse"
	"github.com/rickb777/sqlgen2/parse/exit"
	"gopkg.in/yaml.v2"
)

func main() {
	start := time.Now()

	var oFile, pkgImport, typeName, prefix, list, kind, tableName, tagsFile, genSetters string
	var flags = load.FuncFlags{}
	var force, pgx, all, join, read, create, gofmt, jsonFile, yamlFile, showVersion bool

	flag.StringVar(&oFile, "o", "", "Output file name; optional. Use '-' for stdout.\n"+
		"\tIf omitted, the first input filename is used with '_sql.go' suffix.")
	flag.StringVar(&pkgImport, "pkg", "", "The package holding the model type.\n"+
		"\tThis is expressed in the form 'github.com/owner/repo/foo/bar'\n"+
		"\tIt is only required when the output file is in a different package.")
	flag.StringVar(&typeName, "type", "", "The type to analyse; required.\n"+
		"\tThis is expressed in the form 'pkg.Name'")
	flag.StringVar(&prefix, "prefix", "", "Prefix for names of generated types; optional.\n"+
		"\tUse this if you need to avoid name collisions.")
	flag.StringVar(&list, "list", "", "List type for slice of model objects; optional.")
	flag.StringVar(&kind, "kind", "Table", "Kind of model: you could use 'Table', 'View', 'Join' etc as required.")
	flag.StringVar(&tableName, "table", "", "The name for the database table; default is based on the struct name as a plural.")
	flag.StringVar(&tagsFile, "tags", "", "A YAML file containing tags that augment and override any in the Go struct(s); optional.\n"+
		"\tTags control the SQL type, size, column name, indexes etc.")

	// filters for what gets generated
	flag.BoolVar(&force, "f", false, "Force output code generation, even if already up to date.")
	flag.BoolVar(&pgx, "pgx", false, "Generates code for github.com/jackc/pgx.")
	flag.BoolVar(&all, "all", false, "Shorthand for '-schema -exec -query -select -count -insert -update -upsert -delete -slice'; recommended for normal tables.\n"+
		"\tThis does not affect -setters.")
	flag.BoolVar(&join, "join", false, "Shorthand for '-query'; recommended for joins.\n"+
		"\tThis does not affect -setters.")
	flag.BoolVar(&read, "read", false, "Alias for -select")
	flag.BoolVar(&create, "create", false, "Alias for -insert")
	flag.BoolVar(&flags.Schema, "schema", false, "Generate SQL schema create/drop methods.")
	flag.BoolVar(&flags.Insert, "insert", false, "Generate SQL insert (create) methods.")
	flag.BoolVar(&flags.Exec, "exec", false, "Generate Exec method. This is also provided with -update or -delete.")
	flag.BoolVar(&flags.Query, "query", false, "Generate basic SQL query methods.")
	flag.BoolVar(&flags.Select, "select", false, "Generate SQL select (read) methods; also enables -query, -count.")
	flag.BoolVar(&flags.Count, "count", false, "Generate SQL count methods.")
	flag.BoolVar(&flags.Update, "update", false, "Generate SQL update methods.")
	flag.BoolVar(&flags.Upsert, "upsert", false, "Generate SQL upsert methods; ignored if there is no primary key.")
	flag.BoolVar(&flags.Delete, "delete", false, "Generate SQL delete methods.")
	flag.BoolVar(&flags.Slice, "slice", false, "Generate SQL slice (column select) methods.")
	flag.BoolVar(&flags.Scan, "scan", false, "Generate exported row scan functions (these are normally unexported).")
	flag.StringVar(&genSetters, "setters", "none", "Generate setters for fields of your type (see -type): none, optional, exported, all.\n"+
		"\tFields that are pointers are assumed to be optional.")

	flag.BoolVar(&output.Verbose, "v", false, "Show progress messages.")
	flag.BoolVar(&parse.Debug, "z", false, "Show debug messages.")
	flag.BoolVar(&parse.PrintAST, "ast", false, "Trace the whole astract syntax tree (very verbose).")
	flag.BoolVar(&gofmt, "gofmt", false, "Format and simplify the generated code nicely.")
	flag.BoolVar(&jsonFile, "json", false, "Read/print the table description in JSON (overrides Go parsing if the JSON file exists).")
	flag.BoolVar(&yamlFile, "yaml", false, "Read/print the table description in YAML (overrides Go parsing if the YAML file exists).")
	flag.BoolVar(&showVersion, "version", false, "Show the version.")

	flag.Parse()

	if showVersion {
		fmt.Println(appVersion)
		os.Exit(0)
	}

	output.Require(flag.NArg() > 0, "At least one input file (or path) is required; put this after the other arguments.\n")

	if read {
		flags.Select = true
	}

	if flags.Select {
		flags.Count = true
		flags.Query = true
	}

	if create {
		flags.Insert = true
	}

	if flags.Upsert {
		flags.Insert = true
		flags.Update = true
	}

	if all {
		flags = load.AllFuncFlags
	} else if join {
		flags = load.FuncFlags{Query: true}
	}

	output.Require(len(typeName) > 3, "-type is required. This must specify a type, qualified with its local package in the form 'pkg.Name'.\n", typeName)
	words := strings.Split(typeName, ".")
	output.Require(len(words) == 2, "type %q requires a package name prefix (as in 'time.Time').\n", typeName)

	pkg, name := words[0], words[1]
	typePkg := pkg
	tablePkg := pkg

	if oFile == "" {
		oFile = flag.Args()[0]
		output.Require(strings.HasSuffix(oFile, ".go"), oFile+": must end '.go'")
		oFile = oFile[:len(oFile)-3] + "_sql.go"
		parse.DevInfo("oFile: %s\n", oFile)
	} else {
		tablePkg = load.LastDirName(oFile)
		output.Require(tablePkg == typePkg || pkgImport != "", typeName+": must be accompanied with -pkg if the output file is in a different directory")
		parse.DevInfo("typePkg: %s, tablePkg: %s\n", typePkg, tablePkg)
	}

	outputDeps := filemod.New(oFile)
	o := output.NewOutput(oFile)

	var table *schema.TableDescription

	if yamlFile {
		buf := &bytes.Buffer{}
		dec := yaml.NewDecoder(buf)
		table = readTableJson(o.Derive(".yml"), buf, dec)
		if table != nil {
			yamlFile = false
		}
	} else if jsonFile {
		buf := &bytes.Buffer{}
		dec := json.NewDecoder(buf)
		table = readTableJson(o.Derive(".json"), buf, dec)
		if table != nil {
			jsonFile = false
		}
	}

	inputSources := filemod.New(flag.Args()...)
	if outputDeps.Compare(inputSources) == filemod.AllAreYounger {
		output.Info("skipped %s (%s is already up to date)\n", strings.Join(flag.Args(), ", "), oFile)
		return
	}

	if table == nil {
		output.Info("parsing %s\n", strings.Join(flag.Args(), ", "))

		// parse the Go source code file(s) to extract the required struct and return it as an AST.
		pkgStore, err := parse.Parse(flag.Args())
		output.Require(err == nil, "%v\n", err)
		//utter.Dump(pkgStore)

		tags, err := types.ReadTagsFile(tagsFile)
		if err != nil && !os.IsNotExist(err) {
			exit.Fail(1, "tags file %s failed: %s.\n", tagsFile, err)
		}

		output.Info("loading %s\n", strings.Join(flag.Args(), ", "))

		// load the Tree into a schema Object
		table, err = load.Load(pkgStore, parse.LType{PkgName: pkg, Name: name}, typePkg, tags)
		if err != nil {
			exit.Fail(1, "Go parser failed: %v.\n", err)
		}

		if parse.Debug {
			utter.Dump(table)
		}

		if len(table.Fields) < 1 {
			exit.Fail(1, "no fields found. Check earlier parser warnings.\n")
		}
	} else {
		output.Info("ignored %s\n", strings.Join(flag.Args(), ", "))
	}

	if yamlFile {
		buf := &bytes.Buffer{}
		enc := yaml.NewEncoder(buf)
		writeTableJson(o.Derive(".yml"), buf, enc, table)
	}

	if jsonFile {
		buf := &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		enc.SetIndent("", "  ")
		writeTableJson(o.Derive(".json"), buf, enc, table)
	}

	if typePkg == tablePkg {
		typePkg = ""
		pkgImport = "" // no import needed
	}

	writeSqlGo(o, name, prefix, tableName, kind, list, pkgImport, typePkg, tablePkg, genSetters, table, flags, pgx, gofmt)

	output.Info("%s took %v\n", o.Path(), time.Now().Sub(start))
}

func readTableJson(o output.Output, buf io.ReadWriter, dec decoder) *schema.TableDescription {
	err := o.ReadTo(buf)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		exit.Fail(1, "YAML reader from %s failed: %v.\n", o.Path(), err)
	}

	output.Info("reading %s\n", o.Path())
	table := schema.TableDescription{}
	err = dec.Decode(&table)
	if err != nil {
		exit.Fail(1, "YAML reader from %s failed: %v.\n", o.Path(), err)
	}
	return &table
}

func writeTableJson(o output.Output, buf io.ReadWriter, enc encoder, table *schema.TableDescription) {
	err := enc.Encode(table)
	if err != nil {
		exit.Fail(1, "YAML writer to %s failed: %v.\n", o.Path(), err)
	}

	err = o.Write(buf)
	if err != nil {
		exit.Fail(1, "YAML writer to %s failed: %v.\n", o.Path(), err)
	}
}

func writeSqlGo(o output.Output, name, prefix, tableName, kind, list, pkgImport, typePkg, tablePkg, genSetters string, table *schema.TableDescription, flags load.FuncFlags, pgx, gofmt bool) {
	sql := "sql"
	api := "sqlapi"
	if pgx {
		sql = "pgx"
		api = "pgxapi"
	}
	if typePkg != "" {
		typePkg += "."
	}
	view := code.NewView(typePkg, tablePkg, name, prefix, tableName, list, sql, api)
	view.Table = table
	view.Thing = kind
	view.Thinger = ender(kind)
	view.Interface1 = api + "." + load.PrimaryInterface(table, flags.Schema)
	if flags.Scan {
		view.Scan = "Scan"
	}

	setters := view.FilterSetters(genSetters)

	importSet := load.PackagesToImport(flags, pgx)
	if pkgImport != "" {
		importSet.Add(pkgImport)
	}

	code.ImportsForFields(table, importSet)
	code.ImportsForSetters(setters, importSet)

	headerBuf := &bytes.Buffer{}
	tablerBuf := &bytes.Buffer{}
	queryerBuf := &bytes.Buffer{}
	structBuf := &bytes.Buffer{}

	code.WritePackageHeader(headerBuf, tablePkg, appVersion)

	code.WriteImports(headerBuf, importSet)

	code.WriteType(tablerBuf, queryerBuf, structBuf, view)

	code.WritePrimaryDeclarations(structBuf, view)

	if flags.Schema {
		code.WriteSchemaDeclarations(structBuf, view)
		code.WriteSchemaFunctions(tablerBuf, structBuf, view)
	}

	if flags.Exec || flags.Update || flags.Delete {
		code.WriteExecFunc(queryerBuf, structBuf, view)
	}

	if flags.Query {
		code.WriteQueryRows(queryerBuf, structBuf, view)
		code.WriteQueryThings(queryerBuf, structBuf, view)
	}

	code.WriteScanRows(structBuf, view)

	if flags.Select {
		code.WriteGetRow(queryerBuf, structBuf, view)
		code.WriteSelectRowsFuncs(queryerBuf, structBuf, view)
	}

	if flags.Count {
		code.WriteCountRowsFuncs(queryerBuf, structBuf, view)
	}

	if flags.Slice {
		code.WriteSliceColumn(queryerBuf, structBuf, view)
	}

	if flags.Insert {
		code.WriteConstructInsert(structBuf, view)
	}

	if flags.Update {
		code.WriteConstructUpdate(structBuf, view)
	}

	if flags.Insert {
		code.WriteInsertFunc(queryerBuf, structBuf, view)
	}

	if flags.Update {
		code.WriteUpdateFunc(queryerBuf, structBuf, view)
	}

	if flags.Upsert {
		code.WriteUpsertFunc(queryerBuf, structBuf, view)
	}

	if flags.Delete {
		code.WriteDeleteFunc(queryerBuf, structBuf, view)
	}

	code.WriteSetters(queryerBuf, structBuf, view, setters)

	io.WriteString(tablerBuf, "}\n")
	io.WriteString(queryerBuf, "}\n")

	finishWriting(o, gofmt, constructFullBuffer(headerBuf, tablerBuf, queryerBuf, structBuf))
}

func constructFullBuffer(headerBuf, tablerBuf, queryerBuf, structBuf io.Reader) *bytes.Buffer {
	whole := &bytes.Buffer{}
	io.Copy(whole, headerBuf)
	io.Copy(whole, tablerBuf)
	io.Copy(whole, queryerBuf)
	io.Copy(whole, structBuf)
	return whole
}

func finishWriting(o output.Output, gofmt bool, buf *bytes.Buffer) {
	var pretty io.Reader = buf
	if gofmt {
		var err error
		pretty, err = load.GoFmt(buf)
		output.Require(err == nil, "%s\n%v\n", buf.String(), err)
	}

	o.Write(pretty)
}

func ender(s string) string {
	last := lastLetter(s)
	switch last {
	case 'e', 'o':
		return s + "r"
	}
	return s + "er"
}

func lastLetter(s string) byte {
	if s == "" {
		return 0
	}
	return s[len(s)-1]
}

type encoder interface {
	Encode(v interface{}) (err error)
}

type decoder interface {
	Decode(v interface{}) (err error)
}
