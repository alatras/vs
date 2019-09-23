package main

import (
	"bitbucket.verifone.com/validation-service/http"
	"errors"
	"fmt"
	"github.com/go-chi/chi"
	"os"
)

var version = "unknown"
var appName = "Validation Service"

const port = ":8080"

func main() {
	logger := createLogger(appName, version)

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

	logger.Output.Infof("Starting REST API server at port %s", port)

	err := http.NewServer(
		port,
		chi.NewRouter(),
		logger,
		createValidateTransactionApp(logger, mongoHost, mongoPort),
	).Start()

	if err != nil {
		logger.Error.WithError(err).Error("Failed to start REST API server")
		os.Exit(1)
	}
}
