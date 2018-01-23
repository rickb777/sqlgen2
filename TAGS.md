# Tag In More Detail

Tags are optional annotations on the fields of the structs used to define your tables in Go.

```go
type User struct {
    ID      int64             `sql:"pk: true, auto: true"`
    Login   string            `sql:"unique: loginIdx"`
    Email   string            `sql:"index: emailIdx"`
    Age     uint              `sql:"check: age>=18"`
    Prefs   map[string]string `sql:"encode: json"`
}
```

## Tags Summary

| Tag      | Value         | Purpose                                                           |
| -------- | ------------- | ----------------------------------------------------------------- |
| pk       | true or false | the column is the primary key                                     |
| auto     | true or false | the column is auto-incrementing (ignored if not using MySQL)      |
| prefixed | true or false | the column name is made unique using a computed prefix            |
| name     | string        | the column name (default is the field's name)                     |
| type     | string        | overrides the column type explicitly                              |
| size     | integer       | sets the storage size for the column                              |
| encode   | string        | encodes as "json", "text" or using the "driver"                   |
| index    | string        | the column has an index                                           |
| unique   | string        | the column has a unique index                                     |
| check    | string        | apply the expression as a check constraint on writes to the table |
| fk       | string        | the column is a foreign key of some parent table                  |
| onupdate | string        | one of "cascade", "delete", "restrict", "set default", "set null" |
| ondelete | string        | one of "cascade", "delete", "restrict", "set default", "set null" |

Driver encoding means explicitly deferring to the `sql.Scanner` and `driver.Valuer` methods on your type.

## Encode

* "json" the struct field will be marshalled using Go's JSON marshaller and stored as a string of JSON.
* "text" the struct field will be marshalled using a TextMarshaler and stored as a string. You have to provide the TextMarshaller.
* "driver" the struct field will be marshalled by Go's SQL driver using interfaces implemented by you: `sql.Scanner` and `driver.Valuer`. The column type is whatever you want; your `Scanner` and `Valuer` functions must support the type you choose.

## Check Constraints

Use a tag such as `sql:"check: age>=18"`when you want the value of a column to be constrained by the expression you specify.

More info: https://www.postgresql.org/docs/9.6/static/ddl-constraints.html#DDL-CONSTRAINTS-CHECK-CONSTRAINTS

## Foreign Key Constraints

When `fk` is specified, it must contain a table name and the primary key in that table. It's in the form `table.pk`. Example: `sql:"fk: user.Id"`.
