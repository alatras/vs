package main

import (
	"bitbucket.verifone.com/validation-service/app/validateTransaction"
	"bitbucket.verifone.com/validation-service/cmd"
	"bitbucket.verifone.com/validation-service/logger"
	"bitbucket.verifone.com/validation-service/ruleSet"
	"fmt"
	"github.com/jessevdk/go-flags"
	"log"
	"os"
	"runtime"
)

var version = "unknown"
var appName = "Validation Service"

func main() {
	var opts cmd.Options

	p := flags.NewParser(&opts, flags.Default)

	if _, err := p.Parse(); err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	}

	logger := setupLogger(opts.Log)

	server := setupServer(logger, opts)

	err := server.Start()

	if err != nil {
		logger.Error.WithError(err).Error("Failed to start REST API server")
		os.Exit(1)
	}
}

func setupLogger(logConfig cmd.LogGroup) *logger.Logger {
	logger, err := logger.NewLogger(
		appName,
		version,
		logConfig.FormatValue(),
		logConfig.LevelValue(),
	)

	if err != nil {
		log.Panic("Failed to initialize logger")
	}

	return logger
}

func setupServer(logger *logger.Logger, opts cmd.Options) *cmd.HttpServer {
	ruleSetRepository, err := ruleSet.NewMongoRepository(opts.Mongo.URL, opts.Mongo.DB)

	if err != nil {
		logger.Error.WithError(err).Error("Failed to initialize RuleSetRepository")
		os.Exit(1)
	}

	validatorService := validateTransaction.NewValidatorService(runtime.NumCPU(), ruleSetRepository, logger)

	serverAddress := fmt.Sprintf(":%d", opts.HTTPPort)

	logger.Output.Infof("Starting REST API server at port %d", opts.HTTPPort)

	return cmd.NewHttpServer(serverAddress, validatorService)
}
