package rule

import (
	"testing"

	"bitbucket.verifone.com/validation-service/transaction"
)

func TestNewCardValidatorSuccess(t *testing.T) {
	var err error

	// Should create a new card validator where validation equals value
	_, err = newCardValidator(OperatorEqual, "1234")

	if err != nil {
		t.Error("unexpected error while creating new card validator", err.Error())
	}

	// Should create a new card validator where validation does not equal value
	_, err = newCardValidator(OperatorNotEqual, "1234")

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
	validator, err = newCardValidator(OperatorEqual, "1234")

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
		t.Error("expected to not match, card rules are skipped")
	}

	if ok := validator.Validate(transaction.Transaction{
		Amount:       1,
		CurrencyCode: transaction.USD,
		EntityId:     "1",
		Card:         "12345",
	}); ok {
		t.Error("expected to not match, card rules are skipped")
	}

	// Not equal
	validator, err = newCardValidator(OperatorNotEqual, "1234")

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
		t.Error("expected to not match, card rules are skipped")
	}

	if ok := validator.Validate(transaction.Transaction{
		Amount:       1,
		CurrencyCode: transaction.USD,
		EntityId:     "1",
		Card:         "12345",
	}); ok {
		t.Error("expected to not match, card rules are skipped")
	}
}
