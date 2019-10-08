package rule

import (
	"bitbucket.verifone.com/validation-service/transaction"
	"testing"
)

func TestNewCustomerIPCountryCodeValidator(t *testing.T) {
	var err error

	// Should create a new customer ip country code validator where validation equals value
	_, err = newCustomerIPCountryCodeValidator(equal, "NL")

	if err != nil {
		t.Error("unexpected error while creating new customer ip country code validator", err.Error())
	}

	// Should create a new customer ip country code validator where validation does not equal value
	_, err = newCustomerIPCountryCodeValidator(notEqual, "NL")

	if err != nil {
		t.Error("unexpected error while creating new customer ip country code validator", err.Error())
	}

	// Should return an error when factory receives an invalid operator
	_, err = newCustomerIPCountryCodeValidator("foo", "NL")

	if err != InvalidOperatorError {
		t.Error("expected error while creating new customer ip country code validator with invalid operator")
	}

	// Should return an error when factory receives an invalid value
	_, err = newCustomerIPCountryCodeValidator(equal, "a")

	if err != InvalidValueError {
		t.Error("expected error while creating new customer ip country code validator with invalid operator")
	}
}

func TestCustomerIPCountryCodeValidator_Validate(t *testing.T) {
	var validator Validator
	var err error

	// Equal
	validator, err = newCustomerIPCountryCodeValidator(equal, "NL")

	if err != nil {
		t.Error("unexpected error while creating new customer ip country code validator:", err.Error())
		return
	}

	if ok := validator.Validate(transaction.Transaction{
		Amount:            1,
		CurrencyCode:      transaction.EUR,
		EntityId:          "1",
		CustomerIPCountry: "NL",
	}); !ok {
		t.Error("expected validation to pass")
	}

	if ok := validator.Validate(transaction.Transaction{
		Amount:            1,
		CurrencyCode:      transaction.USD,
		EntityId:          "1",
		CustomerIPCountry: "SD",
	}); ok {
		t.Error("expected validation to fail")
	}

	// Not equal
	validator, err = newCustomerIPCountryCodeValidator(notEqual, "NL")

	if err != nil {
		t.Error("unexpected error while creating new customer ip country code validator:", err.Error())
		return
	}

	if ok := validator.Validate(transaction.Transaction{
		Amount:            1,
		CurrencyCode:      transaction.EUR,
		EntityId:          "1",
		CustomerIPCountry: "NL",
	}); ok {
		t.Error("expected validation to fail")
	}

	if ok := validator.Validate(transaction.Transaction{
		Amount:            1,
		CurrencyCode:      transaction.USD,
		EntityId:          "1",
		CustomerIPCountry: "SD",
	}); !ok {
		t.Error("expected validation to pass")
	}
}
