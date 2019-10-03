package cmd

import (
	"bitbucket.verifone.com/validation-service/logger"
	"github.com/sirupsen/logrus"
)

var Version = "unknown"
var AppName = "Validation Service"

// CommonOptionsCommander extends flags.Commander with SetCommon
// All commands should implement this interfaces
type CommonOptionsCommander interface {
	SetCommon(commonOpts CommonOpts)
	Execute(args []string) error
}

// LogGroup logging configuration parameters
type LogGroup struct {
	///nolint:staticcheck
	Level string `long:"level" env:"LOG_LEVEL" default:"info" choice:"trace" choice:"debug" choice:"info" choice:"warn" choice:"error" choice:"fatal" description:"Logging level"`
	///nolint:staticcheck
	Format string `long:"format" env:"LOG_FORMAT" default:"json" choice:"json" choice:"text" description:"Logging format"`
}

type CommonOpts struct {
	Log LogGroup
}

func (log LogGroup) LevelValue() logrus.Level {
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

func (log LogGroup) FormatValue() logger.LogFormat {
	switch log.Format {
	case "text":
		return logger.TextFormat
	default:
		return logger.JsonFormat
	}
}

// SetCommon satisfies CommonOptionsCommander interface and sets common option fields
func (c *CommonOpts) SetCommon(commonOpts CommonOpts) {
	c.Log = commonOpts.Log
}
