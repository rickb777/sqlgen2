#!/bin/bash -e
PATH=$HOME/gopath/bin:$GOPATH/bin:$PATH

if [ -z "$(type -p enumeration)" ]; then
  go get bitbucket.org/rickb777/enumeration
fi

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

for d in dialect schema where; do
  echo ./$d...
  go test $1 -covermode=count -coverprofile=./$d.out ./$d
  go tool cover -func=./$d.out
  [ -z "$COVERALLS_TOKEN" ] || goveralls -coverprofile=$d.out -service=travis-ci -repotoken $COVERALLS_TOKEN
done

cd demo
./build.sh
