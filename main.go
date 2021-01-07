package main

import (
	"log"
	"os"
	"validation-service/cmd"
	"validation-service/config"

	"github.com/jessevdk/go-flags"
	"gopkg.in/yaml.v2"
)

var version = "unknown"
var appName = "Validation Service"

type ConfigFileOpts struct {
	ConfigFile string `long:"config" short:"f" default:"config.yml" description:"YAML configuration file path"`
}

func main() {
	config.AppName = appName
	config.Version = version

	var opts ConfigFileOpts

	_, err := flags.Parse(&opts)

	if err != nil {
		log.Printf("[ERROR] failed to parse arguments with error: %+v", err)
		os.Exit(1)
	}

	file, err := os.Open(opts.ConfigFile)

	if err != nil {
		log.Printf("[ERROR] failed to read configuration file with error: %+v", err)
		os.Exit(1)
	}

	defer file.Close()

	decoder := yaml.NewDecoder(file)

	var serverConfig config.Server

	if err := decoder.Decode(&serverConfig); err != nil {
		log.Printf("[ERROR] failed to decode configuration file with error: %+v", err)
		os.Exit(1)
	}

	if err := serverConfig.Validate(); err != nil {
		log.Printf("[ERROR] configuration file is invalid: %+v", err)
		os.Exit(1)
	}

	err = cmd.StartServer(serverConfig)

	if err != nil {
		log.Printf("[ERROR] server command failed with error: %+v", err)
		os.Exit(1)
	}
}
