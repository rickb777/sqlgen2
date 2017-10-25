#!/bin/bash -e
cd $(dirname $0)

go generate .
# there are no tests here (yet) but this ensures that the code is compilable
go test .
