package main

import (
	"bitbucket.verifone.com/validation-service/app"
	"bitbucket.verifone.com/validation-service/cmd"
	"bitbucket.verifone.com/validation-service/ruleSet"
	"bitbucket.verifone.com/validation-service/ruleSet/rule"
	"context"
	"fmt"
	"log"
	"os"
)

var version = "unknown"

func main() {
	log.Printf("Validation Service %s\n", version)

	ruleSetRepository, err := ruleSet.NewMongoRepository("mongo", "27017")

	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		return
	}

	rs, err := ruleSet.New("1", "Is greater than 5 and less than 5000", ruleSet.Block, []rule.Metadata{
		{
			Property: "amount",
			Operator: "<",
			Value:    "5000",
		},
		{
			Property: "amount",
			Operator: ">",
			Value:    "5",
		},
	})

	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		return
	}

	err = ruleSetRepository.Create(context.TODO(), rs)
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
