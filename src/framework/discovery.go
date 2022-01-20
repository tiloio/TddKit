package main

import (
	_ "embed"
	"encoding/json"
	"log"
)

type DiscoveryResult struct {
	File       ParsedFile `json:"file"`
	Tests      []Test     `json:"tests"`
	Dependency Dependency `json:"dependency"`
	Resources  []Resource `json:"resources"`
}

type Dependency struct {
	Id           string       `json:"id"`
	Dependencies []Dependency `json:"dependencies"`
}

type Resources struct {
	Resources []Resource `json:"resources"`
}

type Resource struct {
	Id        string     `json:"id"`
	Resources []Resource `json:"resources"`
}
type Test struct {
	Name string `json:"name"`
}

type DiscoveryLog struct {
	Type string `json:"type"`
}

const DEPENDENCY_TYPE = "DEPENDENCY"
const RESOURCE_TYPE = "RESOURCE"
const TEST_TYPE = "TEST"

var newLineAsByte = []byte("\n")
var discoveryEnvironmentVariables = []string{DISCOVERY_PHASE_ENVIRONMENT_VAIRABLE}

func RunDiscoveryPhase(file ParsedFile, resultCh chan DiscoveryResult) {
	stdout := ExecuteEcmascriptTests(&file.content, &discoveryEnvironmentVariables)

	var logs = ReadLogs(stdout)

	var result = DiscoveryResult{
		File:  file,
		Tests: make([]Test, 0),
	}

	var lastDependencLogIndex = 0

	for index, logMsg := range logs {

		var dicoveryLog = DiscoveryLog{}
		if err := json.Unmarshal(logMsg, &dicoveryLog); err != nil {
			log.Fatalln("RunDiscovery: Could not parse discovery log:", err)
		}

		switch dicoveryLog.Type {
		case DEPENDENCY_TYPE:
			if err := json.Unmarshal(logMsg, &result.Dependency); err != nil {
				log.Fatalln("RunDiscovery: Could not parse dependency:", err)
			}
			lastDependencLogIndex = index
		}
	}

	for i := lastDependencLogIndex; i < len(logs); i++ {
		var logMsg = logs[i]

		var dicoveryLog = DiscoveryLog{}
		if err := json.Unmarshal(logMsg, &dicoveryLog); err != nil {
			log.Fatalln("RunDiscovery: Could not parse discovery log:", err)
		}

		switch dicoveryLog.Type {
		case RESOURCE_TYPE:
			if err := json.Unmarshal(logMsg, &result.Resources); err != nil {
				log.Fatalln("RunDiscovery: Could not parse dependency:", err)
			}
		case TEST_TYPE:
			var test = Test{}
			if err := json.Unmarshal(logMsg, &test); err != nil {
				log.Fatalln("RunDiscovery: Could not parse test:", err)
			}
			result.Tests = append(result.Tests, test)
		}
	}

	resultCh <- result
}
