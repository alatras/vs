package logger

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

type LogFormat = string

const (
	TextFormat LogFormat = "text"
	JsonFormat LogFormat = "json"
)

type Logger struct {
	Output *logrus.Entry
	Error *logrus.Entry
}

func NewLogger(appName string, appVersion string, format LogFormat, level logrus.Level) (*Logger, error) {
	logFields := logrus.Fields{
		"name": appName,
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
		return nil, errors.New(fmt.Sprintf("Invalid log format %s", format))
	}

	logger := logrus.New()
	logger.SetLevel(level)
	logger.SetOutput(os.Stdout)
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
		Error: errorLogger.WithFields(logFields),
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

func (l *Logger) WithMetadata(metadata interface{}) *Logger {
	return &Logger{
		Output: l.Output.WithField("metadata", metadata),
		Error:  l.Error.WithField("metadata", metadata),
	}
}
