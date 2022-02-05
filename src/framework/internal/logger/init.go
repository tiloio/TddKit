package logger

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
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


func (command *Command) log() {
	var path = statsPath("log.json")
	statsDir(path)

	file, err := os.Create(path)
	if err != nil {
		log.Panicln("Could not create log file", err)
	}
	defer file.Close()

	statLog := statLogMessage{
		Time:      command.time.UnixMilli(),
		Stderr:    command.stderr,
		Framework: command.framework,
		Type:      command.executionType,
		Message:   string(command.message),
	}

	statLogJson, err := json.MarshalIndent(statLog, "", " ")
	if err != nil {
		log.Panicln("Could not create stat log json", err)
	}

	_, err = file.Write(append(statLogJson, []byte("\n")...))
	if err != nil {
		log.Panicln("Could not write to log file", err)
	}
}

func statsPath(fileName string) string {
	return filepath.Join(".", "stats", fileName)
}

func statsDir(path string) {
	err := os.MkdirAll(filepath.Dir(path), os.ModePerm)
	if err != nil {
		log.Panicln("Could not create stats dir", err)
	}
}
