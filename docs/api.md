# The API

The generated API contains a useful range of data-definition methods (create table, etc) and data-manipulation methods. The latter are all specialisations of Query and Exec. Query is the only method that is always included.

The API is summarised in [package sqlgen2](https://godoc.org/github.com/rickb777/sqlgen2) and [execer.go](https://github.com/rickb777/sqlgen2/blob/master/execer.go). See the interfaces `Table`, `TableCreator`, `TableWithIndexes`, and especially `TableWithCrud`.

But these interfaces are only part of the story. Where appropriate, the methods are type-safe. So, `Query` returns `[]User` (in the case of supporting the type `User`).

Because of the nature of Go's interfaces, the type-safe methods aren't expressly listed in the `TableWithCrud` interface. It's really useful to examine the generated code directly, in order to understand what's going on. There are some examples in the `demo` folder - see in particular [**user_sql.go**](https://github.com/rickb777/sqlgen2/blob/master/demo/user_sql.go), also [**shown here**](https://godoc.org/github.com/rickb777/sqlgen2/demo#DbUserTable).
 

## Declarative requirements

Package `github.com/rickb777/sqlgen2/require` provides:

 * requirements that specify the expected size of the result set, or the number of rows affected, as appropriate.  

Exec, Select, Insert, Update and many other generated API methods accept a `require.Requirement` argument as the first parameter.

The requirement specifies the expected size of the result set, or the number of rows affected, as appropriate. If a requirement is not met, an error is returned with helpful diagnostics.

```Go
    err := tbl.Insert(require.Exactly(5), a, b c, d, e)
    ...
```

This can be easily ignored (just set it to nil). But it is often very useful, especially for Insert, Update and Delete methods.


## Where-expressions

Package `github.com/rickb777/sqlgen2/where` provides:

 * a builder API for 'where' clauses
 * query constraints  

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


## See Also

 * back to [**the tutorial**](tutorial.md)
