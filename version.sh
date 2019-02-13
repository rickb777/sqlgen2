#!/bin/sh -e
V=$1
if [ -z "$1" ]; then
  V=$(git describe --tags --always 2>/dev/null)
fi

VFILE=version.go

echo "// Updated automatically (altered manually just prior to each release)" > $VFILE
echo "" >> $VFILE
echo "package main" >> $VFILE
echo "" >> $VFILE
echo "const appVersion = \"$V\"" >> $VFILE
