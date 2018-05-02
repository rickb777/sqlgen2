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
go get -u github.com/rickb777/sqlgen2/sqlgen
```

This will fetch the source code, compile it, and leave a `sqlgen` binary in your bin folder ready to use.

You will also need to import the `sqlgen2` package and other sub-packages in your source code.


## Usage

See the [**command line usage**](docs/usage.md).


## Tutorial

See the [**tutorial**](docs/tutorial.md).


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

Example selecting a single row:

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

CAUTION - these figures might not be up to date.


## Credits

This tool was derived from [sqlgen](https://github.com/drone/sqlgen) by drone.io, which was itself
inspired by [scaneo](https://github.com/variadico/scaneo).
