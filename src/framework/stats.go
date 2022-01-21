package main

import (
	"encoding/json"
	"log"
	"os"
)

type StatLogMessage struct {
	Time      int64  `json:"time"`
	Stderr    bool   `json:"stderr"`
	Framework bool   `json:"framework"`
	Message   string `json:"message"`
}

func readAndSaveLogs(logs chan CommandLog) {
	file, err := os.Create("./log.json")
	if err != nil {
		log.Panicln("Could not create log file", err)
	}
	defer file.Close()

	for logMsg := range logs {
		statLog := StatLogMessage{
			Time:      logMsg.time.UnixMilli(),
			Stderr:    logMsg.stderr,
			Framework: logMsg.framework,
			Message:   string(logMsg.message),
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
}
