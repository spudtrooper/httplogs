#!/bin/sh

set -e

rm -f main
go build main.go
rm -f ~/go/bin/httplogs
ln -fns /Users/jeff/Projects/httplogs/main ~/go/bin/httplogs

~/go/bin/httplogs --help

