#!/bin/bash -e
cd $(dirname $0)
rm -f *_sql.go

go generate .

# also...

# This has only Exec and Query, no other DDL or DML methods; the type will be called X1UserJoin
sqlgen -type demo.User -o user_ex_exec_query_sql.go -v -prefix X1 -kind Join -schema=false -funcs=none user.go

# These demonstrate the various filters that control what methods are generated
sqlgen -type demo.User -o user_ex_xxxxx_sql.go -v -prefix X -schema=false -funcs=none                                                        user.go
sqlgen -type demo.User -o user_ex_Cxxxx_sql.go -v -prefix C -schema=false -create=true  -read=false -update=false -delete=false -slice=false user.go
sqlgen -type demo.User -o user_ex_xRxxx_sql.go -v -prefix R -schema=false -create=false -read=true  -update=false -delete=false -slice=false user.go
sqlgen -type demo.User -o user_ex_xxUxx_sql.go -v -prefix U -schema=false -create=false -read=false -update=true  -delete=false -slice=false user.go
sqlgen -type demo.User -o user_ex_xxxDx_sql.go -v -prefix D -schema=false -create=false -read=false -update=false -delete=true  -slice=false user.go
sqlgen -type demo.User -o user_ex_xxxxS_sql.go -v -prefix S -schema=false -create=false -read=false -update=false -delete=false -slice=true  user.go
sqlgen -type demo.User -o user_ex_CRUDS_sql.go -v -prefix A -schema=false -funcs=all user.go

go get github.com/mattn/go-sqlite3

go test .
