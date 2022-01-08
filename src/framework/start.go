package main

import (
	_ "embed"
	"log"
	"io/fs"
	"flag"
	"path/filepath"
	"strings"
)


var dirToSearch = flag.String("path", ".", "Path")
var esModule = flag.Bool("esm", false, "Is code es module?")

func main() {
    flag.Parse()
	var path, err = filepath.Abs(*dirToSearch)
	if err != nil {
		log.Fatal(err)
	}

	var files []string


	log.Println("Search dir", path)
	 err = filepath.WalkDir(*dirToSearch, func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			log.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}

		if strings.HasSuffix(path, ".test.js") {
			files = append(files, path)
		}

		return nil
	})
	if err != nil {
		log.Printf("error walking the path %q: %v\n", dirToSearch, err)
		return
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