# sqlgen Command Line Usage

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
