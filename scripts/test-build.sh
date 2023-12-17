#!/usr/bin/env sh

# test
/Users/jordanstewart/go/go1.21.5/bin/go test ./...
# compile
env GOOS=linux /Users/jordanstewart/go/go1.21.5/bin/go build -ldflags="-s -w" -o bin/main main/functions.go main/main.go
