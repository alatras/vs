package rule

import (
	"testing"
	"validation-service/transaction"
)

func TestThreeDSecureErrorNoValidatorSuccess(t *testing.T) {
	var err error

	// Should create a new customer id validator where validation equals value
	_, err = newThreeDSecureErrorNoValidator(OperatorEqual, "1")

	if err != nil {
		t.Error("unexpected error while creating new 3D secure error no validator with equal operator", err.Error())
	}

	// Should create a new customer id validator where validation does not equal value
	_, err = newThreeDSecureErrorNoValidator(OperatorNotEqual, "1")

	if err != nil {
		t.Error("unexpected error while creating new 3D secure error no validator with not-equal operator", err.Error())
	}
}

func TestNewThreeDSecureErrorNoValidatorFailure(t *testing.T) {

	// Should return an error when factory receives an invalid operator
	_, err := newThreeDSecureErrorNoValidator("foo", "1")

	if err != InvalidOperatorError {
		t.Error("unexpected error while creating new 3D secure error no validator with invalid operator")
	}
}

func TestThreeDSecureErrorNoValidator_Validate(t *testing.T) {
	var validator Validator
	var err error

	// Equal
	validator, err = newThreeDSecureErrorNoValidator(OperatorEqual, "1")

	if err != nil {
		t.Error("unexpected error while creating new 3D secure error no validator for validation test:", err.Error())
		return
	}

	if ok := validator.Validate(transaction.Transaction{
		Amount:              1,
		CurrencyCode:        transaction.EUR,
		EntityId:            "1",
		CustomerId:          "1234",
		ThreeDSecureErrorNo: "1",
	}); !ok {
		t.Error("expected validation to pass")
	}

	if ok := validator.Validate(transaction.Transaction{
		Amount:              1,
		CurrencyCode:        transaction.USD,
		EntityId:            "1",
		CustomerId:          "12345",
		ThreeDSecureErrorNo: "0",
	}); ok {
		t.Error("expected validation to fail")
	}

	// Not equal
	validator, err = newThreeDSecureErrorNoValidator(OperatorNotEqual, "1")

	if err != nil {
		t.Error("unexpected error while creating new customer id validator:", err.Error())
		return
	}

	if ok := validator.Validate(transaction.Transaction{
		Amount:              1,
		CurrencyCode:        transaction.EUR,
		EntityId:            "1",
		CustomerId:          "1234",
		ThreeDSecureErrorNo: "1",
	}); ok {
		t.Error("expected validation to fail")
	}

	if ok := validator.Validate(transaction.Transaction{
		Amount:              1,
		CurrencyCode:        transaction.USD,
		EntityId:            "1",
		CustomerId:          "12345",
		ThreeDSecureErrorNo: "0",
	}); !ok {
		t.Error("expected validation to pass")
	}
}
