package main

import (
	_ "embed"
	"flag"
)

func main() {
	flag.Parse()

	var logs = make(chan CommandLog)
	go readAndSaveLogs(logs)

	var files = SearchFiles()

	RunAllTests(files, logs)
}
