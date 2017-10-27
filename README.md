**sqlgen** generates SQL statements and database helper functions from your Go structs. It can be used in
place of a simple ORM or hand-written SQL.

See the [demo](https://github.com/rickb777/sqlgen2/tree/master/demo) directory for examples.

### Install

Install or upgrade with this command:

```
go get -u github.com/rickb777/sqlgen2
```

### Usage

```
sqlgen [option [arg]] file.go ...

Options:
  -type string
    	type to generate; required
  -prefix
        prefix for names of generated types; optional
  -file string
    	input file name; required
  -o string
    	output file name; required
  -schema
    	generate sql schema and queries; default true
  -funcs
    	generate sql crud functions; default true
  -gofmt
    	format and simplify the generated code nicely; default false
  -v
    	verbose progress; default false
  -z
    	debug info; default false
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
* Imports must not use '.' or be renamed.


### Credits

This tool was derived from [sqlgen](https://github.com/drone/sqlgen) by drone.io, which was itself
inspired by [scaneo](https://github.com/variadico/scaneo).
