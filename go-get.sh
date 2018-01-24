#!/bin/bash -e
# Speeds up build time by skipping unnecessary repeated scanning of Github sources.
# This has not been tested on Windows; it probably works but saves no time.

if [ -z "$GOPATH" ]; then
  echo "GOPATH must be set. Use only one path in it (i.e. no colons)."
  exit 1
fi

export PKGDIR=$(go env GOOS)_$(go env GOARCH)

function go_get
{
  if [ ! -d $GOPATH/pkg/$PKGDIR/$1 -a ! -f $GOPATH/pkg/$PKGDIR/$1.a ]; then
    echo go get $2
    go get $2
  fi
}
