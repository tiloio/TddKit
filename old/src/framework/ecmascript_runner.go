package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

const LOG_PREFIX = "__TFW:"

var LOG_PREFIX_BYTES = len([]byte(LOG_PREFIX))

type ExecuteCommand struct {
	command *exec.Cmd
	stdout  *bufio.Scanner
	stderr  *bufio.Scanner
}
type CommandLog struct {
	time          time.Time
	stderr        bool
	framework     bool
	message       []byte
	last          bool
	executionType string
}

type ExecuteLog struct {
	channel chan CommandLog
	typ     string
}

func ExecuteEcmascriptTests(testCode *[]byte, environmentVariables *[]string, executeLog *ExecuteLog) {

	var cmd *exec.Cmd
	if *esModule {
		cmd = exec.Command("node", "--enable-source-maps", "--input-type=module", "-")
	} else {
		cmd = exec.Command("node", "--enable-source-maps", "-")
	}
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, *environmentVariables...)

	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal(executeLog.typ, "Could not read stderr of exec node command!", err)
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(executeLog.typ, "Could not read stdout of exec node command!", err)
	}

	var execCommand = ExecuteCommand{
		command: cmd,
		stdout:  bufio.NewScanner(stdout),
		stderr:  bufio.NewScanner(stderr),
	}

	var finishedLogReadsChannel = make(chan int)

	go execCommand.readLogs(executeLog, finishedLogReadsChannel)

	buffer := bytes.Buffer{}
	buffer.Write(*testCode)
	cmd.Stdin = &buffer

	if err := cmd.Start(); err != nil {
		log.Fatal(executeLog.typ, "Could not start exec node command", err)
	}

	for i := 0; i < 1; i++ {
		<-finishedLogReadsChannel
	}

	if err := cmd.Wait(); err != nil {
		log.Fatal(executeLog.typ, "Exec node command failed", err)
	}
}

func (command *ExecuteCommand) readLogs(executeLog *ExecuteLog, finishedLogReadsChannel chan int) {
	var combinedChannel = make(chan CommandLog)

	go scanLog(command.stderr, true, combinedChannel, executeLog.typ)
	go scanLog(command.stdout, false, combinedChannel, executeLog.typ)

	var lastLogCounter = 0
	for msg := range combinedChannel {
		if msg.last {
			lastLogCounter++
		} else {
			executeLog.channel <- msg
		}

		if lastLogCounter == 2 {
			close(executeLog.channel)
			finishedLogReadsChannel <- 1
			return
		}
	}

}

func scanLog(scanner *bufio.Scanner, stderr bool, logChannel chan CommandLog, executionType string) {
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			log.Panicln(executionType, "Log scan error", err)
		}

		if len(scanner.Bytes()) == 0 {
			continue
		}

		var log = CommandLog{
			time:          time.Now(),
			stderr:        stderr,
			framework:     false,
			executionType: executionType,
			message:       make([]byte, len(scanner.Bytes())),
		}

		copy(log.message, scanner.Bytes())

		if strings.HasPrefix(scanner.Text(), LOG_PREFIX) {
			log.framework = true
			log.message = log.message[LOG_PREFIX_BYTES:]
		}

		logChannel <- log
	}
	logChannel <- CommandLog{last: true}
}
