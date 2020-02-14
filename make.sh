#!/bin/sh
rm -fr pc
GOOS=linux GOARCH=amd64 go build -o pc main.go
cp pc ../okay
