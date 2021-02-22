package logger

import (
	"errors"
	"log"
	"os"
	"os/exec"
	"time"
	"validation-service/environment"
	"validation-service/utils"

	"github.com/go-co-op/gocron"
)

var (
	// LogFile for main logger
	LogFile            *os.File
	logDirectory       string = "./logs"
	currentLogFileName string = "ValidationService.log"
	rotatedFileInitial string = "ValidationService_"
	rotatedDirectory   string = logDirectory + "/rotated"
)

// Initiate & provision current logfile
func init() {
	f, err := ProvideLogFile()

	if err != nil {
		log.Printf("[ERROR] failed to provide log file with error: %+v", err)
		return
	}

	LogFile = f

	logFileCron()

	err = cleanUp()

	if err != nil {
		log.Printf("[WARN] can't clean up logs: %+v", err)
		return
	}
}

// ProvideLogFile provides current logfile and creates it if non existant
func ProvideLogFile() (*os.File, error) {
	if _, err := os.Stat(logDirectory); os.IsNotExist(err) {
		err = os.MkdirAll(logDirectory, 0700)

		if err != nil {
			return nil, errors.New("[ERROR] failed to create log directory")
		}
	}

	f, err := os.OpenFile(logDirectory+"/"+currentLogFileName,
		os.O_RDWR|os.O_CREATE|os.O_APPEND,
		0666)

	if err != nil {
		return nil, errors.New("failed to open log file")
	}

	return f, nil
}

// Schedule a logfile rotation cron job at midnight
func logFileCron() {
	s := gocron.NewScheduler(time.UTC)

	job, err := s.Every(1).Day().At("00:00").Do(rotateLogFile)

	if err != nil {
		log.Printf("[ERROR] failed to start log rotate cron job: %+v", err)
		return
	}

	s.StartAsync()

	job.ScheduledAtTime()
}

// Rotate log file into rotation directory and create a fresh current file
func rotateLogFile() {
	currentLogFile := logDirectory + "/" + currentLogFileName

	if _, err := os.Stat(currentLogFile); os.IsNotExist(err) {
		log.Printf("[ERROR] failed to find log file to rotate it: %+v", err)
		return
	}

	if _, err := os.Stat(rotatedDirectory); os.IsNotExist(err) {
		err = os.MkdirAll(rotatedDirectory, 0700)

		if err != nil {
			log.Printf("[ERROR] failed to create logfile rotated directory: %+v", err)
			return
		}
	}

	// string: `ValidationService_[date in YYYYMMDD]_[RANDOM STRING].log`
	rotated := rotatedDirectory + "/" + rotatedFileInitial +
		time.Now().AddDate(0, 0, -1).Format("20060102") +
		"_" + utils.RandomString(8) + ".log"

	if _, err := utils.CopyFile(currentLogFile, rotated); err != nil {
		log.Printf("[ERROR] failed to copy today's log file: %+v", err)
		return
	}

	if err := os.Remove(currentLogFile); err != nil {
		log.Printf("[ERROR] failed to remove current log file: %+v", err)
		return
	}

	if _, err := os.Create(currentLogFile); err != nil {
		log.Printf("[ERROR] failed to create a logfile after rotation: %+v", err)
		return
	}
}

// Cleanup logfiles older than rotation days' count
func cleanUp() error {
	rotationCount := environment.Get("LOG_ROTATING_COUNT", "30")

	_, err := exec.Command("find ./logs/rotated -type f -ctime +" + rotationCount + " | xargs rm -f").Output()

	if err != nil {
		return err
	}

	return nil
}
