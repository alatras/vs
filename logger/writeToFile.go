package logger

import (
	"encoding/json"
	"log"
)

// WriteToFile : write logs into log file
func WriteToFile(fl FileLogger) {
	m, err := json.Marshal(fl)

	if err != nil {
		log.Printf("[ERROR] failed to marshal log with error: %+v", err)
		return
	}

	l := log.New(LogFile, string(m), 0)

	l.Println()
}
