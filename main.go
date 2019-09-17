package main

import (
	"bitbucket.verifone.com/validation-service/app"
	"bitbucket.verifone.com/validation-service/cmd"
	"bitbucket.verifone.com/validation-service/infra/instrumentation/main"
	"bitbucket.verifone.com/validation-service/infra/logger"
	"bitbucket.verifone.com/validation-service/infra/repository"
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

	instrumentation := instrumentation.NewMainInstrumentation(logger)

	ruleSetRepository, err := repository.NewStubRuleSetRepository()

	if err != nil {
		instrumentation.FailedToInitRuleSetRepository(err)
		os.Exit(1)
	}

	validatorService := app.NewValidatorService(6, ruleSetRepository)

	serverPort := 8080
	serverAddress := fmt.Sprintf(":%d", serverPort)

	instrumentation.StartingRestApiServer(serverPort)

	err = cmd.NewHttpServer(serverAddress, validatorService).Start()

	if err != nil {
		instrumentation.FailedToStartRestApiServer(err)
		os.Exit(1)
	}
}
