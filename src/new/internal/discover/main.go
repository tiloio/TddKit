package discover

import (
	"test-framework/internal/logger"
	"test-framework/internal/parser"
)

type Result struct {
	File      parser.File `json:"file"`
	TestSuite TestSuite   `json:"testSuite"`
	Tests     []Test      `json:"tests"`
}

type TestSuite struct {
	Id           string       `json:"id"`
	Dependencies []Dependency `json:"dependencies"`
	Resources    []Resource   `json:"resources"`
}

type Dependency struct {
	Id           string       `json:"id,omitempty"`
	Dependencies []Dependency `json:"dependencies,omitempty"`
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
	message      logger.Command
}

func Run(parsedFiles *[]parser.File) {
	discoveryResultCh := make(chan Result, len(parsedFiles))
	for _, file := range parsedFiles {
		go discover(file, discoveryResultCh, logs)
	}

	var statsChannel = make(chan Stats)
	var run = TestRun{
		DiscoveryResults: make([]DiscoveryResult, filesLength),
		Running:          make([]DiscoveryResult, 0),
		Errors:           make([]TestResult, 0),
		Finished:         make([]TestResult, 0),
		resultChannel:    make(chan TestResult, filesLength),
	}
	go readAndSaveCurrentStats(statsChannel)

	for index := range run.DiscoveryResults {
		run.DiscoveryResults[index] = <-discoveryResultCh
		go transmitStat(statsChannel, run)
	}
}