**sqlgen** generates SQL statements and database helper functions from your Go structs. It can be used in
place of a simple ORM or hand-written SQL.

See the [demo](https://github.com/rickb777/sqlgen2/tree/master/demo) directory for examples. Look in the
generated files `*_sql.go` and the hand-crafted files (`hook.go`, `issue.go`, `user.go`).

### Install

Install or upgrade with this command:

```
go get -u github.com/rickb777/sqlgen2
```

### Usage

```
sqlgen [option [arg]] file.go ...

Options:
  -type pkg.Type
    	primary type to analyse, which must be a struct type; required
  -o string
    	output file name; required
  -tags string
    	a YAML file containing tags that augment and override any attached to the fields in Go struct(s); optional
  -list string
    	names some type that is a collection of the primary type; optional; otherwise []*Type is used
  -prefix string
        prefix for names of generated types; optional, default is blank
  -kind string
    	suffix for names of generated types to indicate the intent; optional, default is "Table" but you might find nouns like "Join" or "View" are helpful
  -schema=true
    	generate sql schema and queries; default true
  -funcs=true
    	generate sql crud functions; default true
  -setters string
    	generate setter methods for fields in the primary type: none, optional (i.e. fields that are pointers), exported, all; default: none
  -gofmt
    	format and simplify the generated code nicely; default false
  -v
    	verbose progress; default false
  -z
    	debug info; default false
  -ast
    	debug AST info; default false
```

### Tutorial

`sqlgen` gives you a table-focussed view of your database. It generates code that maps each `struct`
you specify onto a database table (or view, or join result).

`sqlgen` is *not* an ORM: it does not automatically handle 'object' - relational mapping. It's simpler -
but just as comprehensive. 

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

The tool outputs the following generated code:

```Go
func ScanUser(row *sql.Row) (*User, error) {
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

	return v, nil
}

const CreateUserStmt = `
CREATE TABLE IF NOT EXISTS users (
 id     INTEGER,
 login  TEXT,
 email  TEXT
);
`

const SelectUserStmt = `
SELECT 
 id,
 login,
 email
FROM users 
`

const SelectUserRangeStmt = `
SELECT 
 id,
 login,
 email
FROM users 
LIMIT ? OFFSET ?
`


// more functions and sql statements not displayed
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

The tool also assumes that we probably intend to fetch data from the database using this index. The tool will therefore automatically generate the following queries:

```Go
const SelectUserLoginStmt = `
SELECT 
 id,
 login,
 email
WHERE user_login=?
`

const UpdateUserLoginStmt = `
UPDATE users SET 
 id=?,
 login=?,
 email=?
WHERE user_login=?
`

const DeleteUserLoginStmt = `
DELETE FROM users 
WHERE user_login=?
`
```

#### Tags Summary

| Tag      | Value         | Purpose                                                      |
| -------- | ------------- | ------------------------------------------------------------ |
| pk       | true or false | the column is the primary key                                |
| auto     | true or false | the column is auto-incrementing (ignored if not using MySQL) |
| prefixed | true or false | the column name is made unique using a computed prefixed     |
| name     | string        | the column name                                              |
| type     | string        | overrides the column type explicitly                         |
| size     | integer       | sets the storage size for the column                         |
| encode   | string        | encodes as "json" or "text"                                  |
| index    | string        | the column has an index                                      |
| unique   | string        | the column has a unique index                                |

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

Some types in your struct may not have native equivalents in your database such as `[]string`. These values can be marshaled and stored as JSON in the database.

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

### Where-expressions

Select, count amd update methods accept where-expressions as parameters. Example:

```Go
    ...  set up tbl
    wh := where.Eq("name", "Andy").And(where.Gt("age", 18))
    value, err := tbl.SelectOne(wh)
    ...
```


### Go Generate

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


### Benchmarks

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


### Restrictions

* Compound primaries are not supported
* In the structs used for tables, the imports must not use '.' or be renamed.


### Credits

This tool was derived from [sqlgen](https://github.com/drone/sqlgen) by drone.io, which was itself
inspired by [scaneo](https://github.com/variadico/scaneo).
