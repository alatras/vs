package rule

import (
	"testing"

	"bitbucket.verifone.com/validation-service/transaction"
)

func TestNewIssuerCountryCodeValidatorSuccess(t *testing.T) {
	var err error

	// Should create a new issuer country code validator where validation equals value
	_, err = newIssuerCountryCodeValidator(OperatorEqual, "NLD")

	if err != nil {
		t.Error("unexpected error while creating new issuer country code validator", err.Error())
	}

	// Should create a new issuer country code validator where validation does not equal value
	_, err = newIssuerCountryCodeValidator(OperatorNotEqual, "NLD")

	if err != nil {
		t.Error("unexpected error while creating new issuer country code validator", err.Error())
	}
}

func TestNewIssuerCountryCodeValidatorFailure(t *testing.T) {
	// Should return an error when factory receives an invalid operator
	_, err := newIssuerCountryCodeValidator("foo", "NLD")

	if err != InvalidOperatorError {
		t.Error("expected error while creating new issuer country code validator with invalid operator")
	}

	// Should return an error when factory receives an invalid value
	_, err = newIssuerCountryCodeValidator(OperatorEqual, "foo")

	if err != InvalidValueError {
		t.Error("expected error while creating new issuer country code validator with invalid operator")
	}
}

func TestIssuerCountryCodeValidator_Validate(t *testing.T) {
	var validator Validator
	var err error

	// Equal
	validator, err = newIssuerCountryCodeValidator(OperatorEqual, "NLD")

	if err != nil {
		t.Error("unexpected error while creating new issuer country code validator:", err.Error())
		return
	}

	if ok := validator.Validate(transaction.Transaction{
		IssuerCountryCode: transaction.NLD,
	}); !ok {
		t.Error("expected validation to pass")
	}

	if ok := validator.Validate(transaction.Transaction{
		IssuerCountryCode: transaction.USA,
	}); ok {
		t.Error("expected validation to fail")
	}

	// Not equal
	validator, err = newIssuerCountryCodeValidator(OperatorNotEqual, "NLD")

	if err != nil {
		t.Error("unexpected error while creating new issuer country code validator:", err.Error())
		return
	}

	if ok := validator.Validate(transaction.Transaction{
		IssuerCountryCode: transaction.NLD,
	}); ok {
		t.Error("expected validation to fail")
	}

	if ok := validator.Validate(transaction.Transaction{
		IssuerCountryCode: transaction.USA,
	}); !ok {
		t.Error("expected validation to pass")
	}
}
