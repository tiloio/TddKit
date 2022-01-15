package main

import (
	_ "embed"
	"encoding/json"
	"log"
)

type DiscoveryResult struct {
	file       ParsedFile
	dependency Dependency
	resources  []Resource
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

type DiscoveryLog struct {
	Type string `json:"type"`
}

const DEPENDENCY_TYPE = "DEPENDENCY"
const RESOURCE_TYPE = "RESOURCE"

var newLineAsByte = []byte("\n")
var discoveryEnvironmentVariables = []string{DISCOVERY_PHASE_ENVIRONMENT_VAIRABLE}

func RunDiscoveryPhase(file ParsedFile, resultCh chan DiscoveryResult) {
	stdout := ExecuteEcmascriptTests(&file.content, &discoveryEnvironmentVariables)

	var logs = ReadLogs(stdout)

	var dependency = Dependency{}
	var resources = Resources{}
	for _, logMsg := range logs {

		var dicoveryLog = DiscoveryLog{}
		if err := json.Unmarshal(logMsg, &dicoveryLog); err != nil {
			log.Fatalln("RunDiscovery: Could not parse dependency:", err)
		}

		switch dicoveryLog.Type {
		case DEPENDENCY_TYPE:
			if err := json.Unmarshal(logMsg, &dependency); err != nil {
				log.Fatalln("RunDiscovery: Could not parse dependency:", err)
			}
		case RESOURCE_TYPE:
			if err := json.Unmarshal(logMsg, &resources); err != nil {
				log.Fatalln("RunDiscovery: Could not parse dependency:", err)
			}
		}
	}

	resultCh <- DiscoveryResult{
		file:       file,
		dependency: dependency,
		resources:  resources.Resources,
	}
}
