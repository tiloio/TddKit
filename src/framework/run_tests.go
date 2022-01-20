package main

import (
	_ "embed"
	"encoding/json"
	"log"
)

type TestResult struct {
	discoveryResult DiscoveryResult
	tests           int
	errors          int
	dependencyError bool
}

type TestLog struct {
	Name string `json:"name"`
	Err  string `json:"err,omitempty"`
}

var byteNewLine = []byte("\n")
var environmentVariables = []string{TEST_PHASE_ENVIRONMENT_VAIRABLE}

func RunTest(discoveryResult DiscoveryResult, resultChannel chan TestResult) {
	log.Println("Running test of", discoveryResult.File.Name, discoveryResult.Dependency.Id)

	var environment = append(environmentVariables, TEST_ID_ENVIRONMENT_VARIABLE+"="+discoveryResult.Dependency.Id)
	stdout := ExecuteEcmascriptTests(&discoveryResult.File.content, &environment)

	var fileResult = TestResult{
		tests:           0,
		errors:          0,
		discoveryResult: discoveryResult,
	}

	var logs = ReadLogs(stdout)

	var results = make([]TestLog, len(logs))
	for index, logMsg := range logs {
		var currentResult = results[index]
		err := json.Unmarshal(logMsg, &currentResult)
		if err != nil {
			log.Fatalln("RunTest: Could not parse result:", err)
		}
		fileResult.tests = fileResult.tests + 1

		if currentResult.Err != "" {
			fileResult.errors = fileResult.errors + 1
			log.Println("Test '", currentResult.Name, "' failed!", currentResult.Err)
		}
	}

	resultChannel <- fileResult
}
