package validateTransaction

import (
	"bitbucket.verifone.com/validation-service/logger"
	"bitbucket.verifone.com/validation-service/report"
	"bitbucket.verifone.com/validation-service/ruleSet"
	"bitbucket.verifone.com/validation-service/transaction"
	"context"
	"errors"
	"github.com/google/uuid"
	"testing"
)

var trx = transaction.Transaction{
	EntityId:            uuid.New().String(),
	Amount:              100,
	CurrencyCode:        "EUR",
	CustomerCountryCode: "NL",
	IssuerCountryCode:   "NLD",
}

func Test_App_ValidateTransaction_Success(t *testing.T) {
	log := logger.NewStubLogger()
	repo, _ := ruleSet.NewStubRepository(nil)

	app := NewValidatorService(1, repo, log)

	reportChan, errorChan := app.Enqueue(context.TODO(), trx)

	select {
	case rep := <-reportChan:
		if rep.Action != report.Pass {
			t.Errorf("expected action PASS, but got %s", rep.Action)
		}
	case validationError := <-errorChan:
		t.Error(validationError)
	}
}

func Test_App_ValidateTransaction_UnexpectedError(t *testing.T) {
	log := logger.NewStubLogger()
	repo, _ := ruleSet.NewStubRepository(errors.New("unexpected"))

	app := NewValidatorService(1, repo, log)

	reportChan, errorChan := app.Enqueue(context.TODO(), trx)

	select {
	case <-reportChan:
		t.Error("expected validation to fail but it succeeded")
	case validationError := <-errorChan:
		if !validationError.Is(UnexpectedErr) {
			t.Errorf("expected error to be unexpected but got %v", validationError)
		}
	}
}
