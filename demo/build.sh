#!/bin/bash -e
cd $(dirname $0)
rm -f *_sql.go

go generate .

# also...

# This has no DDL methods; the type will be called V2UserJoin
sqlgen -type demo.User -o user_no_ddl_sql.go   -v -prefix V2 -schema=false              -kind Join user.go

# This has no crud methods; the type will be called UserTbl
sqlgen -type demo.User -o user_no_crud_sql.go  -v            -funcs=false               -kind Tbl  user.go

# This has neither DDL nor crud methods; the type will be called V4UserTable
sqlgen -type demo.User -o user_no_stuff_sql.go -v -prefix V4 -schema=false -funcs=false            user.go

go get github.com/mattn/go-sqlite3

go test .
