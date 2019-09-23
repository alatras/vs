package main

import (
	"bitbucket.verifone.com/validation-service/app/validateTransaction"
	"bitbucket.verifone.com/validation-service/http"
	"bitbucket.verifone.com/validation-service/logger"
	"bitbucket.verifone.com/validation-service/ruleSet"
	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
	"log"
	"os"
)

var version = "unknown"
var appName = "Validation Service"

const port = ":8080"

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

	logger.Output.Infof("Starting REST API server at port %s", port)

	err = http.NewServer(
		port,
		chi.NewRouter(),
		logger,
		createValidateTransactionApp(logger),
	).Start()

	if err != nil {
		logger.Error.WithError(err).Error("Failed to start REST API server")
		os.Exit(1)
	}
}

func createValidateTransactionApp(logger *logger.Logger) *validateTransaction.ValidatorService {
	ruleSetRepository, err := ruleSet.NewStubRuleSetRepository()

	if err != nil {
		logger.Error.WithError(err).Error("Failed to initialize RuleSetRepository")
		os.Exit(1)
	}

	validator := validateTransaction.NewValidatorService(6, ruleSetRepository, logger)
	return &validator
}
