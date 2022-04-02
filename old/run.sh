#!/bin/bash

DEFAULT_PATH='../test_examples'
DEFAULT_DIR='simple'
DEFAULT_GLOB='/**/*.test.[tj]s'

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
TEST_PATH="$DEFAULT_PATH/${1:-$DEFAULT_DIR}"
ESM=${2:-false}
GLOB=${3:-$DEFAULT_GLOB}
EXECUTABLE="$DIR/src/framework"

echo "Run $EXECUTABLE go run . -path \"$TEST_PATH\" -glob \"$GLOB\" -esm=$ESM"

cd "$EXECUTABLE"
time go run . -path "$TEST_PATH" -glob "$GLOB" -esm=$ESM
