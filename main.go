package main

import (
	"bitbucket.verifone.com/validation-service/app/validateTransaction"
	"bitbucket.verifone.com/validation-service/cmd"
	"bitbucket.verifone.com/validation-service/logger"
	"bitbucket.verifone.com/validation-service/ruleSet"
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
	"os"
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

	ruleSetRepository, err := ruleSet.NewStubRuleSetRepository()

	if err != nil {
		logger.Error.WithError(err).Error("Failed to initialize RuleSetRepository")
		os.Exit(1)
	}

	validatorService := validateTransaction.NewValidatorService(6, ruleSetRepository, logger)

	serverPort := 8080
	serverAddress := fmt.Sprintf(":%d", serverPort)

	logger.Output.Infof("Starting REST API server at port %d", serverPort)

	err = cmd.NewHttpServer(serverAddress, validatorService).Start()

	if err != nil {
		logger.Error.WithError(err).Error("Failed to start REST API server")
		os.Exit(1)
	}
}
