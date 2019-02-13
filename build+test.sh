#!/bin/bash -e
cd $(dirname $0)

function announce
{
  echo
  echo $@
}

PATH=$HOME/go/bin:$PATH
unset GOPATH

go mod download

# delete artefacts from previous build (if any)
mkdir -p reports
rm -f reports/*.out reports/*.html */*.txt demo/*_sql.go

### Build Phase 1 ###

./version.sh $1
go test . ./code ./load ./output ./parse
go build -o sqlgen *.go

[ -d $HOME/go/bin ] && cp -v sqlgen $HOME/go/bin

for d in code output parse; do
  announce sqlgen/$d
  go test -covermode=count -coverprofile=reports/sqlgen-$d.out ./$d
  go tool cover -html=reports/sqlgen-$d.out -o reports/sqlgen-$d.html
  [ -z "$COVERALLS_TOKEN" ] || goveralls -coverprofile=reports/sqlgen-$d.out -service=travis-ci -repotoken $COVERALLS_TOKEN
done

### Build Phase 2 ###

#sqlgen -type vanilla.Record -o vanilla/vanilla_sql.go -read -delete -slice -v vanilla/vanilla.go

### Demo ###

announce demo
cd demo
./build.sh sqlite mysql postgres pgx

cd ..
echo
echo go vet -shadow=true -composites=false ./...
go vet -shadow=true -composites=false ./...
