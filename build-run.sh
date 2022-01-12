#!/bin/bash


DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

DEFAULT_PATH="$DIR/src/test_examples"
DEFAULT_DIR='simple'
DEFAULT_GLOB='/**/*.test.[tj]s'

TEST_PATH="$DEFAULT_PATH/${1:-$DEFAULT_DIR}"
ESM=${2:-false}
GLOB=${3:-$DEFAULT_GLOB}
EXECUTABLE="$DIR/dist/runTests"

$DIR/build.sh

echo "Run $EXECUTABLE -path \"$TEST_PATH\" -glob \"$GLOB\" -esm=$ESM"


time $EXECUTABLE -path "$TEST_PATH" -glob "$GLOB" -esm=$ESM
