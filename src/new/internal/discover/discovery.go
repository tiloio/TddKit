
const TESTSUITE_TYPE = "TESTSUITE"
const TEST_TYPE = "TEST"

var newLineAsByte = []byte("\n")
var discoveryEnvironmentVariables = []string{DISCOVERY_PHASE_ENVIRONMENT_VAIRABLE}

func discover(file ParsedFile, resultCh chan DiscoveryResult, logger chan CommandLog) {
	var executeLog = ExecuteLog{
		channel: make(chan CommandLog),
		typ:     "DISCOVERY",
	}
	go ExecuteEcmascriptTests(&file.content, &discoveryEnvironmentVariables, &executeLog)

	var result = DiscoveryResult{
		File:  file,
		Tests: make([]Test, 0),
	}

	var lastTestSuiteLogIndex = 0
	var logMessages = make([]LogMessage, 0)

	for logMsg := range executeLog.channel {
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

		if dicoveryLog.Type == TESTSUITE_TYPE {
			lastTestSuiteLogIndex = len(logMessages) - 1
		}
	}

	for i := lastTestSuiteLogIndex; i < len(logMessages); i++ {
		var logMsg = logMessages[i]

		switch logMsg.discoveryLog.Type {
		case TESTSUITE_TYPE:
			if err := json.Unmarshal(logMsg.message.message, &result.TestSuite); err != nil {
				log.Fatalln("RunDiscovery: Could not parse test suite:", err)
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
