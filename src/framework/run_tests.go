package main

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"log"
	"path/filepath"
)

type TestResult struct {
	Name string `json:"name"`
	Err  string `json:"err,omitempty"`
}

type FileResult struct {
	tests  int
	errors int
}

var byteNewLine = []byte("\n")

func RunTest(file *string, resultChannel chan FileResult) {
	absolutePath, err := filepath.Abs(*file)
	if err != nil {
		log.Fatal("Could not read file:", err)
	}

	stdout := ExecuteEcmascriptTests(&absolutePath)

	var fileResult = FileResult{tests: 0, errors: 0}

	// todo rewrite to one loop
	var resultJsons = bytes.Split(*stdout, byteNewLine)
	var results = make([]TestResult, cap(resultJsons))
	for index, resultJson := range resultJsons {

		if len(resultJson) == 0 {
			continue
		}
		var currentResult = results[index]
		err := json.Unmarshal(resultJson, &currentResult)
		if err != nil {
			log.Println("Could not parse result:", err)
		}
		fileResult.tests = fileResult.tests + 1

		if currentResult.Err != "" {
			fileResult.errors = fileResult.errors + 1
			log.Println("Test '", currentResult.Name, "' failed!", currentResult.Err)
		}
	}

	resultChannel <- fileResult
}
