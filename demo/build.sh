#!/bin/bash -e
cd $(dirname $0)

PATH=$PWD/..:$HOME/go/bin:$PATH

if [[ $1 = "-v" ]]; then
  V=-v
  shift
fi

rm -f *_sql.go *_sql.json

go generate .

# also...

# These demonstrate the various filters that control what methods are generated
sqlgen -type demo.User -o user_ex_xxxxxxxxx_sql.go $V -prefix X -json         user.go role.go
sqlgen -type demo.User -o user_ex_Exxxxxxxx_sql.go $V -prefix E -exec=true    user.go role.go
sqlgen -type demo.User -o user_ex_xQxxxxxxx_sql.go $V -prefix Q -query=true   user.go role.go
sqlgen -type demo.User -o user_ex_xxIxxxxxx_sql.go $V -prefix I -create=true  user.go role.go
sqlgen -type demo.User -o user_ex_xxxCxxxxx_sql.go $V -prefix C -count=true   user.go role.go
sqlgen -type demo.User -o user_ex_xxxxSxxxx_sql.go $V -prefix S -select=true  user.go role.go
sqlgen -type demo.User -o user_ex_xxxxxUxxx_sql.go $V -prefix U -update=true  user.go role.go
sqlgen -type demo.User -o user_ex_xxxxxxPxx_sql.go $V -prefix P -upsert=true  user.go role.go
sqlgen -type demo.User -o user_ex_xxxxxxxDx_sql.go $V -prefix D -delete=true  user.go role.go
sqlgen -type demo.User -o user_ex_xxxxxxxxL_sql.go $V -prefix L -slice=true   user.go role.go
sqlgen -type demo.User -o user_ex_EQICRUPDL_sql.go $V -prefix A -all          user.go role.go

# incomplete trial of sub-directory output
#mkdir -p sub
#sqlgen -type demo.Hook -pkg github.com/rickb777/sqlgen2/demo -o sub/xhook_sql.go -list demo.HookList -all $V .

unset GO_DRIVER GO_DSN GO_QUOTER

export DB_CONNECT_TIMEOUT=1s
export PGHOST=localhost
export PGDATABASE=test
export PGUSER PGPASSWORD
if [[ -z $PGUSER ]]; then
  PGUSER=testuser
  PGPASSWORD=TestPasswd.9.9.9
fi

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
      GO_DRIVER=mysql GO_DSN="$PGUSER:$PGPASSWORD@/test" go test $V .
      ;;

    postgres)
      echo
      echo "PostgreSQL (no quotes)...."
      GO_DRIVER=postgres GO_DSN="postgres://$PGUSER:$PGPASSWORD@/test" GO_QUOTER=none go test $V . ||:
      echo
      echo "PostgreSQL (ANSI)...."
      GO_DRIVER=postgres GO_DSN="postgres://$PGUSER:$PGPASSWORD@/test" GO_QUOTER=ansi go test $V . ||:
      ;;

    pgx)
      echo
      echo "PGX (no quotes)...."
      GO_DRIVER=pgx GO_DSN="postgres://$PGUSER:$PGPASSWORD@/test" GO_QUOTER=none go test $V . ||:
      echo
      echo "PGX (ANSI)...."
      GO_DRIVER=pgx GO_DSN="postgres://$PGUSER:$PGPASSWORD@/test" GO_QUOTER=ansi go test $V . ||:
      ;;

    sqlite)
      unset GO_DRIVER GO_DSN
      echo
      echo "SQLite3 (no quotes)..."
      GO_QUOTER=none go test $V .
      echo
      echo "SQLite3 (ANSI)..."
      GO_QUOTER=ansi go test $V .
      ;;

    *)
      echo "$db: unrecognised; must be sqlite, mysql, or postgres. Use 'all' for all of these."
      exit 1
      ;;
  esac
done
