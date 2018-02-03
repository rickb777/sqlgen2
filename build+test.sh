#!/bin/bash -e
cd $(dirname $0)

. go-get.sh

PATH=$HOME/gopath/bin:$GOPATH/bin:$PATH

# delete artefacts from previous build (if any)
rm -f *.out */*.txt demo/*_sql.go

### Dependencies ###

go_get bitbucket.org/pkg/inflect.a        bitbucket.org/pkg/inflect
go_get github.com/acsellers/inflections.a github.com/acsellers/inflections
go_get github.com/kortschak/utter.a       github.com/kortschak/utter
go_get gopkg.in/yaml.v2.a                 gopkg.in/yaml.v2

if ! type -p goveralls; then
  echo go get github.com/mattn/goveralls
  go get github.com/mattn/goveralls
fi

### Collection Types ###
# these generated files hardly ever need to change (see github.com/rickb777/runtemplate to do so)
[ -f schema/type_set.go ]   || runtemplate -tpl simple/set.tpl  -output schema/type_set.go  Type=Type   Comparable:true Ordered:false Numeric:false
[ -f util/int64_set.go ]    || runtemplate -tpl simple/set.tpl  -output util/int64_set.go   Type=int64  Comparable:true Ordered:true  Numeric:true
[ -f util/string_list.go ]  || runtemplate -tpl simple/list.tpl -output util/string_list.go Type=string Comparable:true Ordered:true  Numeric:false
[ -f util/string_set.go ]   || runtemplate -tpl simple/set.tpl  -output util/string_set.go  Type=string Comparable:true Ordered:true  Numeric:false

### Build Phase 1 ###

cd sqlgen
go install .

for d in code output parse; do
  echo sqlgen/$d...
  go test $1 -covermode=count -coverprofile=../$d.out ./$d
  go tool cover -func=../$d.out
  [ -z "$COVERALLS_TOKEN" ] || goveralls -coverprofile=$d.out -service=travis-ci -repotoken $COVERALLS_TOKEN
done

cd ..

### Build Phase 2 ###

go install ./...

for d in constraint require schema sqlgen where; do
  echo ./$d...
  go test $1 -covermode=count -coverprofile=./$d.out ./$d
  go tool cover -func=./$d.out
  [ -z "$COVERALLS_TOKEN" ] || goveralls -coverprofile=$d.out -service=travis-ci -repotoken $COVERALLS_TOKEN
done

#echo .
#go test . -covermode=count -coverprofile=dot.out .
#go tool cover -func=dot.out
#[ -z "$COVERALLS_TOKEN" ] || goveralls -coverprofile=$d.out -service=travis-ci -repotoken $COVERALLS_TOKEN

### Demo ###

cd demo
./build.sh sqlite mysql # postgres
