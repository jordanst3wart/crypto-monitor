#!/usr/bin/env sh

# test
go test ./...
# compile
env GOOS=linux go build -ldflags="-s -w" -o bin/main main/functions.go main/main.go
