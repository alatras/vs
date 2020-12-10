package rule

import (
	"testing"

	"bitbucket.verifone.com/validation-service/transaction"
)

func TestNewEntityIdValidatorSuccess(t *testing.T) {
	var err error

	// Should create a new entity validator where validation equals value
	_, err = newEntityIdValidator(OperatorEqual, "1234")

	if err != nil {
		t.Error("unexpected error while creating new entity validator", err.Error())
	}

	// Should create a new entity validator where validation does not equal value
	_, err = newEntityIdValidator(OperatorNotEqual, "1234")

	if err != nil {
		t.Error("unexpected error while creating new entity validator", err.Error())
	}
}

func TestNewEntityIdValidatorFailure(t *testing.T) {
	// Should return an error when factory receives an invalid operator
	_, err := newEntityIdValidator("foo", "1234")

	if err != InvalidOperatorError {
		t.Error("expected error while creating new entity validator with invalid operator")
	}
}

func TestEntityIdValidator_Validate(t *testing.T) {
	var validator Validator
	var err error

	// Equal
	validator, err = newEntityIdValidator(OperatorEqual, "1234")

	if err != nil {
		t.Error("unexpected error while creating new entity validator:", err.Error())
		return
	}

	if ok := validator.Validate(transaction.Transaction{
		Amount:       1,
		CurrencyCode: transaction.EUR,
		EntityId:     "1234",
		Card:         "1234",
	}); !ok {
		t.Error("expected validation to pass")
	}

	if ok := validator.Validate(transaction.Transaction{
		Amount:       1,
		CurrencyCode: transaction.USD,
		EntityId:     "1",
		Card:         "12345",
	}); ok {
		t.Error("expected validation to fail")
	}

	// Not equal
	validator, err = newEntityIdValidator(OperatorNotEqual, "1234")

	if err != nil {
		t.Error("unexpected error while creating new entity validator:", err.Error())
		return
	}

	if ok := validator.Validate(transaction.Transaction{
		Amount:       1,
		CurrencyCode: transaction.EUR,
		EntityId:     "1234",
		Card:         "1234",
	}); ok {
		t.Error("expected validation to fail")
	}

	if ok := validator.Validate(transaction.Transaction{
		Amount:       1,
		CurrencyCode: transaction.USD,
		EntityId:     "1",
		Card:         "12345",
	}); !ok {
		t.Error("expected validation to pass")
	}
}
