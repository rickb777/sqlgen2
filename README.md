# sqlgen2

[![GoDoc](https://img.shields.io/badge/api-Godoc-blue.svg)](https://pkg.go.dev/github.com/rickb777/sqlgen2)
[![Build Status](https://travis-ci.org/rickb777/sqlgen2.svg?branch=master)](https://travis-ci.org/rickb777/sqlgen2)
[![Issues](https://img.shields.io/github/issues/rickb777/sqlgen2.svg)](https://github.com/rickb777/sqlgen2/issues)

**sqlgen** generates SQL statements from your Go structs, used along with
[**sqlapi** database helper functions](https://github.com/rickb777/sqlapi). It takes the
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

Currently, support is included for **MySQL**, **PostgreSQL** and **SQLite**. Other `database/sql` dialects can be added relatively easy - send a Pull Request!

There is also direct support for `pgx`, which differs somewhat from the `database/sql` but is noteworthy and worth supporting here.

## Install

 -- this is being revised --

## Usage

See the [**command line usage**](docs/usage.md).


## Tutorial

See the [**tutorial**](docs/tutorial.md).


## Credits

This tool was derived from [sqlgen](https://github.com/drone/sqlgen) by drone.io, which was itself
inspired by [scaneo](https://github.com/variadico/scaneo).
