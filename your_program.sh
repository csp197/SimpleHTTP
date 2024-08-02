#!/bin/sh

set -e # Exit early if any commands fail

# - Edit this to change how your program compiles locally
go build -o /tmp/simple-http app/*.go

# - Edit this to change how your program runs locally
exec /tmp/simple-http "$@"
