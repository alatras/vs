package validateTransaction

import (
	"context"
	"errors"
	"testing"

	"validation-service/logger"
	"validation-service/report"
	"validation-service/ruleSet"
	"validation-service/ruleSet/rule"
	"validation-service/transaction"

	"github.com/google/uuid"
)

var trx = transaction.Transaction{
	EntityId:            uuid.New().String(),
	Amount:              100,
	CurrencyCode:        "EUR",
	CustomerCountryCode: "NL",
	IssuerCountryCode:   "NLD",
	Card:                "123123123123",
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

func Test_App_ValidateTransaction_Success_SkipCardValidation(t *testing.T) {
	log := logger.NewStubLogger()
	repo, _ := ruleSet.NewStubRepository(nil)

	err := repo.Create(context.TODO(), ruleSet.RuleSet{
		Id:       "1234",
		EntityId: trx.EntityId,
		Action:   ruleSet.Block,
		Name:     "block ruleset",
		RuleMetadata: []rule.Metadata{
			{
				Property: rule.PropertyCard,
				Operator: rule.OperatorEqual,
				Value:    trx.Card,
			},
			{
				Property: rule.PropertyCurrencyCode,
				Operator: rule.OperatorEqual,
				Value:    string(trx.CurrencyCode),
			},
		},
	})

	if err != nil {
		t.Errorf("expected create blocking ruleset, but got error %s", err)
	}

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
