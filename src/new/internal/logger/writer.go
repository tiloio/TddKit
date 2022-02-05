package logger

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

var file *os.File

func init() {
	var path = statsPath("log.json")
	statsDir(path)

	curFile, err := os.Create(path)
	if err != nil {
		log.Panicln("Could not create log file", err)
	}
	file = curFile
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

func writeLog(statLog interface{}) {
	statLogJson, err := json.MarshalIndent(statLog, "", " ")
	if err != nil {
		log.Panicln("Could not create stat log json", err)
	}

	_, err = file.Write(append(statLogJson, []byte("\n")...))
	if err != nil {
		log.Panicln("Could not write to log file", err)
	}
}
