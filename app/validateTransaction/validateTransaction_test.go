package validateTransaction

import (
	"bitbucket.verifone.com/validation-service/entityService"
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

type stubEntityService struct {
	err error
}

func (s *stubEntityService) Ping() error {
	return s.err
}

func (s *stubEntityService) GetAncestorsOf(entityId string) ([]string, error) {
	if s.err != nil {
		return []string{}, s.err
	}

	return []string{entityId}, nil
}

func (s *stubEntityService) GetDescendantsOf(entityId string) ([]string, error) {
	if s.err != nil {
		return []string{}, s.err
	}

	return []string{entityId}, nil
}

func Test_App_ValidateTransaction_Success(t *testing.T) {
	log := logger.NewStubLogger()
	repo, _ := ruleSet.NewStubRepository()
	entityService := stubEntityService{}

	app := NewValidatorService(1, &entityService, repo, log)

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

func Test_App_ValidateTransaction_EntityNotFound(t *testing.T) {
	log := logger.NewStubLogger()
	repo, _ := ruleSet.NewStubRepository()
	entityService := stubEntityService{
		err: entityService.EntityNotFound,
	}

	app := NewValidatorService(1, &entityService, repo, log)

	reportChan, errorChan := app.Enqueue(context.TODO(), trx)

	select {
	case <-reportChan:
		t.Error("expected validation to fail but it succeeded")
	case validationError := <-errorChan:
		if !validationError.Is(EntityIdNotFoundErr) {
			t.Errorf("expected error to be entity not found but got %v", validationError)
		}
	}
}

func Test_App_ValidateTransaction_UnexpectedError(t *testing.T) {
	log := logger.NewStubLogger()
	repo, _ := ruleSet.NewStubRepository()
	entityService := stubEntityService{
		err: errors.New("unexpected"),
	}

	app := NewValidatorService(1, &entityService, repo, log)

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
