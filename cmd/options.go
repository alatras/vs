package cmd

import (
	"bitbucket.verifone.com/validation-service/logger"
	"github.com/sirupsen/logrus"
)

// Options all configuration parameters
type Options struct {
	HTTPPort int `long:"httpPort" env:"HTTP_PORT" default:"8080" description:"HTTP port"`

	Log   LogGroup   `group:"log" namespace:"log"`
	Mongo MongoGroup `group:"mongo" namespace:"mongo"`
}

// MongoGroup MongoDB configuration parameters
type MongoGroup struct {
	URL string `long:"url" env:"MONGO_URL" required:"MongoDB url required" description:"MongoDB url"`
	DB  string `long:"db" env:"MONGO_DB" default:"validationService" description:"MongoDB database name"`
}

// LogGroup logging configuration parameters
type LogGroup struct {
	Level  string `long:"level" env:"LOG_LEVEL" default:"info" choice:"trace" choice:"debug" choice:"info" choice:"warn" choice:"error" choice:"fatal" description:"Logging level"`
	Format string `long:"format" env:"LOG_FORMAT" default:"json" choice:"json" choice:"text" description:"Logging format"`
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
