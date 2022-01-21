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

type LogMessage struct {
	discoveryLog DiscoveryLog
	message      CommandLog
}

const DEPENDENCY_TYPE = "DEPENDENCY"
const RESOURCE_TYPE = "RESOURCE"
const TEST_TYPE = "TEST"

var newLineAsByte = []byte("\n")
var discoveryEnvironmentVariables = []string{DISCOVERY_PHASE_ENVIRONMENT_VAIRABLE}

func RunDiscoveryPhase(file ParsedFile, resultCh chan DiscoveryResult, logger chan CommandLog) {
	var logs = make(chan CommandLog)
	go ExecuteEcmascriptTests(&file.content, &discoveryEnvironmentVariables, logs)

	var result = DiscoveryResult{
		File:  file,
		Tests: make([]Test, 0),
	}

	var lastDependencLogIndex = 0
	var logMessages = make([]LogMessage, 0)

	for logMsg := range logs {
		logger <- logMsg
		logMessage := string(logMsg.message)

		if logMsg.stderr {
			log.Println(file.Name, "ERR:", logMessage)
		} else {
			log.Println(file.Name+":", logMessage)
		}

		if !logMsg.framework {
			continue
		}

		var dicoveryLog = DiscoveryLog{}
		if err := json.Unmarshal(logMsg.message, &dicoveryLog); err != nil {
			log.Fatalln("RunDiscovery: Could not parse discovery log:'", logMessage, "' Err:", err)
		}

		logMessages = append(logMessages, LogMessage{
			discoveryLog: dicoveryLog,
			message:      logMsg,
		})

		if dicoveryLog.Type == DEPENDENCY_TYPE {
			lastDependencLogIndex = len(logMessages) - 1
		}
	}

	for i := lastDependencLogIndex; i < len(logMessages); i++ {
		var logMsg = logMessages[i]

		switch logMsg.discoveryLog.Type {
		case DEPENDENCY_TYPE:
			if err := json.Unmarshal(logMsg.message.message, &result.Dependency); err != nil {
				log.Fatalln("RunDiscovery: Could not parse dependency:", err)
			}
		case RESOURCE_TYPE:
			if err := json.Unmarshal(logMsg.message.message, &result.Resources); err != nil {
				log.Fatalln("RunDiscovery: Could not parse dependency:", err)
			}
		case TEST_TYPE:
			var test = Test{}
			if err := json.Unmarshal(logMsg.message.message, &test); err != nil {
				log.Fatalln("RunDiscovery: Could not parse test:", err)
			}
			result.Tests = append(result.Tests, test)
		}
	}

	resultCh <- result
}
