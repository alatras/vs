package logger

import (
	"fmt"
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

type LogFormat = string

const (
	TextFormat LogFormat = "text"
	JsonFormat LogFormat = "json"
)

type Metadata = map[string]interface{}

type Logger struct {
	Output *logrus.Entry
	Error  *logrus.Entry
}

func NewStubLogger() *Logger {
	logger := logrus.New()
	logger.SetLevel(logrus.PanicLevel)

	errorLogger := logrus.New()
	errorLogger.SetLevel(logrus.PanicLevel)

	return &Logger{
		Output: logger.WithField("stub", true),
		Error:  errorLogger.WithField("stub", true),
	}
}

func NewLogger(appName string, appVersion string, format LogFormat, level logrus.Level) (*Logger, error) {
	logFields := logrus.Fields{
		"name":    appName,
		"version": appVersion,
	}

	var formatter logrus.Formatter

	if format == TextFormat {
		formatter = &logrus.TextFormatter{
			ForceColors: true,
		}
	} else if format == JsonFormat {
		formatter = &logrus.JSONFormatter{}
	} else {
		return nil, fmt.Errorf("invalid log format %s", format)
	}

	logFile, err := ProvideLogFile()

	if err != nil {
		return nil, fmt.Errorf("[ERROR] failed to provide log file with error: %+v", err)
	}

	mw := io.MultiWriter(os.Stdout, logFile)

	logger := logrus.New()
	logger.SetLevel(level)
	logger.SetOutput(mw)
	logger.SetFormatter(formatter)

	errorLogger := logrus.New()
	errorLogger.SetOutput(os.Stderr)
	errorLogger.SetFormatter(formatter)

	if level < logrus.ErrorLevel {
		errorLogger.SetLevel(level)
	} else {
		errorLogger.SetLevel(logrus.ErrorLevel)
	}

	return &Logger{
		Output: logger.WithFields(logFields),
		Error:  errorLogger.WithFields(logFields),
	}, nil
}

func (l *Logger) Scoped(scope string) *Logger {
	return &Logger{
		Output: l.Output.WithField("scope", scope),
		Error:  l.Error.WithField("scope", scope),
	}
}

func (l *Logger) WithTraceId(traceId string) *Logger {
	return &Logger{
		Output: l.Output.WithField("traceId", traceId),
		Error:  l.Error.WithField("traceId", traceId),
	}
}

func (l *Logger) WithMetadata(metadata Metadata) *Logger {
	return &Logger{
		Output: l.Output.WithField("metadata", metadata),
		Error:  l.Error.WithField("metadata", metadata),
	}
}
