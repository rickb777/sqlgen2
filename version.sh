#!/bin/sh -e

VFILE=version.go

echo "// Updated automatically (altered manually just prior to each release)" > $VFILE
echo "" >> $VFILE
echo "package main" >> $VFILE
echo "" >> $VFILE
echo "const appVersion = \"$(git describe --tags --always 2>/dev/null)\"" >> $VFILE
