package rule

import (
	"bitbucket.verifone.com/validation-service/transaction"
	"testing"
)

func TestNewCountryCodeValidator(t *testing.T) {
	var err error

	// Should create a new country code validator where validation equals value
	_, err = newCurrencyCodeValidator(equal, "EUR")

	if err != nil {
		t.Error("unexpected error while creating new country code validator", err.Error())
	}

	// Should create a new country code validator where validation does not equal value
	_, err = newCurrencyCodeValidator(notEqual, "EUR")

	if err != nil {
		t.Error("unexpected error while creating new country code validator", err.Error())
	}

	// Should return an error when factory receives an invalid operator
	_, err = newCurrencyCodeValidator("foo", "EUR")

	if err != InvalidOperatorError {
		t.Error("expected error while creating new amount validator with invalid operator")
	}

	// Should return an error when factory receives an invalid value
	_, err = newCurrencyCodeValidator(equal, "a")

	if err != InvalidValueError {
		t.Error("expected error while creating new amount validator with invalid operator")
	}
}

func TestCountryCodeValidator_Validate(t *testing.T) {
	var validator Validator
	var err error

	// Equal
	validator, err = newCurrencyCodeValidator(equal, "EUR")

	if err != nil {
		t.Error("unexpected error while creating new country code validator:", err.Error())
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
	validator, err = newCurrencyCodeValidator(notEqual, "EUR")

	if err != nil {
		t.Error("unexpected error while creating new country code validator:", err.Error())
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
