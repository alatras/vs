package config

import (
	"fmt"
	"log"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

var Version = "1.0.12"
var AppName = "Validation Service"

type Server struct {
	HTTPPort int   `yaml:"httpPort" env:"HTTP_PORT" env-default:"8080" env-description:"App HTTP port"`
	Mongo    Mongo `yaml:"mongo"`
	Log      Log   `yaml:"log"`
	AppD     AppD  `yaml:"appd"`
}

type Mongo struct {
	URL               string `yaml:"url" env:"MONGO_URL" env-default:"mongodb://mongo" env-description:"Mongo database URL"`
	DB                string `yaml:"db" env:"MONGO_DB_NAME" env-default:"validationService" env-description:"Mongo database name"`
	RetryMilliseconds int    `yaml:"retryMilliseconds" env:"MONGO_RETRY_MILLISECONDS" env-default:"0" env-description:"Mondo DB retry milliseconds"`
}

type AppD struct {
	AppName     string `yaml:"appName" env:"APP_DYNAMICS_APP_NAME" env-default:"Validation Service" env-description:"App name"`
	TierName    string `yaml:"tierName" env:"APP_DYNAMICS_TIRENAME" env-default:"transaction" env-description:"App dynamics tier name"`
	NodeName    string `yaml:"nodeName" env:"APP_DYNAMICS_NODE_NAME" env-default:"transaction01" env-description:"App dynamics node name"`
	InitTimeout int    `yaml:"initTimeout" env:"APP_DYNAMICS_INIT_TIMEOUT" env-default:"1000" env-description:"App dynamics node name"`
	Controller  struct {
		Host      string `yaml:"host" env:"APP_DYNAMICS_HOST" env-default:"" env-description:"App Dynamics host"`
		Port      uint16 `yaml:"port" env:"APP_DYNAMICS_PORT" env-default:"443" env-description:"App Dynamics port"`
		ProxyHost string `yaml:"proxyHost" env:"APP_DYNAMICS_PROXY_HOST" env-default:"" env-description:"App Dynamics proxy host needed for some environments"`
		ProxyPort string `yaml:"proxyPort" env:"APP_DYNAMICS_PROXY_PORT" env-default:"" env-description:"App Dynamics proxy port needed for some environments"`
		UseSSL    bool   `yaml:"useSSL" env:"APP_DYNAMICS_PROXY_PORT" env-default:"true" env-description:"App Dynamics use SSL"`
		Account   string `yaml:"account" env:"APP_DYNAMICS_ACCOUNT" env-default:"account" env-description:"App Dynamics account"`
		AccessKey string `yaml:"accessKey" env:"APP_DYNAMICS_ACCESS_KEY" env-default:"accessKey" env-description:"App Dynamics access key"`
	} `yaml:"controller"`
}

type Log struct {
	Level                        string `yaml:"level" env:"LOG_LEVEL" env-default:"info" env-description:"Log level"`
	Format                       string `yaml:"format" env:"LOG_FORMAT" env-default:"json" env-description:"Log format"`
	TraceIdHeader                string `yaml:"traceIdHeader" env:"TRACE_ID_HEADER" env-default:"x-b3-traceid" env-description:"Trace ID header key name"`
	LogFile                      string `yaml:"logFile" env:"LOG_FILE" env-default:"./logs/main.log" env-description:"Path of the log file"`
	LogFileMaxMb                 int    `yaml:"logFileMaxMb" env:"LOG_FILE_MAX_SIZE" env-default:"1" env-description:"Log file max size"`
	LogRotationPeriod            int    `yaml:"logRotatingPeriod" env:"LOG_ROTATING_PERIOD" env-default:"1" env-description:"Log file rotation period"`
	LogRotationCount             int    `yaml:"logRotatingCount" env:"LOG_ROTATING_COUNT" env-default:"30" env-description:"Log file rotation count"`
	HealthLogFilePath            string `yaml:"healthCheckLogFile" env:"HEALTH_LOG_FILE" env-default:"./logs/health.log" env-description:"Path of the health check log file"`
	HealthCheckLogRotatingPeriod int    `yaml:"healthCheckLogRotatingPeriod" env:"HEALTH_CHECK_LOG_ROTATING_PERIOD" env-default:"10" env-description:"Log file rotation period"`
	HealthCheckLogRotationCount  int    `yaml:"healthCheckLogRotatingCount" env:"HEALTH_CHECK_LOG_ROTATING_COUNT" env-default:"2" env-description:"Log file rotation count"`
}

var App Server

func Read(path string) {
	errEnv := godotenv.Load(".env")
	if errEnv != nil {
		fmt.Println("No .env file")
	}

	readError := cleanenv.ReadConfig(path+"config.yml", &App)
	if readError != nil {
		log.Panic("Failed to read yaml file", readError)
	}
}
