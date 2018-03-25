#!/bin/bash -e
# Speeds up build time by skipping unnecessary repeated scanning of Github sources.
# This has not been tested on Windows; it probably works but saves no time.

if [ -z "$GOPATH" ]; then
  echo "GOPATH must be set. Use only one path in it (i.e. no colons)."
  exit 1
fi

function getBinary
{
  if ! type -p $1 >/dev/null; then
    echo go get $2
    go get $2
  fi
}

getBinary dep         github.com/golang/dep/cmd/dep
getBinary runtemplate github.com/rickb777/runtemplate
getBinary goveralls   github.com/mattn/goveralls

dep ensure
