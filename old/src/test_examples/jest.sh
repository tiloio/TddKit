#!/bin/bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
TARGET_DIR="$DIR/${1:-simple}"
TIMES=${2:-10}


if ! command -v jest &> /dev/null
then
    echo 'Error: need jest to work, install via'
    echo '  npm i -g jest'
    exit 1
fi

echo "Running JEST $TIMES times on $TARGET_DIR"

cd $TARGET_DIR
jest --clearCache

perfRun() {
    for (( i = 0; i < $TIMES; i++ )); do
        jest > /dev/null 2>&1
    done
}

time perfRun
