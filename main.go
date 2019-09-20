package main

import (
	"bitbucket.verifone.com/validation-service/app"
	"bitbucket.verifone.com/validation-service/cmd"
	"bitbucket.verifone.com/validation-service/ruleSet"
	"errors"
	"fmt"
	"log"
	"os"
	"runtime"
)

var version = "unknown"

func main() {
	log.Printf("Validation Service %s\n", version)

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
		_, _ = fmt.Fprintln(os.Stderr, err)
		return
	}

	validatorService := app.NewValidatorService(runtime.NumCPU(), ruleSetRepository)

	err = cmd.NewHttpServer(":8080", validatorService).Start()

	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
	}
}
