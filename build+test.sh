#!/bin/bash -e
cd $(dirname $0)
PATH=$HOME/gopath/bin:$GOPATH/bin:$PATH

# delete artefacts from previous build (if any)
rm -f *.out */*.txt demo/*_sql.go

go get bitbucket.org/pkg/inflect
go get github.com/acsellers/inflections
go get github.com/kortschak/utter
go get gopkg.in/yaml.v2

# these generated files hardly ever need to change (see github.com/rickb777/runtemplate to do so)
[ -f schema/type_set.go ]        || runtemplate -tpl simple/set.tpl -output schema/type_set.go        Type=Type
[ -f sqlgen/code/string_set.go ] || runtemplate -tpl simple/set.tpl -output sqlgen/code/string_set.go Type=string

cd sqlgen
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
go install .

for d in schema sqlgen where; do
  echo ./$d...
  go test $1 -covermode=count -coverprofile=./$d.out ./$d
  go tool cover -func=./$d.out
  [ -z "$COVERALLS_TOKEN" ] || goveralls -coverprofile=$d.out -service=travis-ci -repotoken $COVERALLS_TOKEN
done

#echo .
#go test . -covermode=count -coverprofile=dot.out .
#go tool cover -func=dot.out
#[ -z "$COVERALLS_TOKEN" ] || goveralls -coverprofile=$d.out -service=travis-ci -repotoken $COVERALLS_TOKEN

cd demo
./build.sh
