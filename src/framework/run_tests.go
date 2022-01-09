package main

import (
	_ "embed"
	"log"
	"path/filepath"
	"os"
	"os/exec"
	"bytes"
	"encoding/json"
)

//go:embed adapters/node/index.js
var nodeRequire []byte
var byteNewLine = []byte("\n")


type TestResult struct {
	Name string `json:"name"`
	Err string `json:"err,omitempty"`
}

type FileResult struct {
	tests int 
	errors int 
}

func executeTest(file *string) *exec.Cmd {
	fileContent, err := os.ReadFile(*file)
	if err != nil {
		log.Fatal("Could not read", file, "got err:", err)
	}

	var cmd *exec.Cmd
	if *esModule {
		cmd = exec.Command("node", "--input-type=module", "-")
	} else {
		cmd = exec.Command("node", "-")
	}

	buffer := bytes.Buffer{}
	buffer.Write(nodeRequire)
	buffer.Write(fileContent)
	cmd.Stdin = &buffer

	return cmd
}

func RunTest(file *string, resultChannel chan FileResult) {
	cmd := executeTest(file)
	cmd.Dir = filepath.Dir(*file)
	stdout, err := cmd.CombinedOutput()
	if err != nil {
		log.Println("Command finished with error:", err)
		log.Fatal(string(stdout))
	}


	var fileResult = FileResult{tests: 0, errors: 0}

	// todo rewrite to one loop
	var resultJsons = bytes.Split(stdout, byteNewLine)
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