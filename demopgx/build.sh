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
#sqlgen -json -type demo.User -o user_ex_xxxxx_sql.go $V -prefix X -schema=false                                                              user.go role.go
#sqlgen -type demo.User -o user_ex_Cxxxx_sql.go $V -prefix C -schema=false -create=true  -read=false -update=false -delete=false -slice=false user.go role.go
#sqlgen -type demo.User -o user_ex_xRxxx_sql.go $V -prefix R -schema=false -create=false -read=true  -update=false -delete=false -slice=false user.go role.go
#sqlgen -type demo.User -o user_ex_xxUxx_sql.go $V -prefix U -schema=false -create=false -read=false -update=true  -delete=false -slice=false user.go role.go
#sqlgen -type demo.User -o user_ex_xxxDx_sql.go $V -prefix D -schema=false -create=false -read=false -update=false -delete=true  -slice=false user.go role.go
#sqlgen -type demo.User -o user_ex_xxxxS_sql.go $V -prefix S -schema=false -create=false -read=false -update=false -delete=false -slice=true  user.go role.go
#sqlgen -type demo.User -o user_ex_CRUDS_sql.go $V -prefix A -schema=false -all user.go role.go

unset GO_DRIVER GO_DSN PGQUOTE

go clean -testcache ||:

export DB_CONNECT_TIMEOUT=1s
export PGHOST=localhost
export PGDATABASE=test
export PGUSER PGPASSWORD
if [[ -z $PGUSER ]]; then
  PGUSER=test
  PGPASSWORD=test
fi

echo
echo "PGX (no quotes)...."
GO_DRIVER=pgx GO_DSN="postgres://$PGUSER:$PGPASSWORD@/test" PGQUOTE=none go test $V . ||:
echo
echo "PGX (ANSI)...."
GO_DRIVER=pgx GO_DSN="postgres://$PGUSER:$PGPASSWORD@/test" PGQUOTE=ansi go test $V . ||:
