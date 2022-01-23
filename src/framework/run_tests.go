package main

import (
	_ "embed"
	"encoding/json"
	"log"
)

type TestResult struct {
	DiscoveryResult   DiscoveryResult `json:"discoveryResult"`
	Success           []TestResultLog `json:"success"`
	Errors            []TestResultLog `json:"errors"`
	DependencyErrored Dependency      `json:"dependencyErrored,omitempty"`
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
		Success:         make([]TestResultLog, 0),
		Errors:          make([]TestResultLog, 0),
		DiscoveryResult: discoveryResult,
	}

	var logMessages = make([]CommandLog, 0)
	for logMsg := range logs {
		logger <- logMsg
		if logMsg.framework {
			var currentResult = TestResultLog{}

			err := json.Unmarshal(logMsg.message, &currentResult)
			if err != nil {
				log.Fatalln("RunTest: Could not parse result:", err)
			}

			if currentResult.Err != "" {
				fileResult.Errors = append(fileResult.Errors, currentResult)
				log.Println("Test '", currentResult.Name, "' failed!", currentResult.Err)
			} else {
				fileResult.Success = append(fileResult.Success, currentResult)
			}

		}
		logMessages = append(logMessages, logMsg)
	}

	resultChannel <- fileResult
}
