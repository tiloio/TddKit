package main

import (
	_ "embed"
	"flag"
)

func main() {
	flag.Parse()

	var files = SearchFiles()

	RunAllTests(files)
}
