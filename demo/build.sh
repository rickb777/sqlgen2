#!/bin/bash -e
cd $(dirname $0)

go generate .

# also...
sqlgen -type demo.User -o user_no_ddl_sql.go -v -prefix V2 -schema=false user.go
sqlgen -type demo.User -o user_no_crud_sql.go -v -prefix V3 -funcs=false user.go
sqlgen -type demo.User -o user_no_stuff_sql.go -v -prefix V4 -schema=false -funcs=false user.go

go test .
