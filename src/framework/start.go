package main

import (
	_ "embed"
	"log"
	"flag"
	"github.com/yargevad/filepathx"
)

var dirToSearch = flag.String("path", ".", "Path")
var globPattern = flag.String("glob", "/**/*.test.[tj]s", "Glob")
var EsModule = flag.Bool("esm", false, "Is code es module?")

func main() {
    flag.Parse()
	path := *dirToSearch + *globPattern

	log.Println("Searching", path)

	files, err := filepathx.Glob(path)
	if err != nil {
		log.Fatal("Could not find", path, "got error:",err)
	}

	log.Println("Found", len(files), "files\n")

	ch := make(chan FileResult, len(files))

	for index, _ := range files {
		go RunTest(&files[index], ch)
	}

	var result = FileResult{tests: 0, errors: 0}

	for i := 0; i < len(files); i++ {
		currentResult := <-ch
		result.tests = result.tests + currentResult.tests
		result.errors = result.errors + currentResult.errors
	}

	log.Println("Tests:",  result.tests, " - Failed:", result.errors)
}