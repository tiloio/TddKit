package runner

import (
	_ "embed"
	"log"
	"test-frameworl/internal/ecmascript"
	"test-frameworl/internal/discover"
)

type TestRun struct {
	DiscoveryResults []discover.Result `json:"discoveryResults"`
	Running          []discover.Result `json:"running"`
	// Errors           []TestResult      `json:"errors"`
	// Finished         []TestResult      `json:"finished"`
	// resultChannel    chan TestResult
}

func RunAllTests(files *[]string) {
	var parsedFiles = ecmascript.ParseFiles(files)

	discoveryResultCh := make(chan DiscoveryResult, filesLength)
	for _, file := range parsedFiles {
		go RunDiscoveryPhase(file, discoveryResultCh, logs)
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

	run.runPossibleTests(logs)

	log.Println("Currently running", len(run.Running))
	for len(run.Running) != 0 {
		run.evaluateRunningTest()
		run.runPossibleTests(logs)
		go transmitStat(statsChannel, run)
	}
	transmitStat(statsChannel, run)

	var tests = 0
	var erroredTests = 0
	for _, test := range run.Finished {
		tests = tests + len(test.Success)
	}
	for _, test := range run.Errors {
		tests = tests + len(test.Success) + len(test.Errors)
		erroredTests = erroredTests + len(test.Errors)
	}

	log.Println(len(run.Finished)+len(run.Errors), "Files -", len(run.Errors), "Failed")
	log.Println(tests, "Tests -", erroredTests, "Failed")
}

func (run *TestRun) evaluateRunningTest() {
	currentResult := <-run.resultChannel

	for index, discoveryItem := range run.Running {
		if discoveryItem.File.Name == currentResult.DiscoveryResult.File.Name {
			if len(currentResult.Errors) == 0 {
				run.Finished = append(run.Finished, currentResult)
			} else {
				run.Errors = append(run.Errors, currentResult)
			}

			run.Running = remove(run.Running, index)
			return
		}
	}

	log.Panicln("Finished test not found in running tests:", currentResult.DiscoveryResult.File.Name)
}

const (
	STILL_WAITING             = -1
	ERRORED_DEPENDENCY        = 0
	ALL_DEPENDENCIES_FINISHED = 1
)

func (run *TestRun) runPossibleTests(logs chan CommandLog) {

	var leftoverResults = make([]DiscoveryResult, 0)

	for _, result := range run.DiscoveryResults {

		var dependencyStatus, dependency = CheckDependencies(result, run.Errors, run.Finished)
		log.Println("dependencyStatus", result.File.Name, dependencyStatus)

		if dependencyStatus == STILL_WAITING {
			leftoverResults = append(leftoverResults, result)
			continue
		}
		if dependencyStatus == ERRORED_DEPENDENCY {
			run.Errors = append(run.Errors, TestResult{
				DiscoveryResult:   result,
				DependencyErrored: *dependency,
			})
			continue
		}

		go RunTest(result, run.resultChannel, logs)
		run.Running = append(run.Running, result)
	}

	run.DiscoveryResults = leftoverResults
}

func CheckDependencies(result DiscoveryResult, errors []TestResult, finished []TestResult) (int, *Dependency) {
	for _, dependency := range result.TestSuite.Dependencies {

		if resultMachtes(errors, dependency.Id) {
			return ERRORED_DEPENDENCY, &dependency
		}

		if !resultMachtes(finished, dependency.Id) {
			return STILL_WAITING, &dependency
		}
	}

	return ALL_DEPENDENCIES_FINISHED, nil
}

func resultMachtes(s []TestResult, id string) bool {
	for _, result := range s {
		if id == result.DiscoveryResult.TestSuite.Id {
			return true
		}
	}

	return false
}

func remove(s []DiscoveryResult, i int) []DiscoveryResult {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
