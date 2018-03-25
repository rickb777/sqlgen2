#!/bin/sh -e

VFILE=util/version.go

echo "// Updated automatically" > $VFILE
echo "package util" >> $VFILE
echo "" >> $VFILE
echo "const Version    = \"$(git describe --tags --dirty 2>/dev/null)\"" >> $VFILE
#echo "const CommitHash = \"$(git rev-parse --short HEAD 2>/dev/null)\"" >> $VFILE
echo "const GitBranch  = \"$(git symbolic-ref -q --short HEAD 2>/dev/null)\"" >> $VFILE
echo "const GitOrigin  = \"$(git remote get-url origin 2>/dev/null)\"" >> $VFILE
echo "const BuildDate  = \"$(date '+%FT%T')\"" >> $VFILE
