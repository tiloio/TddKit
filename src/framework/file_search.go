package main

import (
	_ "embed"
	"log"
	"flag"
	"github.com/yargevad/filepathx"
)

var dirToSearch = flag.String("path", ".", "Path")
var globPattern = flag.String("glob", "/**/*.test.[tj]s", "Glob")


func SearchFiles() *[]string{
    path := *dirToSearch + *globPattern

	log.Println("Searching", path)

	files, err := filepathx.Glob(path)
	if err != nil {
		log.Fatal("Could not find", path, "got error:",err)
	}

	log.Println("Found", len(files), "files\n")
	return &files
} 