#!/bin/bash

ls /usr/local/go/pkg/tool/linux_amd64/
ls /usr/local/go/pkg/linux_amd64/google.golang.org/
find / -name "grpc" 2>/dev/null

cd ../proto
protoc --go_out=plugins=grpc:. *.proto
cd -
go run main.go
