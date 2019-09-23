package main

import (
	"bitbucket.verifone.com/validation-service/http"
	"bitbucket.verifone.com/validation-service/logger"
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

	logger.Output.Infof("Starting REST API server at port %d", port)

	err = http.NewServer(port, chi.NewRouter(), logger).Start()

	if err != nil {
		logger.Error.WithError(err).Error("Failed to start REST API server")
		os.Exit(1)
	}
}
