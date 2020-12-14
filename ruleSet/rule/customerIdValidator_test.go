package rule

import (
	"testing"

	"bitbucket.verifone.com/validation-service/transaction"
)

func TestNewCustomerIdValidatorSuccess(t *testing.T) {
	var err error

	// Should create a new customer id validator where validation equals value
	_, err = newCustomerIdValidator(OperatorEqual, "1234")

	if err != nil {
		t.Error("unexpected error while creating new customer id validator", err.Error())
	}

	// Should create a new customer id validator where validation does not equal value
	_, err = newCustomerIdValidator(OperatorNotEqual, "1234")

	if err != nil {
		t.Error("unexpected error while creating new customer id validator", err.Error())
	}
}

func TestNewCustomerIdValidatorFailure(t *testing.T) {

	// Should return an error when factory receives an invalid operator
	_, err := newCustomerIdValidator("foo", "1234")

	if err != InvalidOperatorError {
		t.Error("expected error while creating new customer id validator with invalid operator")
	}
}

func TestCustomerIdValidator_Validate(t *testing.T) {
	var validator Validator
	var err error

	// Equal
	validator, err = newCustomerIdValidator(OperatorEqual, "1234")

	if err != nil {
		t.Error("unexpected error while creating new customer id validator:", err.Error())
		return
	}

	if ok := validator.Validate(transaction.Transaction{
		Amount:       1,
		CurrencyCode: transaction.EUR,
		EntityId:     "1",
		CustomerId:   "1234",
	}); !ok {
		t.Error("expected validation to pass")
	}

	if ok := validator.Validate(transaction.Transaction{
		Amount:       1,
		CurrencyCode: transaction.USD,
		EntityId:     "1",
		CustomerId:   "12345",
	}); ok {
		t.Error("expected validation to fail")
	}

	// Not equal
	validator, err = newCustomerIdValidator(OperatorNotEqual, "1234")

	if err != nil {
		t.Error("unexpected error while creating new customer id validator:", err.Error())
		return
	}

	if ok := validator.Validate(transaction.Transaction{
		Amount:       1,
		CurrencyCode: transaction.EUR,
		EntityId:     "1",
		CustomerId:   "1234",
	}); ok {
		t.Error("expected validation to fail")
	}

	if ok := validator.Validate(transaction.Transaction{
		Amount:       1,
		CurrencyCode: transaction.USD,
		EntityId:     "1",
		CustomerId:   "12345",
	}); !ok {
		t.Error("expected validation to pass")
	}
}
