package main

import (
	"flag"
	"strconv"
	"test-framework/internal/files"
	"test-framework/internal/logger"
)

func main() {
	flag.Parse()

	var fileList, err = files.Search()
	if err != nil {
		err.Fatal()
	}

	logger.Info("Found " + strconv.Itoa(len(*fileList)) + " files.")
}
