package rule

import (
	"testing"

	"bitbucket.verifone.com/validation-service/transaction"
)

func TestNewCurrencyCodeValidator(t *testing.T) {
	var err error

	// Should create a new currency code validator where validation equals value
	_, err = newCurrencyCodeValidator(OperatorEqual, "EUR")

	if err != nil {
		t.Error("unexpected error while creating new currency code validator", err.Error())
	}

	// Should create a new currency code validator where validation does not equal value
	_, err = newCurrencyCodeValidator(OperatorNotEqual, "EUR")

	if err != nil {
		t.Error("unexpected error while creating new currency code validator", err.Error())
	}

	// Should return an error when factory receives an invalid operator
	_, err = newCurrencyCodeValidator("foo", "EUR")

	if err != InvalidOperatorError {
		t.Error("expected error while creating new currency code validator with invalid operator")
	}

	// Should return an error when factory receives an invalid value
	_, err = newCurrencyCodeValidator(OperatorEqual, "a")

	if err != InvalidValueError {
		t.Error("expected error while creating new currency code validator with invalid operator")
	}
}

func TestCurrencyCodeValidator_Validate(t *testing.T) {
	var validator Validator
	var err error

	// Equal
	validator, err = newCurrencyCodeValidator(OperatorEqual, "EUR")

	if err != nil {
		t.Error("unexpected error while creating new currency code validator:", err.Error())
		return
	}

	if ok := validator.Validate(transaction.Transaction{
		Amount:       1,
		CurrencyCode: transaction.EUR,
		EntityId:     "1",
	}); !ok {
		t.Error("expected validation to pass")
	}

	if ok := validator.Validate(transaction.Transaction{
		Amount:       1,
		CurrencyCode: transaction.USD,
		EntityId:     "1",
	}); ok {
		t.Error("expected validation to fail")
	}

	// Not equal
	validator, err = newCurrencyCodeValidator(OperatorNotEqual, "EUR")

	if err != nil {
		t.Error("unexpected error while creating new currency code validator:", err.Error())
		return
	}

	if ok := validator.Validate(transaction.Transaction{
		Amount:       1,
		CurrencyCode: transaction.EUR,
		EntityId:     "1",
	}); ok {
		t.Error("expected validation to fail")
	}

	if ok := validator.Validate(transaction.Transaction{
		Amount:       1,
		CurrencyCode: transaction.USD,
		EntityId:     "1",
	}); !ok {
		t.Error("expected validation to pass")
	}
}
