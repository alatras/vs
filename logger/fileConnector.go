package logger

import (
	"errors"
	"log"
	"os"
	"os/exec"
	"time"
	"validation-service/environment"
)

// LogFile : current logfile
var LogFile *os.File

// Initiate current logfile get it ready
func init() {
	f, err := provideLogFile()

	if err != nil {
		log.Printf("[ERROR] failed to provide log file with error: %+v", err)
		return
	}

	LogFile = f

	err = cleanUp()

	if err != nil {
		log.Printf("[WARN] can't clean up logs: %+v", err)
		return
	}
}

// Provide current logfile
func provideLogFile() (*os.File, error) {
	path := "./logs/"
	fileName := "ValidationService_" + time.Now().Format("20060102") + ".log"

	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, 0700)

		if err != nil {
			return nil, errors.New("failed to create log path")
		}
	}

	f, err := os.OpenFile(path+fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		return nil, errors.New("failed to open log file")
	}

	return f, nil
}

// Delete logfiles older than rotation days' count
func cleanUp() error {
	rotationCount := environment.Get("LOG_ROTATING_COUNT", "30")

	_, err := exec.Command("find ./logs -type f -ctime +" + rotationCount + " | xargs rm -f").Output()

	if err != nil {
		return err
	}

	return nil
}
