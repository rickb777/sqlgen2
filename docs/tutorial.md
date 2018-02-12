# Using sqlgen2

*sqlgen* is *not* an ORM: it does not automatically handle 'object' - relational mapping. It's simpler - but just as comprehensive.

You are expected to understand your own database; *sqlgen* removes the hard work of reliably writing the necessary `database/sql` API calls.

*sqlgen* gives you a table-focussed view of your database. The principal concept is a simple one.

 * You write a `struct` for every database table you want to manage.
 * If necessary, you also write a `struct` for every view or join result.

Then, *sqlgen* generates code that maps these onto the corresponding database tables.

Consider these two `structs`, both of which are in the [demo](https://github.com/rickb777/sqlgen2/tree/master/demo) package.

```Go
type User struct {
    Uid          int64
    Name         string
    EmailAddress string
    AddressId    *int64
    // ... other fields not showm
}

type Address struct {
    Id       int64
    Lines    []string
    Postcode string
}
```

These represent tables that have a one-to-many relationship. By the way, notice how `User.AddressId` is a pointer - as a convention, we use pointers for optional items.

![demo tables](demo-tables.png)

We can run the following commands:

```
sqlgen -type demo.User user.go
sqlgen -type demo.Address address.go
```

The tool generates code that includes `user_sql.go`:

```Go
type UserTable struct {
    // contains filtered or unexported fields
}

// NewUserTable returns a new table instance.
// If a blank table name is supplied, the default name "users" will be used instead.
// The request context is initialised with the background.
func NewUserTable(name sqlgen2.TableName, d *sqlgen2.Database) DbUserTable {
	// ... etc
}
```

and `address_sql.go`:

```Go
type AddressTable struct {
    // contains filtered or unexported fields
}

// NewAddressTable returns a new table instance.
// If a blank table name is supplied, the default name "users" will be used instead.
// The request context is initialised with the background.
func NewAddressTable(name sqlgen2.TableName, d *sqlgen2.Database) DbUserTable {
	// ... etc
}
```

The table `structs` have many methods to access the table. You have some control over what is provided and this will be described further later on. For example, you might not need update methods on a log table.

The two generated `structs` are related to a provided type called `Database`.

![database-and-tables](database-and-tables.png)

You normally have exactly one [`*Database`](https://godoc.org/github.com/rickb777/sqlgen2#Database) in your app: it wraps the `*sql.DB` connection and logger (if you need one).

## Controlling What the Columns Mean

This is a great start, but what if we want to specify primary keys, column sizes and more? This may be achieved by annotating your code using Go tags. For example, we can tag the `ID` field to indicate it is a primary key and will auto increment:

```diff
type User struct {
-   Uid     int64
+   Uid     int64  `sql:"pk: true, auto: true"`
    Name    string
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
    Uid     int64  `sql:"pk: true, auto: true"`
-   Name    string
+   Name    string `sql:"unique: login"`
    Email   string
}
```

This information instructs the tool to generate the following:


```Go
const CreateUserName  = `
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

See more details in [tags.md](tags.md).

### Nesting

Nested Go structures can be flattened into a single database table. As an example, we have a `User` and `Address` with a one-to-one relationship. In some cases, we may prefer to de-normalize our data and store in a single table, avoiding un-necessary joins.

```diff
type User struct {
    Uid    int64  `sql:"pk: true"`
    Name   string
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
    Uid    int64  `sql:"pk: true"`
    Name   string
    Email  string
+   Label  []string `sql:"encode: json"
}
```

### Driver Encoding

If you have a field type that implements `sql.Scanner`, `driver.Valuer` and you want to use this explicitly as the column encoding to string data, specify `encode: driver`, e.g.

```
type User struct {
    Uid    int64    `sql:"pk: true"`
    Stuff  MyStruct `sql:"encode: driver"`
}
```

You don't always have to do this because your types are inspected and will normally be auto-detected. However, for struct types that contain other fields, ambiguity arises. You can use this setting to resolve the ambiguity; otherwise the internal fields will be treated as table columns.


### Text Encoding

If you have a field type that implements `encoding.MashalText`, `encoding.UnmashalText` and you want to use this as the column encoding to string data, specify `encode: text`, e.g.

```
type User struct {
    Uid    int64    `sql:"pk: true"`
    Thing  MyStruct `sql:"encode: text"`
}
```

There is no auto-detection for these enciding interfaces. Use this setting as and when you need it.


### Dialects

The generated code supports the following SQL dialects: `postgres`, `mysql` and `sqlite`. You decide at runtime which you need to use.


### Indexes

If your columns need indexes, `sqlgen` includes extra code for CRUD operations based on the indexed columns as well as on the primary key. This example shows a primary key column `Id`, an indexed column `Name `, and a uniquely indexed column `EmailAddress`.

```Go
type User struct {
	Uid          int64    `sql:"pk: true, auto: true"`
	Name         string   `sql:"index: user_name"`
	EmailAddress string   `sql:"unique: user_email"`
	AddressId    *int64   `sql:"fk: addresses.id, onupdate: restrict, ondelete: restrict"`
    // ... other fields not showm
}
```

[List of all the tags](tags.md).


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

//go:generate sqlgen -type demo.User -gofmt -all -setters=all user.go role.go

type User struct {
    Uid    int64  `sql:"pk: true, auto: true"`
    Name   string `sql:"unique: user_login"`
    Email  string `sql:"size: 1024"`
    Avatar string
}
```

See the [command line options](usage.md).


## Restrictions

* Compound primaries are not supported
* In the structs used for tables, the imports must not use '.' or be renamed.

