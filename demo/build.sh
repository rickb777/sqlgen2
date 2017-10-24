#!/bin/bash -e
cd $(dirname $0)

sqlgen -v -type demo.Hook hook.go
sqlgen -v -type demo.Issue issue.go
sqlgen -v -type demo.User user.go

# there are no tests here (yet) but this ensures that the code is compilable
go test .
