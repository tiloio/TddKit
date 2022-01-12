package main

import (
	_ "embed"
	"log"
)


func RunAllTests(files *[]string) {
	var filesLength = len(*files)
	ch := make(chan FileResult, filesLength)

	for index, _ := range *files {
		var file = &(*files)[index]
		go RunTest(file, ch)
	}

	var result = FileResult{tests: 0, errors: 0}

	for i := 0; i < filesLength; i++ {
		currentResult := <-ch
		result.tests = result.tests + currentResult.tests
		result.errors = result.errors + currentResult.errors
	}

	log.Println("Tests:", result.tests, " - Failed:", result.errors)
}
