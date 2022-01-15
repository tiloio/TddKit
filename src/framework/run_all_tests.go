package main

import (
	_ "embed"
	"flag"
	"log"
)

var esModule = flag.Bool("esm", false, "Is code es module?")

func RunAllTests(files *[]string) {
	var filesLength = len(*files)
	parsedFileChannel := make(chan ParsedFile, filesLength)

	for _, file := range *files {
		go ParseFileAsync(file, esModule, parsedFileChannel)
	}

	var parsedFiles = make([]ParsedFile, filesLength)

	for i := 0; i < filesLength; i++ {
		parsedFiles[i] = <-parsedFileChannel
	}

	discoveryResultCh := make(chan DiscoveryResult, filesLength)
	for _, file := range parsedFiles {
		go RunDiscoveryPhase(file, discoveryResultCh)
	}

	discoveryResults := make([]DiscoveryResult, filesLength)
	for index, _ := range discoveryResults {
		discoveryResults[index] = <-discoveryResultCh
	}

	testResultCh := make(chan TestResult, filesLength)
	for _, discoveryResult := range discoveryResults {
		go RunTest(discoveryResult, testResultCh)
	}

	var result = TestResult{tests: 0, errors: 0}
	for i := 0; i < filesLength; i++ {
		currentResult := <-testResultCh
		result.tests = result.tests + currentResult.tests
		result.errors = result.errors + currentResult.errors
	}

	log.Println("Tests:", result.tests, " - Failed:", result.errors)
}
