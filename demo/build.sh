#!/bin/bash -e
cd $(dirname $0)

go generate .

# also...
sqlgen -type demo.User -o user2_sql.go -v -prefix V2 -schema=false user.go
#sqlgen -type demo.User -o user3_sql.go -v -prefix V3 -extras=false user.go
#sqlgen -type demo.User -o user4_sql.go -v -prefix V4 -funcs=false user.go

go test .
