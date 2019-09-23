package main

import (
	"bitbucket.verifone.com/validation-service/app/validateTransaction"
	"bitbucket.verifone.com/validation-service/logger"
	"bitbucket.verifone.com/validation-service/ruleSet"
	"github.com/sirupsen/logrus"
	"log"
	"os"
)

func createLogger(app string, version string) *logger.Logger {
	l, err := logger.NewLogger(
		appName,
		version,
		logger.TextFormat,
		logrus.TraceLevel,
	)
	if err != nil {
		log.Panic("Failed to initialize logger")
	}
	return l
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
