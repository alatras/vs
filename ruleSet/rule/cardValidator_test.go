package rule

import (
	"bitbucket.verifone.com/validation-service/transaction"
	"testing"
)

func TestNewCardValidatorSuccess(t *testing.T) {
	var err error

	// Should create a new card validator where validation equals value
	_, err = newCardValidator(equal, "1234")

	if err != nil {
		t.Error("unexpected error while creating new card validator", err.Error())
	}

	// Should create a new card validator where validation does not equal value
	_, err = newCardValidator(notEqual, "1234")

	if err != nil {
		t.Error("unexpected error while creating new card validator", err.Error())
	}
}

func TestNewCardValidatorFailure(t *testing.T) {
	// Should return an error when factory receives an invalid operator
	_, err := newCardValidator("foo", "1234")

	if err != InvalidOperatorError {
		t.Error("expected error while creating new card validator with invalid operator")
	}
}

func TestCardValidator_Validate(t *testing.T) {
	var validator Validator
	var err error

	// Equal
	validator, err = newCardValidator(equal, "1234")

	if err != nil {
		t.Error("unexpected error while creating new card validator:", err.Error())
		return
	}

	if ok := validator.Validate(transaction.Transaction{
		Amount:       1,
		CurrencyCode: transaction.EUR,
		EntityId:     "1",
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
	validator, err = newCardValidator(notEqual, "1234")

	if err != nil {
		t.Error("unexpected error while creating new card validator:", err.Error())
		return
	}

	if ok := validator.Validate(transaction.Transaction{
		Amount:       1,
		CurrencyCode: transaction.EUR,
		EntityId:     "1",
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
