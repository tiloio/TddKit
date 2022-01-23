package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type StatLogMessage struct {
	Time      int64  `json:"time"`
	Stderr    bool   `json:"stderr"`
	Framework bool   `json:"framework"`
	Type      string `json:"type"`
	Message   string `json:"message"`
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

func readAndSaveLogs(logs chan CommandLog) {
	var path = statsPath("log.json")
	statsDir(path)

	file, err := os.Create(path)
	if err != nil {
		log.Panicln("Could not create log file", err)
	}
	defer file.Close()

	for logMsg := range logs {
		statLog := StatLogMessage{
			Time:      logMsg.time.UnixMilli(),
			Stderr:    logMsg.stderr,
			Framework: logMsg.framework,
			Type:      logMsg.executionType,
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

type Stats struct {
	DiscoveryResults []DiscoveryResult `json:"discoveryResults"`
	Running          []DiscoveryResult `json:"running"`
	Errors           []TestResult      `json:"errors"`
	Finished         []TestResult      `json:"finished"`
}

func transmitStat(statsChannel chan Stats, testRun TestRun) {
	statsChannel <- Stats{
		DiscoveryResults: testRun.DiscoveryResults,
		Running:          testRun.Running,
		Errors:           testRun.Errors,
		Finished:         testRun.Finished,
	}
}

func readAndSaveCurrentStats(statsChannel chan Stats) {
	var path = statsPath("stats.json")
	statsDir(path)

	for stats := range statsChannel {
		statLogJson, err := json.MarshalIndent(stats, "", " ")
		if err != nil {
			log.Panicln("Could not create stats json", err)
		}

		err = ioutil.WriteFile(path, statLogJson, 0644)
		if err != nil {
			log.Panicln("Could not create stats file", err)
		}
	}
}
