package main

import (
	"bytes"
	_ "embed"
	"log"
	"strings"
)

const LOG_PREFIX = "__TFW:"

var LOG_PREFIX_BYTES = len([]byte(LOG_PREFIX))

func ReadLogs(stdout *[]byte) [][]byte {
	var stdoutLines = bytes.Split(*stdout, byteNewLine)
	var results = [][]byte{}
	for _, stdoutLine := range stdoutLines {

		log.Println(string(stdoutLine))
		if len(stdoutLine) == 0 || !strings.HasPrefix(string(stdoutLine), LOG_PREFIX) {
			continue
		}
		var currentLine = stdoutLine[LOG_PREFIX_BYTES:]
		results = append(results, currentLine)
	}

	return results
}
