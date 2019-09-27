package rule

import (
	"bitbucket.verifone.com/validation-service/transaction"
	"testing"
)

func TestNewCountryCodeValidator(t *testing.T) {
	var err error

	// Should create a new country code validator where validation equals value
	_, err = newCountryCodeValidator(equal, "NL")

	if err != nil {
		t.Error("unexpected error while creating new country code validator", err.Error())
	}

	// Should create a new country code validator where validation does not equal value
	_, err = newCountryCodeValidator(notEqual, "NL")

	if err != nil {
		t.Error("unexpected error while creating new country code validator", err.Error())
	}
}

func TestCountryCodeValidator_Validate(t *testing.T) {
	var validator Validator
	var err error

	// Equal
	validator, err = newCountryCodeValidator(equal, "NL")

	if err != nil {
		t.Error("unexpected error while creating new country code validator:", err.Error())
		return
	}

	if ok := validator.Validate(transaction.Transaction{
		Amount:      1,
		CountryCode: "NL",
		EntityId:    "1",
	}); !ok {
		t.Error("expected validation to pass")
	}

	if ok := validator.Validate(transaction.Transaction{
		Amount:      1,
		CountryCode: "AB",
		EntityId:    "1",
	}); ok {
		t.Error("expected validation to fail")
	}

	// Not equal
	validator, err = newCountryCodeValidator(notEqual, "NL")

	if err != nil {
		t.Error("unexpected error while creating new country code validator:", err.Error())
		return
	}

	if ok := validator.Validate(transaction.Transaction{
		Amount:      1,
		CountryCode: "NL",
		EntityId:    "1",
	}); ok {
		t.Error("expected validation to fail")
	}

	if ok := validator.Validate(transaction.Transaction{
		Amount:      1,
		CountryCode: "AB",
		EntityId:    "1",
	}); !ok {
		t.Error("expected validation to pass")
	}
}
