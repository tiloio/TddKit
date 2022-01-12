#!/bin/bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
TARGET_DIR="$DIR/${1:-simple}"
EXECUTABLE="$DIR/../../dist/runTests"
ROOT_DIR="$DIR/../../"
TIMES=${2:-10}

$ROOT_DIR/build.sh

echo "Running FRAMEWORK $TIMES times on $TARGET_DIR"

perfRun() {
    for (( i = 0; i < $TIMES; i++ )); do
        $EXECUTABLE -path $TARGET_DIR -jest  > /dev/null 2>&1
    done
}

time perfRun
