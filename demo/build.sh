#!/bin/bash -e
cd $(dirname $0)

PATH=$PWD/..:$HOME/go/bin:$PATH

rm -f *_sql.go

go generate .

# also...

# These demonstrate the various filters that control what methods are generated
sqlgen -json -type demo.User -o user_ex_xxxxx_sql.go -v -prefix X -schema=false                                                                    user.go role.go
sqlgen -type demo.User -o user_ex_Cxxxx_sql.go -v -prefix C -schema=false -create=true  -read=false -update=false -delete=false -slice=false user.go role.go
sqlgen -type demo.User -o user_ex_xRxxx_sql.go -v -prefix R -schema=false -create=false -read=true  -update=false -delete=false -slice=false user.go role.go
sqlgen -type demo.User -o user_ex_xxUxx_sql.go -v -prefix U -schema=false -create=false -read=false -update=true  -delete=false -slice=false user.go role.go
sqlgen -type demo.User -o user_ex_xxxDx_sql.go -v -prefix D -schema=false -create=false -read=false -update=false -delete=true  -slice=false user.go role.go
sqlgen -type demo.User -o user_ex_xxxxS_sql.go -v -prefix S -schema=false -create=false -read=false -update=false -delete=false -slice=true  user.go role.go
sqlgen -type demo.User -o user_ex_CRUDS_sql.go -v -prefix A -schema=false -all user.go role.go

unset GO_DRIVER GO_DSN GO_QUOTER

for db in $@; do
  echo
  go clean -testcache ||:

  case $db in
    mysql)
      echo
      echo "MySQL...."
      GO_DRIVER=mysql GO_DSN=testuser:TestPasswd.9.9.9@/test go test -v .
      ;;

    postgres)
      echo
      echo "PostgreSQL (no quotes)...."
      GO_DRIVER=postgres GO_DSN="postgres://testuser:TestPasswd.9.9.9@/test" GO_QUOTER=none go test -v . ||:
      echo
      echo "PostgreSQL (ANSI)...."
      GO_DRIVER=postgres GO_DSN="postgres://testuser:TestPasswd.9.9.9@/test" GO_QUOTER=ansi go test -v . ||:
      ;;

    pgx)
      echo
      echo "PGX (no quotes)...."
      GO_DRIVER=pgx GO_DSN="postgres://testuser:TestPasswd.9.9.9@/test" GO_QUOTER=none go test -v . ||:
      echo
      echo "PGX (ANSI)...."
      GO_DRIVER=pgx GO_DSN="postgres://testuser:TestPasswd.9.9.9@/test" GO_QUOTER=ansi go test -v . ||:
      ;;

    sqlite) # default - see above
      unset GO_DRIVER GO_DSN
      echo
      echo "SQLite3 (no quotes)..."
      GO_QUOTER=none go test -v .
      echo
      echo "SQLite3 (ANSI)..."
      GO_QUOTER=ansi go test -v .
      ;;

    *)
      echo "$db: unrecognised; must be sqlite, mysql, or postgres"
      exit 1
      ;;
  esac
done
