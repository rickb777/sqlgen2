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

## Tutorial

`sqlgen` gives you a table-focussed view of your database. For each table, you write a `struct` that describes the table's columns. It generates code that maps your `struct` onto its database table (this also works for views and join results).

`sqlgen` is *not* an ORM: it does not automatically handle 'object' - relational mapping. It's simpler -
but just as comprehensive.

You are expected to understand your own database; `sqlgen` removes the hard work of reliably writing the necessary `database/sql` API calls.

First, let's start with a simple `User` struct in `user.go`:

```Go
type User struct {
	ID     int64
	Login  string
	Email  string
}
```

We can run the following command:

```
sqlgen -file user.go -type User -pkg demo
```

The tool outputs the following generated code (slightly simplified for clarity):

```Go
func scanUsers(rows *sql.Rows, firstOnly bool) (vv []*User, n int64, err error) {
	for rows.Next() {
		n++

        var v0 int64
        var v1 string
        var v2 string

        err := row.Scan(
            &v0,
            &v1,
            &v2,
        )
        if err != nil {
            return nil, err
        }
    
        v := &User{}
        v.ID = v0
        v.Login = v1
        v.Email = v2

		vv = append(vv, v)
	}

	return vv, n, rows.Err()
}

const CreateUserStmt = `
CREATE TABLE IF NOT EXISTS users (
 id     INTEGER,
 login  TEXT,
 email  TEXT
);
`

// many other functions and sql statements not displayed
```

This is a great start, but what if we want to specify primary keys, column sizes and more? This may be acheived by annotating your code using Go tags. For example, we can tag the `ID` field to indicate it is a primary key and will auto increment:

```diff
type User struct {
-   ID      int64
+   ID      int64  `sql:"pk: true, auto: true"`
    Login   string
    Email   string
}
```

This information allows the tool to generate smarter SQL statements:

```diff
CREATE TABLE IF NOT EXISTS users (
-user_id     INTEGER
+user_id     INTEGER PRIMARY KEY AUTOINCREMENT
,user_login  TEXT
,user_email  TEXT
);
```

Including SQL statements to select, insert, update and delete data using the primary key:

```Go
const SelectUserPkeyStmt = `
SELECT 
 id,
 login,
 email
WHERE user_id=?
`

const UpdateUserPkeyStmt = `
UPDATE users SET 
 id=?,
 login=?,
 email=?
WHERE user_id=?
`

const DeleteUserPkeyStmt = `
DELETE FROM users 
WHERE user_id=?
`
```

We can take this one step further and annotate indexes. In our example, we probably want to make sure the `user_login` field has a unique index:

```diff
type User struct {
    ID      int64  `sql:"pk: true, auto: true"`
-   Login   string
+   Login   string `sql:"unique: login"`
    Email   string
}
```

This information instructs the tool to generate the following:


```Go
const CreateUserLogin = `
CREATE UNIQUE INDEX IF NOT EXISTS user_login ON users (user_login)
```

#### Tags Summary

| Tag      | Value         | Purpose                                                      |
| -------- | ------------- | ------------------------------------------------------------ |
| pk       | true or false | the column is the primary key                                |
| auto     | true or false | the column is auto-incrementing (ignored if not using MySQL) |
| prefixed | true or false | the column name is made unique using a computed prefix       |
| name     | string        | the column name                                              |
| type     | string        | overrides the column type explicitly                         |
| size     | integer       | sets the storage size for the column                         |
| encode   | string        | encodes as "json", "text" or using the "driver"              |
| index    | string        | the column has an index                                      |
| unique   | string        | the column has a unique index                                |

Driver encoding means explicitly deferring to the `sql.Scanner` and `driver.Valuer` methods on your type.

See more details in [TAGS.md](TAGS.md).

### Nesting

Nested Go structures can be flattened into a single database table. As an example, we have a `User` and `Address` with a one-to-one relationship. In some cases, we may prefer to de-normalize our data and store in a single table, avoiding un-necessary joins.

```diff
type User struct {
    ID     int64  `sql:"pk: true"`
    Login  string
    Email  string
+   Addr   *Address
}

type Address struct {
    City   string
    State  string
    Zip    string `sql:"index: user_zip"`
}
```

The above relationship is flattened into a single table (see below). When the data is retrieved from the database the nested structure is restored.

```sql
CREATE TALBE IF NOT EXISTS users (
 id         INTEGER PRIMARY KEY AUTO_INCREMENT,
 login      TEXT,
 email      TEXT,
 addr_city  TEXT,
 addr_state TEXT,
 addr_zip   TEXT
);
```

### Shared Nested Structs With Tags

If you want to nest a struct into several types that you want to process with sqlgen, you might run up against a challenge: the field tags needed in one case might be inappropriate for another case. But there's an easy way around this issue: you can supply a Yaml file that sets the tags needed in each case. The Yaml file has tags that override the tags in the Go source code.

If you prefer, you can even use this approach for all tags; this means you don't need *any* tags in the Go source code.

This example has two fields with tags.

```go
type User struct {
    Uid uint64 `sql:"pk: true, auto: true"`
    Name       `sql:"name: username"`
}
```

The Yaml below overrides the tags. When used, the effect is that the `Uid` field would be neither primary key nor auto-incrementing. The `Name` field would have size 50 instead of the default (255), but the `name: username` setting would be kept.

```yaml
Uid:
  pk:   false
  auto: false
