#!/bin/bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
EXECUTABLE="$DIR/dist/runTests"

cd "$DIR/src/framework"
go build -o $EXECUTABLE