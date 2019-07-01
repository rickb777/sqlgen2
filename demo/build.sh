#!/bin/bash -e
cd $(dirname $0)

PATH=$PWD/..:$HOME/go/bin:$PATH

rm -f *_sql.go *_sql.json

go generate .

# also...

# These demonstrate the various filters that control what methods are generated
sqlgen -json -type demo.User -o user_ex_xxxxxx_sql.go -v -prefix X -schema=false                                                                    user.go role.go
sqlgen -type demo.User -o user_ex_Ixxxxx_sql.go -v -prefix I -schema=false -create=true  -count=false -read=false -update=false -delete=false -slice=false user.go role.go
sqlgen -type demo.User -o user_ex_xCxxxx_sql.go -v -prefix C -schema=false -create=false -count=true  -read=false -update=false -delete=false -slice=false user.go role.go
sqlgen -type demo.User -o user_ex_xxRxxx_sql.go -v -prefix R -schema=false -create=false -count=false -read=true  -update=false -delete=false -slice=false user.go role.go
sqlgen -type demo.User -o user_ex_xxxUxx_sql.go -v -prefix U -schema=false -create=false -count=false -read=false -update=true  -delete=false -slice=false user.go role.go
sqlgen -type demo.User -o user_ex_xxxxDx_sql.go -v -prefix D -schema=false -create=false -count=false -read=false -update=false -delete=true  -slice=false user.go role.go
sqlgen -type demo.User -o user_ex_xxxxxS_sql.go -v -prefix S -schema=false -create=false -count=false -read=false -update=false -delete=false -slice=true  user.go role.go
sqlgen -type demo.User -o user_ex_ICRUDS_sql.go -v -prefix A -schema=false -all user.go role.go

unset GO_DRIVER GO_DSN GO_QUOTER

DBS=$@
if [ "$1" = "all" ]; then
  DBS="sqlite mysql postgres pgx"
fi

for db in $DBS; do
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

    sqlite)
      unset GO_DRIVER GO_DSN
      echo
      echo "SQLite3 (no quotes)..."
      GO_QUOTER=none go test -v .
      echo
      echo "SQLite3 (ANSI)..."
      GO_QUOTER=ansi go test -v .
      ;;

    *)
      echo "$db: unrecognised; must be sqlite, mysql, or postgres. Use 'all' for all of these."
      exit 1
      ;;
  esac
done