Name:
  size: 50
```

### JSON Encoding

Some types in your struct may not have native equivalents in your database such as `[]string`. These values can be marshalled and stored as JSON in the database.

```diff
type User struct {
    ID     int64  `sql:"pk: true"`
    Login  string
    Email  string
+   Label  []string `sql:"encode: json"
}
```

### Driver Encoding

If you have a field type that implements `sql.Scanner`, `driver.Valuer` and you want to use this explicitly as the column encoding to string data, specify `encode: driver`, e.g.

```
type User struct {
    ID     int64    `sql:"pk: true"`
    Stuff  MyStruct `sql:"encode: driver"`
}
```

You don't always have to do this because your types are inspected and will normally be auto-detected. However, for struct types that contain other fields, ambiguity arises. You can use this setting to resolve the ambiguity; otherwise the internal fields will be treated as table columns.


### Text Encoding

If you have a field type that implements `encoding.MashalText`, `encoding.UnmashalText` and you want to use this as the column encoding to string data, specify `encode: text`, e.g.

```
type User struct {
    ID     int64    `sql:"pk: true"`
    Thing  MyStruct `sql:"encode: text"`
}
```

There is no auto-detection for these enciding interfaces. Use this setting as and when you need it.


### Dialects

The generated code supports the following SQL dialects: `postgres`, `mysql` and `sqlite`. You decide at runtime which you need to use.


### Indexes

If your columns are indexes, `sqlgen` includes extra code for CRUD operations based on the indexed columns as well
as on the primary key. This example shows a primary key column `Id`, a uniquely-indexed column `Login`, and an
ordinary indexed column `Email`.

```Go
type User struct {
    Id     int64  `sql:"pk: true, auto: true"`
    Login  string `sql:"unique: user_login_idx"`
    Email  string `sql:"index: email_idx"`
    ...  other fields
}
```

## Joins and Views

Writing join queries is easy if you already know how to write the SQL. Sqlgen2 doesn't do this for you, but provides a `Query` method to do the heavy lifting.

All you need to write is a view type that is a Go struct to match the columns returned from the join. I like to name these something like `FirstSecondJoin`, referring to the first and second table types (and more if needed).

Because the only use-case is via the `Query` method, don't enable any of the command-line `-select`, `-update` etc flags. You won't need `schema` either.

You will probably find it helpful to use `-kind View`, which will cause your generated code to be `FirstSecondJoinView` (instead of `FirstSecondJoinTable`).  


## The API

The generated API contains a useful range of data-definition methods (create table, etc) and data-manipulation methods. The latter are all specialisations of Query and Exec. Query is the only method that is always included.

The API is summarised in [package sqlgen2](https://godoc.org/github.com/rickb777/sqlgen2) and [execer.go](https://github.com/rickb777/sqlgen2/blob/master/execer.go). See the interfaces `Table`, `TableCreator`, `TableWithIndexes`, and especially `TableWithCrud`.

But these interfaces are only part of the story. Where appropriate, the methods are type-safe. So, `Query` returns `[]User` (in the case of supporting the type `User`).

Because of the nature of Go's interfaces, the type-safe methods aren't expressly listed in the `TableWithCrud` interface. It's really useful to examine the generated code directly, in order to understand what's going on. There are some examples in the `demo` folder - see in particular [**user_sql.go**](https://github.com/rickb777/sqlgen2/blob/master/demo/user_sql.go), also [**shown here**](https://godoc.org/github.com/rickb777/sqlgen2/demo#DbUserTable).
 

### Declarative requirements

Exec, Select, Insert, Update and many other generated API methods accept a `require.Requirement` argument as the first parameter.

The requirement specifies the expected size of the result set, or the number of rows affected, as appropriate. If a requirement is not met, an error is returned with helpful diagnostics.

```Go
    err := tbl.Insert(require.Exactly(5), a, b c, d, e)
    ...
```

This can be easily ignored (just set it to nil). But it is often very useful, especially for Insert, Update and Delete methods.


### Where-expressions

Select, Count, Update and Delete methods accept where-expressions as parameters. Example:

```Go
    wh := where.Eq("name", "Andy").And(where.Gt("age", 18))
    value, err := tbl.Select(nil, wh, nil)
    ...
```

### Query Constraints

Select, Count, Update and Delete methods also accept query constraints as parameters. These specify result ordering and limit/offset parameters. Example:

```Go
    qc := where.OrderBy("name").Limit(10)
    value, err := tbl.Select(nil, nil, qc)
    ...
```


### Slice methods

This `SliceXxxx` set of methods returns vertical slices from the table, i.e. a set of values (strings, ints or floats) that come from one column of the table. They have a common format

```Go
SliceXxxx(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]T, error)
```

where *T* is the type of the corresponding field.


## Go Generate

Example use with `go:generate`:

```Go
package demo

//go:generate sqlgen -type User -pkg demo -o user_sql.go user.go

type User struct {
    ID     int64  `sql:"pk: true, auto: true"`
    Login  string `sql:"unique: user_login"`
    Email  string `sql:"size: 1024"`
    Avatar string
}
```

The current version is not smart enough to find the whole tree of source code containing dependent types. So you need
to list all the Go source files you want it to parse. This is an implementation limitation in the current version.


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
