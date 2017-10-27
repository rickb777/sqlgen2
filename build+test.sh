#!/bin/bash -e
PATH=$HOME/gopath/bin:$GOPATH/bin:$PATH

if [ -z "$(type -p enumeration)" ]; then
  go get bitbucket.org/rickb777/enumeration
fi

go get bitbucket.org/pkg/inflect
go get github.com/acsellers/inflections
go get github.com/kortschak/utter
go get gopkg.in/yaml.v2

enumeration -i schema/dialect.go -o schema/dialect_enum.go -package schema -type DialectId

cd sqlgen
enumeration -i parse/kind.go -o parse/kind_enum.go -package parse -type Kind

go install .

if ! type -p goveralls; then
  echo go get github.com/mattn/goveralls
  go get github.com/mattn/goveralls
fi

for d in code output parse; do
  echo sqlgen/$d...
  go test $1 -covermode=count -coverprofile=../$d.out ./$d
  go tool cover -func=../$d.out
  [ -z "$COVERALLS_TOKEN" ] || goveralls -coverprofile=$d.out -service=travis-ci -repotoken $COVERALLS_TOKEN
done

cd ..
#go install .

for d in schema where; do
  echo ./$d...
  go test $1 -covermode=count -coverprofile=./$d.out ./$d
  go tool cover -func=./$d.out
  [ -z "$COVERALLS_TOKEN" ] || goveralls -coverprofile=$d.out -service=travis-ci -repotoken $COVERALLS_TOKEN
done

echo .
go test . -covermode=count -coverprofile=dot.out .
go tool cover -func=dot.out
[ -z "$COVERALLS_TOKEN" ] || goveralls -coverprofile=$d.out -service=travis-ci -repotoken $COVERALLS_TOKEN

cd demo
./build.sh
