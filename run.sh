#!/bin/sh

set -e # Exit early if any commands fail

go build -o /tmp/simple-http app/*.go # Compile go code located in app directory into a binary located in the /tmp directory

exec /tmp/simple-http "$@" # Execute the newly created binary located in the /tmp directory
