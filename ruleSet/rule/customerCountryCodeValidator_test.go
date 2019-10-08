package rule

import (
	"bitbucket.verifone.com/validation-service/transaction"
	"testing"
)

func TestNewCustomerCountryCodeValidator(t *testing.T) {
	var err error

	// Should create a new customer country code validator where validation equals value
	_, err = newCustomerCountryCodeValidator(equal, "NL")

	if err != nil {
		t.Error("unexpected error while creating new customer country code validator", err.Error())
	}

	// Should create a new customer country code validator where validation does not equal value
	_, err = newCustomerCountryCodeValidator(notEqual, "NL")

	if err != nil {
		t.Error("unexpected error while creating new customer country code validator", err.Error())
	}

	// Should return an error when factory receives an invalid operator
	_, err = newCustomerCountryCodeValidator("foo", "NL")

	if err != InvalidOperatorError {
		t.Error("expected error while creating new customer country code validator with invalid operator")
	}

	// Should return an error when factory receives an invalid value
	_, err = newCustomerCountryCodeValidator(equal, "foo")

	if err != InvalidValueError {
		t.Error("expected error while creating new customer country code validator with invalid operator")
	}
}

func TestCustomerCountryCodeValidator_Validate(t *testing.T) {
	var validator Validator
	var err error

	// Equal
	validator, err = newCustomerCountryCodeValidator(equal, "NL")

	if err != nil {
		t.Error("unexpected error while creating new customer country code validator:", err.Error())
		return
	}

	if ok := validator.Validate(transaction.Transaction{
		CustomerCountryCode: transaction.NL,
	}); !ok {
		t.Error("expected validation to pass")
	}

	if ok := validator.Validate(transaction.Transaction{
		CustomerCountryCode: transaction.US,
	}); ok {
		t.Error("expected validation to fail")
	}

	// Not equal
	validator, err = newCustomerCountryCodeValidator(notEqual, "NL")

	if err != nil {
		t.Error("unexpected error while creating new customer country code validator:", err.Error())
		return
	}

	if ok := validator.Validate(transaction.Transaction{
		CustomerCountryCode: transaction.NL,
	}); ok {
		t.Error("expected validation to fail")
	}

	if ok := validator.Validate(transaction.Transaction{
		CustomerCountryCode: transaction.US,
	}); !ok {
		t.Error("expected validation to pass")
	}
}
