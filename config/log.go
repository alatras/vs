package config

import (
	"validation-service/logger"

	"github.com/sirupsen/logrus"
)

// Log logging configuration parameters
type Log struct {
	Level                 string `yaml:"level"`
	Format                string `yaml:"format"`
	LogFile               string `yaml:"logFile"`
	logFileMaxSize        int    `yaml:"logFileMaxSize"`
	logFileRotationPeriod int    `yaml:"logFileRotationPeriod"`
	logFileRotationCount  int    `yaml:"logFileRotationPeriod"`
	traceIdHeader         string `yaml:"traceIdHeader"`
}

func (log Log) LevelValue() logrus.Level {
	switch log.Level {
	case "trace":
		return logrus.TraceLevel
	case "debug":
		return logrus.DebugLevel
	case "info":
		return logrus.InfoLevel
	case "warn":
		return logrus.WarnLevel
	case "error":
		return logrus.WarnLevel
	case "fatal":
		return logrus.FatalLevel
	default:
		return logrus.InfoLevel
	}
}

func (log Log) FormatValue() logger.LogFormat {
	switch log.Format {
	case "text":
		return logger.TextFormat
	default:
		return logger.JsonFormat
	}
}

func (log Log) LogFileValue() logger.LogFormat {
	return log.LogFile
}

func (log Log) LogFileMaxMbValue() int {
	return log.logFileMaxSize
}

func (log Log) LogFileRotationCountValue() int {
	return log.logFileRotationCount
}

func (log Log) LogFileRotationDaysValue() int {
	return log.logFileRotationPeriod
}

func (log Log) TraceIdHeader() string {
	return log.traceIdHeader
}
