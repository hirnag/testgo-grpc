#!/bin/bash

protoc --go_out=plugins=grpc:. pb/*.proto
go run main.go
