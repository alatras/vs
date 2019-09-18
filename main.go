package main

import (
	"bitbucket.verifone.com/validation-service/app"
	"bitbucket.verifone.com/validation-service/cmd"
	"bitbucket.verifone.com/validation-service/ruleSet"
	"fmt"
	"log"
	"os"
)

var version = "unknown"

func main() {
	log.Printf("Validation Service %s\n", version)

	ruleSetRepository, err := ruleSet.NewStubRepository()

	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		return
	}

	validatorService := app.NewValidatorService(6, ruleSetRepository)

	err = cmd.NewHttpServer(":8080", validatorService).Start()

	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
	}
}
