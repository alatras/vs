package config

import (
	"errors"
)

var Version = "unknown"
var AppName = "Validation Service"

// ServerConfiguration with command line flags and env
type Server struct {
	Mongo Mongo `yaml:"mongo"`

	HTTPPort int `yaml:"httpPort"`

	Log Log `yaml:"log"`
}

// Mongo MongoDB configuration parameters
type Mongo struct {
	URL string `yaml:"url"`
	DB  string `yaml:"db"`
}

func (s Server) Validate() error {
	if s.HTTPPort == 0 {
		return errors.New("httpPort is missing")
	}

	if s.Mongo.URL == "" {
		return errors.New("mongo.url is missing")
	}

	if s.Mongo.DB == "" {
		return errors.New("mongo.db is missing")
	}

	if s.Log.Format == "" {
		return errors.New("log.format is missing")
	}

	if s.Log.Level == "" {
		return errors.New("log.level is missing")
	}

	return nil
}
