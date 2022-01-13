package main

import (
	_ "embed"
	"encoding/json"
	"log"
)

type DiscoveryResult struct {
	Dependencies []Dependency
}

type Dependency struct {
	Id           string       `json:"id"`
	Dependencies []Dependency `json:"dependencies"`
}

type DiscoveryLog struct {
	Type string `json:"type"`
}

var newLineAsByte = []byte("\n")
var discoveryEnvironmentVariables = []string{DISCOVERY_PHASE_ENVIRONMENT_VAIRABLE}

func RunDiscoveryPhase(file ParsedFile, resultCh chan DiscoveryResult) {
	stdout := ExecuteEcmascriptTests(&file.content, &discoveryEnvironmentVariables)

	var logs = ReadLogs(stdout)

	var dependencies = []Dependency{}
	for _, logMsg := range logs {

		var dicoveryLog = DiscoveryLog{}
		if err := json.Unmarshal(logMsg, &dicoveryLog); err != nil {
			log.Fatalln("RunDiscovery: Could not parse dependency:", err)
		}

		switch dicoveryLog.Type {
		case "DEPENDENCY":
			var dependency = Dependency{}
			if err := json.Unmarshal(logMsg, &dependency); err != nil {
				log.Fatalln("RunDiscovery: Could not parse dependency:", err)
			}
			dependencies = append(dependencies, dependency)
		}
	}

	resultCh <- DiscoveryResult{
		Dependencies: dependencies,
	}
}
