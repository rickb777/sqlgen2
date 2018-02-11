**sqlgen** generates SQL statements and database helper functions from your Go structs. It can be used in
place of a simple ORM or hand-written SQL.

See the [demo](https://github.com/rickb777/sqlgen2/tree/master/demo) directory for examples. Look in the
generated files `*_sql.go` and the hand-crafted files (`hook.go`, `issue.go`, `user.go`).

## Features

* Auto-generates DAO-style table-support code for SQL databases.
* Sophisticated parsing of a Go `struct` that describes records in the table.
* Allows nesting of structs, fields that are structs or pointers to structs etc.
* Struct tags give fine control over the semantics of each field.
* Supports indexes and constraints.
* Supports foreign key relationships between tables.
* Helps you develop code for joins and views.
* Supports JSON-encoded columns, allowing a more no-SQL model when needed.
* Provides a builder-style API for constructing where-clauses and query constraints.
* Allows declarative requirements on the expected result of each query, enhancing error checking. 
* Very flexible configuration.
* Fast and easy to use.

Currently, support is included for **MySQL**, **PostgreSQL** and **SQLite**. Other dialects can be added relatively easy - send a Pull Request!

## Install

Install or upgrade with this command:

```
go get -u github.com/rickb777/sqlgen2
```

## Usage

```
sqlgen [option [arg]] file.go ...

Usage of sqlgen:

  -prefix string
    	Prefix for names of generated types; optional. Use this if you need to avoid name collisions.

  -type string
    	The type to analyse; required. This is expressed in the form 'pkg.Name'

  -kind string
    	Kind of model: you could use 'Table', 'View', 'Join' etc as required (default "Table")

  -list string
    	List type for slice of model objects; optional.

  -tags string
    	A YAML file containing tags that augment and override any in the Go struct(s); optional. Tags control the SQL type, size, column name, indexes etc.

  -o string
    	Output file name; optional. Use '-' for stdout. If omitted, the first input filename is used with '_sql.go' suffix.

  -gofmt
    	Format and simplify the generated code nicely.

Options:

  -schema
    	Generate SQL schema create/drop methods.

  -exec
        Generate Exec method. This is also provided with -update or -delete

  -read
  -select
    	Generate SQL select (read) methods.

  -create
  -insert
    	Generate SQL create (insert) methods.

  -update
    	Generate SQL update methods.

  -delete
    	Generate SQL delete methods.

  -slice
    	Generate SQL slice (column select) methods.

  -all
    	Shorthand for '-schema -create -read -update -delete -slice'; recommended. This does not affect -setters.

  -setters string
    	Generate setters for fields of your type (see -type): none, optional, exported, all. Fields that are pointers are assumed to be optional. (default "none")

Information:

  -v	Show progress messages.
  -z	Show debug messages.
  -ast
    	Trace the whole astract syntax tree (very verbose).
```

The options `-schema`, `-read`, `-insert`, -`update`, `-delete`, and `-slice` are all off by default; **you should normally include `-all`**, unless you have other needs.

## Benchmarks

This tool demonstrates performance gains, albeit small, over light-weight ORM packages such as `sqlx` and `meddler`. Over time I plan to expand the benchmarks to include additional ORM packages.

To run the project benchmarks:

```
go get ./...
go generate ./...
go build
cd bench
go test -bench=Bench
```

Example selecing a single row:

```
BenchmarkMeddlerRow-4      30000        42773 ns/op
BenchmarkSqlxRow-4         30000        41554 ns/op
BenchmarkSqlgenRow-4       50000        39664 ns/op

```

Selecting multiple rows:

```
BenchmarkMeddlerRows-4      2000      1025218 ns/op
BenchmarkSqlxRows-4         2000       807213 ns/op
BenchmarkSqlgenRows-4       2000       700673 ns/op
```


## Restrictions

* Compound primaries are not supported
* In the structs used for tables, the imports must not use '.' or be renamed.


## Credits

This tool was derived from [sqlgen](https://github.com/drone/sqlgen) by drone.io, which was itself
inspired by [scaneo](https://github.com/variadico/scaneo).
