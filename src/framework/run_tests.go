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

type TestResultLog struct {
	Name string `json:"name"`
	Err  string `json:"err,omitempty"`
}

var byteNewLine = []byte("\n")
var environmentVariables = []string{TEST_PHASE_ENVIRONMENT_VAIRABLE}

func RunTest(discoveryResult DiscoveryResult, resultChannel chan TestResult, logger chan CommandLog) {
	log.Println("Running test of", discoveryResult.File.Name, discoveryResult.Dependency.Id)

	var environment = append(environmentVariables, TEST_ID_ENVIRONMENT_VARIABLE+"="+discoveryResult.Dependency.Id)

	var logs = make(chan CommandLog)
	go ExecuteEcmascriptTests(&discoveryResult.File.content, &environment, logs)

	var fileResult = TestResult{
		tests:           0,
		errors:          0,
		discoveryResult: discoveryResult,
	}

	var results = make([]TestResultLog, 0)
	var logMessages = make([]CommandLog, 0)
	for logMsg := range logs {
		logger <- logMsg
		if logMsg.framework {
			var currentResult = TestResultLog{}

			err := json.Unmarshal(logMsg.message, &currentResult)
			if err != nil {
				log.Fatalln("RunTest: Could not parse result:", err)
			}
			fileResult.tests = fileResult.tests + 1

			if currentResult.Err != "" {
				fileResult.errors = fileResult.errors + 1
				log.Println("Test '", currentResult.Name, "' failed!", currentResult.Err)
			}

			results = append(results, currentResult)
		}
		logMessages = append(logMessages, logMsg)
	}

	resultChannel <- fileResult
}
