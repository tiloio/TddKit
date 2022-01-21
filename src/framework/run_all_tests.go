package main

import (
	_ "embed"
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
)

type TestRun struct {
	DiscoveryResults []DiscoveryResult `json:"discoveryResults"`
	Running          []DiscoveryResult `json:"running"`
	Errors           []TestResult      `json:"errors"`
	Finished         []TestResult      `json:"finished"`
	resultChannel    chan TestResult
}

var esModule = flag.Bool("esm", false, "Is code es module?")

func RunAllTests(files *[]string) {
	var filesLength = len(*files)
	parsedFileChannel := make(chan ParsedFile, filesLength)
	var logs = make(chan CommandLog)
	go readAndSaveLogs(logs)

	for _, file := range *files {
		go ParseFileAsync(file, esModule, parsedFileChannel)
	}

	var parsedFiles = make([]ParsedFile, filesLength)

	for i := 0; i < filesLength; i++ {
		parsedFiles[i] = <-parsedFileChannel
	}

	discoveryResultCh := make(chan DiscoveryResult, filesLength)
	for _, file := range parsedFiles {
		go RunDiscoveryPhase(file, discoveryResultCh, logs)
	}

	var run = TestRun{
		DiscoveryResults: make([]DiscoveryResult, filesLength),
		Running:          make([]DiscoveryResult, 0),
		Errors:           make([]TestResult, 0),
		Finished:         make([]TestResult, 0),
		resultChannel:    make(chan TestResult, filesLength),
	}

	for index := range run.DiscoveryResults {
		run.DiscoveryResults[index] = <-discoveryResultCh
	}

	file, _ := json.MarshalIndent(run, "", " ")
	_ = ioutil.WriteFile("debug.json", file, 0644)

	run.runPossibleTests(logs)

	log.Println("Currently running", len(run.Running))
	for len(run.Running) != 0 {
		run.evaluateRunningTest()
		run.runPossibleTests(logs)
	}

	var tests = 0
	var erroredTests = 0
	for _, test := range run.Finished {
		tests = tests + test.tests
	}
	for _, test := range run.Errors {
		tests = tests + test.tests
		erroredTests = erroredTests + test.errors
	}

	log.Println(len(run.Finished)+len(run.Errors), "Files -", len(run.Errors), "Failed")
	log.Println(tests, "Tests -", erroredTests, "Failed")
}

func (run *TestRun) evaluateRunningTest() {
	currentResult := <-run.resultChannel

	for index, discoveryItem := range run.Running {
		if discoveryItem.File.Name == currentResult.discoveryResult.File.Name {
			if currentResult.errors == 0 {
				run.Finished = append(run.Finished, currentResult)
			} else {
				run.Errors = append(run.Errors, currentResult)
			}

			run.Running = remove(run.Running, index)
			return
		}
	}

	log.Panicln("Finished test not found in running tests:", currentResult.discoveryResult.File.Name)
}

const (
	STILL_WAITING             = -1
	ERRORED_DEPENDENCY        = 0
	ALL_DEPENDENCIES_FINISHED = 1
)

func (run *TestRun) runPossibleTests(logs chan CommandLog) {

	var leftoverResults = make([]DiscoveryResult, 0)

	for _, result := range run.DiscoveryResults {

		var dependencyStatus = CheckDependencies(result, run.Errors, run.Finished)
		log.Println("dependencyStatus", result.File.Name, dependencyStatus)

		if dependencyStatus == STILL_WAITING {
			leftoverResults = append(leftoverResults, result)
			continue
		}
		if dependencyStatus == ERRORED_DEPENDENCY {
			failedTests := len(result.Tests)
			run.Errors = append(run.Errors, TestResult{
				discoveryResult: result,
				dependencyError: true,
				tests:           failedTests,
				errors:          failedTests,
			})
			continue
		}

		run.Running = append(run.Running, result)
		go RunTest(result, run.resultChannel, logs)
	}

	run.DiscoveryResults = leftoverResults
}

func CheckDependencies(result DiscoveryResult, errors []TestResult, finished []TestResult) int {
	for _, dependency := range result.Dependency.Dependencies {

		if resultMachtes(errors, dependency.Id) {
			return ERRORED_DEPENDENCY
		}

		if !resultMachtes(finished, dependency.Id) {
			return STILL_WAITING
		}
	}

	return ALL_DEPENDENCIES_FINISHED
}

func resultMachtes(s []TestResult, id string) bool {
	for _, result := range s {
		if id == result.discoveryResult.Dependency.Id {
			return true
		}
	}

	return false
}

func remove(s []DiscoveryResult, i int) []DiscoveryResult {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
