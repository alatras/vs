package main

import (
	"bitbucket.verifone.com/validation-service/http"
	"github.com/go-chi/chi"
	"os"
)

var version = "unknown"
var appName = "Validation Service"

const port = ":8080"

func main() {
	logger := createLogger(appName, version)

	logger.Output.Infof("Starting REST API server at port %s", port)

	err := http.NewServer(
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
