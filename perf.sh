#!/bin/bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
TEST_EXAMPLES_DIR="$DIR/src/test_examples"

#$TEST_EXAMPLES_DIR/jest.sh $1 $2
echo ""
$TEST_EXAMPLES_DIR/framework.sh $1 $2
