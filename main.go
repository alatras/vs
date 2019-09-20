package main

import (
	"bitbucket.verifone.com/validation-service/app/validateTransaction"
	"bitbucket.verifone.com/validation-service/cmd"
	"bitbucket.verifone.com/validation-service/logger"
	"bitbucket.verifone.com/validation-service/ruleSet"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"runtime"
)

var version = "unknown"
var appName = "Validation Service"

func main() {
	logger, err := logger.NewLogger(
		appName,
		version,
		logger.TextFormat,
		logrus.TraceLevel,
	)

	if err != nil {
		log.Panic("Failed to initialize logger")
	}

	mongoHost := os.Getenv("MONGO_HOST")
	if mongoHost == "" {
		_, _ = fmt.Fprintln(os.Stderr, errors.New("mongo host undefined"))
		return
	}

	mongoPort := os.Getenv("MONGO_PORT")
	if mongoPort == "" {
		_, _ = fmt.Fprintln(os.Stderr, errors.New("mongo port undefined"))
		return
	}

	ruleSetRepository, err := ruleSet.NewMongoRepository(mongoHost, mongoPort)
	if err != nil {
		logger.Error.WithError(err).Error("Failed to initialize RuleSetRepository")
		os.Exit(1)
	}

	validatorService := validateTransaction.NewValidatorService(runtime.NumCPU(), ruleSetRepository, logger)

	serverPort := 8080
	serverAddress := fmt.Sprintf(":%d", serverPort)

	logger.Output.Infof("Starting REST API server at port %d", serverPort)

	err = cmd.NewHttpServer(serverAddress, validatorService).Start()

	if err != nil {
		logger.Error.WithError(err).Error("Failed to start REST API server")
		os.Exit(1)
	}
}
