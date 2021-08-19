package logger

import (
	"fmt"
	"io"
	"os"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
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

var AppName, AppVersion string

func NewLogger(
	appName string,
	appVersion string,
	format interface{},
	level logrus.Level,
	logfile string,
	logFileMaxMb int,
	logFileRotationCount int,
	logFileRotationDays int,
) (*Logger, error) {
	logFields := logrus.Fields{
		"mdc": logrus.Fields{
			"app_name":    appName,
			"app_version": appVersion,
		},
	}

	AppName = appName
	AppVersion = appVersion

	var formatter logrus.Formatter

	if format == TextFormat {
		formatter = &logrus.TextFormatter{
			ForceColors: true,
		}
	} else if format == JsonFormat {
		formatter = &logrus.JSONFormatter{
			FieldMap: logrus.FieldMap{
				logrus.FieldKeyMsg:  "logger_name",
				logrus.FieldKeyTime: "timestamp",
			},
		}
	} else {
		return nil, fmt.Errorf("invalid log format %s", format)
	}

	logFile := &lumberjack.Logger{
		Filename:   logfile,
		MaxSize:    logFileMaxMb,         // megabytes
		MaxBackups: logFileRotationCount, // no. of files
		MaxAge:     logFileRotationDays,  // days
		Compress:   true,                 // default is false
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
		Output: l.Output.WithField("trace_id", traceId),
		Error:  l.Error.WithField("trace_id", traceId),
	}
}

func (l *Logger) WithMetadata(metadata Metadata) *Logger {
	return &Logger{
		Output: l.Output.WithField("metadata", metadata),
		Error:  l.Error.WithField("metadata", metadata),
	}
}

func (l *Logger) WithCorrelationId(correlationId string) *Logger {
	return &Logger{
		Output: l.Output.WithField("correlation_id", correlationId),
		Error:  l.Error.WithField("correlation_id", correlationId),
	}
}
