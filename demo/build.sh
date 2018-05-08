#!/bin/bash -e
cd $(dirname $0)

PATH=$HOME/gopath/bin:$GOPATH/bin:$PATH

[ -n "$OFFLINE" ] || . ../go-get.sh

rm -f *_sql.go

go generate .

# also...

# These demonstrate the various filters that control what methods are generated
sqlgen -type demo.User -o user_ex_xxxxx_sql.go -v -prefix X -schema=false                                                                    user.go role.go
sqlgen -type demo.User -o user_ex_Cxxxx_sql.go -v -prefix C -schema=false -create=true  -read=false -update=false -delete=false -slice=false user.go role.go
sqlgen -type demo.User -o user_ex_xRxxx_sql.go -v -prefix R -schema=false -create=false -read=true  -update=false -delete=false -slice=false user.go role.go
sqlgen -type demo.User -o user_ex_xxUxx_sql.go -v -prefix U -schema=false -create=false -read=false -update=true  -delete=false -slice=false user.go role.go
sqlgen -type demo.User -o user_ex_xxxDx_sql.go -v -prefix D -schema=false -create=false -read=false -update=false -delete=true  -slice=false user.go role.go
sqlgen -type demo.User -o user_ex_xxxxS_sql.go -v -prefix S -schema=false -create=false -read=false -update=false -delete=false -slice=true  user.go role.go
sqlgen -type demo.User -o user_ex_CRUDS_sql.go -v -prefix A -schema=false -all user.go role.go

unset GO_DRIVER GO_DSN

echo
echo SQLite3...
echo go test .
go test .

for db in $@; do
  echo
  case $db in
    mysql)
      echo MySQL....
      echo go test .
      GO_DRIVER=mysql GO_DSN=testuser:TestPasswd9@/test go test .
      ;;

    postgres)
      echo PostgreSQL....
      echo go test .
      GO_DRIVER=postgres GO_DSN="postgres://testuser:TestPasswd9@/test" go test .
      ;;

    sqlite) # default - see above
      ;;

    *)
      echo "$db: unrecognised; must be sqlite, mysql, or postgres"
      exit 1
      ;;
  esac
done
