# The API

Part of the code you'll be working with is auto-generated; the rest is the core API. We'll discuss the latter first.


## The core API

Unlike the generated API, the core API is not auto-generated. It's summarised in [package sqlgen2](https://godoc.org/github.com/rickb777/sqlgen2) and [execer.go](https://github.com/rickb777/sqlgen2/blob/master/execer.go). See the interfaces [Table](https://godoc.org/github.com/rickb777/sqlgen2#Table), [TableCreator](https://godoc.org/github.com/rickb777/sqlgen2#TableCreator), [TableWithIndexes](https://godoc.org/github.com/rickb777/sqlgen2#TableWithIndexes), and especially [TableWithCrud](https://godoc.org/github.com/rickb777/sqlgen2#TableWithCrud).

But these interfaces are only part of the story. Where appropriate, the methods are type-safe. So, `Query` returns `[]User` (in the case of supporting the type `User`).

Because of the nature of Go's interfaces, the type-safe methods aren't expressly listed in the `TableWithCrud` interface. It's really useful to examine the generated code directly, in order to understand what's going on. There are some examples in the `demo` folder - see in particular [**user_sql.go**](https://github.com/rickb777/sqlgen2/blob/master/demo/user_sql.go), also [**shown here**](https://godoc.org/github.com/rickb777/sqlgen2/demo#DbUserTable).
 

## Declarative requirements

Package [github.com/rickb777/sqlgen2/require](https://godoc.org/github.com/rickb777/sqlgen2/require) provides:

 * requirements that specify the expected size of the result set, or the number of rows affected, as appropriate.  

GetXxx, Exec, Select, Insert, Update and many other generated API methods accept a `require.Requirement` argument as the first parameter - more on these below.

The requirement specifies the expected size of the result set, or the number of rows affected, as appropriate. If a requirement is not met, an error is returned with helpful diagnostics.

```Go
    err := tbl.Insert(require.Exactly(5), a, b c, d, e)
    ...
```

The requirement can easily be ignored: just set it to nil. But it is often very useful, especially for Insert, Update and Delete methods.


## Where-expressions

Package [github.com/rickb777/sqlgen2/where](https://godoc.org/github.com/rickb777/sqlgen2/where) provides two related things:

 * a builder API for 'where' clauses
 * query constraints  

### Where-expression builder

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


## The Generated API

The generated API contains a useful range of data-definition methods (create table, etc) and data-manipulation methods that are specialised for the Go type and its corresponding database table. The data-manipulation methods are all specialisations of Query and Exec. Query is the only method that is always included.

A good example of the generated code is in [demo.AddressTable](https://godoc.org/github.com/rickb777/sqlgen2/demo#AddressTable).


### Get methods

This `GetXxxx` methods return one (or more) records of type *Thing* from the table where a particular field has a specified value. They have a common format

```Go
GetThingByField(req require.Requirement, value *type*) (*T, error)

GetThingsByField(req require.Requirement, value *type*) ([]T, error)
```

where *Thing* is the type of the corresponding field.

These are the easiest access methods to use and are provided for every indexed field (but currently only where the indexes are on a single column).


### Select and Count methods

These methods provide more flexible query options. The first three use parameters from the core API to make it easier to build requests:

```Go
Select(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]*T, error)

SelectOne(req require.Requirement, wh where.Expression, qc where.QueryConstraint) (*Address, error)

Count(wh where.Expression) (count int64, err error)
```

The second three are lower-level, don't need the core API builders, but a bit harder to use:

```Go
SelectWhere(req require.Requirement, where, orderBy string, args ...interface{}) ([]*T, error)

SelectOneWhere(req require.Requirement, where, orderBy string, args ...interface{}) (*T, error)

CountWhere(where string, args ...interface{}) (count int64, err error)
```

### Slice methods

This `SliceXxxx` methods return vertical slices from the table, i.e. a set of values (strings, ints or floats) that come from one column of the table. They have a common format

```Go
SliceXxxx(req require.Requirement, wh where.Expression, qc where.QueryConstraint) ([]T, error)
```

where *T* is the type of the corresponding field.


### Alteration methods

Insert takes one or more values of type *T:

```Go
Insert(req require.Requirement, vv ...*T) error
```

Update is similar; it writes the values of all the fields to the corresponding table records.

```Go
Update(req require.Requirement, vv ...*T) (int64, error)
```

There's a similar method for updating some but not all fields:

```Go
UpdateFields(req require.Requirement, wh where.Expression, fields ...sql.NamedArg) (int64, error)
```


### Delete methods

There are two kinds of delete methods: they are provided by primary key and by where-expression:

```Go
DeleteXxxx(req require.Requirement, id ...int64) (int64, error) // id type may differ

Delete(req require.Requirement, wh where.Expression) (int64, error)
```


## See Also

 * back to [**the tutorial**](tutorial.md)
