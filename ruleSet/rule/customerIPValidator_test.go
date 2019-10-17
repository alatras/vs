package rule

import (
	"bitbucket.verifone.com/validation-service/transaction"
	"testing"
)

func TestNewCustomerIPValidatorSuccess(t *testing.T) {
	var err error

	// Should create a new customer ip validator where validation equals value
	_, err = newCustomerIPValidator(equal, "11.11.11.11")

	if err != nil {
		t.Error("unexpected error while creating new customer ip validator", err.Error())
	}

	// Should create a new customer ip validator where validation does not equal value
	_, err = newCustomerIPValidator(notEqual, "11.11.11.11")

	if err != nil {
		t.Error("unexpected error while creating new customer ip validator", err.Error())
	}
}

func TestNewCustomerIPValidatorFailure(t *testing.T) {
	// Should return an error when factory receives an invalid operator
	_, err := newCustomerIPValidator("foo", "11.11.11.11")

	if err != InvalidOperatorError {
		t.Error("expected error while creating new customer ip validator with invalid operator")
	}

	// Should return an error when factory receives an invalid value
	_, err = newCustomerIPValidator(equal, "a")

	if err != InvalidValueError {
		t.Error("expected error while creating new customer ip validator with invalid operator")
	}
}

func TestCustomerIPValidator_Validate(t *testing.T) {
	var validator Validator
	var err error

	// Equal
	validator, err = newCustomerIPValidator(equal, "11.11.11.11")

	if err != nil {
		t.Error("unexpected error while creating new customer ip validator:", err.Error())
		return
	}

	if ok := validator.Validate(transaction.Transaction{
		Amount:       1,
		CurrencyCode: transaction.EUR,
		EntityId:     "1",
		CustomerIP:   "11.11.11.11",
	}); !ok {
		t.Error("expected validation to pass")
	}

	if ok := validator.Validate(transaction.Transaction{
		Amount:       1,
		CurrencyCode: transaction.USD,
		EntityId:     "1",
		CustomerIP:   "11.11.11.12",
	}); ok {
		t.Error("expected validation to fail")
	}

	// Not equal
	validator, err = newCustomerIPValidator(notEqual, "11.11.11.11")

	if err != nil {
		t.Error("unexpected error while creating new customer ip validator:", err.Error())
		return
	}

	if ok := validator.Validate(transaction.Transaction{
		Amount:       1,
		CurrencyCode: transaction.EUR,
		EntityId:     "1",
		CustomerIP:   "11.11.11.11",
	}); ok {
		t.Error("expected validation to fail")
	}

	if ok := validator.Validate(transaction.Transaction{
		Amount:       1,
		CurrencyCode: transaction.USD,
		EntityId:     "1",
		CustomerIP:   "11.11.11.12",
	}); !ok {
		t.Error("expected validation to pass")
	}
}
