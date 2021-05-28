package config

import (
	"errors"
	"os"
)

var Version = "unknown"
var AppName = "Validation Service"

// ServerConfiguration with command line flags and env
type Server struct {
	Mongo Mongo `yaml:"mongo"`

	HTTPPort int `yaml:"httpPort"`

	AppD AppD `yaml:"appd"`

	Log Log `yaml:"log"`
}

// Mongo MongoDB configuration parameters
type Mongo struct {
	URL               string `yaml:"url"`
	DB                string `yaml:"db"`
	RetryMilliseconds int    `yaml:"retryMilliseconds"`
}

// GetConfig either from struct or environment
func (m Mongo) GetConfig(key string) string {
	v := os.Getenv(key)

	if v != "" {
		return v
	}

	switch key {
	case "MONGO_URL":
		return m.URL
	case "MONGO_DB":
		return m.DB
	default:
		return "n/a"
	}
}

// DefaultMongoRetryMilliseconds default setting for Mongo RetryMilliseconds
const DefaultMongoRetryMilliseconds = 1

// AppD App Dynamics configuration parameters
type AppD struct {
	AppName     string `yaml:"appName"`
	TierName    string `yaml:"tierName"`
	NodeName    string `yaml:"nodeName"`
	InitTimeout int    `yaml:"initTimeout"`
	Controller  struct {
		Host      string `yaml:"host"`
		Port      uint16 `yaml:"port"`
		UseSSL    bool   `yaml:"useSSL"`
		Account   string `yaml:"account"`
		AccessKey string `yaml:"accessKey"`
	} `yaml:"controller"`
}

func (c Server) Validate() error {
	if c.HTTPPort == 0 {
		return errors.New("httpPort is missing")
	}

	if err := c.Mongo.Validate(); err != nil {
		return err
	}

	if err := c.AppD.Validate(); err != nil {
		return err
	}

	if err := c.Log.Validate(); err != nil {
		return err
	}

	return nil
}

func (c Mongo) Validate() error {
	if c.GetConfig("url") == "" {
		return errors.New("mongo.url is missing")
	}

	if c.GetConfig("db") == "" {
		return errors.New("mongo.db is missing")
	}

	return nil
}

func (c AppD) Validate() error {
	if c.AppName == "" {
		return errors.New("appd.appName is missing")
	}

	if c.TierName == "" {
		return errors.New("appd.tierName is missing")
	}

	if c.NodeName == "" {
		return errors.New("appd.nodeName is missing")
	}

	if c.InitTimeout == 0 {
		return errors.New("appd.initTimeout is missing")
	}

	if c.Controller.Host == "" {
		return errors.New("appd.controller.host is missing")
	}

	if c.Controller.Port == 0 {
		return errors.New("appd.controller.port is missing")
	}

	if c.Controller.Account == "" {
		return errors.New("appd.controller.account is missing")
	}

	if c.Controller.AccessKey == "" {
		return errors.New("appd.controller.accessKey is missing")
	}

	return nil
}

func (c Log) Validate() error {
	if c.Format == "" {
		return errors.New("log.format is missing")
	}

	if c.Level == "" {
		return errors.New("log.level is missing")
	}

	return nil
}
