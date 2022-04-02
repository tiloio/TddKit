package logger

import (
	"fmt"
	"log"
	"time"
)

type Command struct {
	time          time.Time
	stderr        bool
	framework     bool
	message       []byte
	last          bool
	executionType string
}

type statLogMessage struct {
	Time      int64  `json:"time"`
	Stderr    bool   `json:"stderr"`
	Framework bool   `json:"framework"`
	Type      string `json:"type"`
	Message   string `json:"message"`
}
type statLogFrameworkMessage struct {
	Time    int64  `json:"time"`
	Type    string `json:"type"`
	Message string `json:"message"`
}

func (command *Command) Log() {
	statLog := statLogMessage{
		Time:      command.time.UnixMilli(),
		Stderr:    command.stderr,
		Framework: command.framework,
		Type:      command.executionType,
		Message:   string(command.message),
	}

	writeLog(statLog)
}

func Info(message string) {
	statLog := statLogFrameworkMessage{
		Time:    time.Now().UnixMilli(),
		Type:    "InfoLog",
		Message: message,
	}

	writeLog(statLog)
}

func Fatal(v ...interface{}) {
	message := fmt.Sprintf("%v", v)
	statLog := statLogFrameworkMessage{
		Time:    time.Now().UnixMilli(),
		Type:    "FatalErrorLog",
		Message: message,
	}

	writeLog(statLog)
	log.Fatal(v)
}
