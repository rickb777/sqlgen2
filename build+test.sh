#!/bin/bash -e
PATH=$HOME/gopath/bin:$GOPATH/bin:$PATH

if ! type -p goveralls; then
  echo go get github.com/mattn/goveralls
  go get github.com/mattn/goveralls
fi

# TODO add 'schema' when it has some tests
for d in output parse; do
  echo $d...
  go test -v -covermode=count -coverprofile=$d.out ./$d
  go tool cover -func=$d.out
  [ -z "$COVERALLS_TOKEN" ] || goveralls -coverprofile=$d.out -service=travis-ci -repotoken $COVERALLS_TOKEN
done
