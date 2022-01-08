#!/bin/bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
NAME=${1:-simple}
ESM=${2:-false}
EXEC_DIR="$DIR/src/test_examples/$NAME"
EXECUTABLE="$DIR/src/framework"

cd "$EXECUTABLE"
time go run . -path $EXEC_DIR -esm=$ESM
