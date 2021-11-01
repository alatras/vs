package rule

import (
	"testing"
	"validation-service/transaction"
)

func TestThreeDSecureEnrollmentStatusSuccess(t *testing.T) {
	var err error

	// Should create a new customer id validator where validation equals value
	_, err = newThreeDSecureEnrollmentStatusValidator(OperatorEqual, "Y")

	if err != nil {
		t.Error("unexpected error while creating new 3D secure enrollment status validator with equal operator", err.Error())
	}

	// Should create a new customer id validator where validation does not equal value
	_, err = newThreeDSecureEnrollmentStatusValidator(OperatorNotEqual, "Y")

	if err != nil {
		t.Error("unexpected error while creating new 3D secure enrollment status validator with not-equal operator", err.Error())
	}
}

func TestThreeDSecureEnrollmentStatusValidatorFailure(t *testing.T) {

	// Should return an error when factory receives an invalid operator
	_, err := newThreeDSecureEnrollmentStatusValidator("foo", "Y")

	if err != InvalidOperatorError {
		t.Error("unexpected error while creating new 3D secure enrollment status validator with invalid operator")
	}
}

func TestThreeDSecureEnrollmentStatus_Validate(t *testing.T) {
	var validator Validator
	var err error

	// Equal
	validator, err = newThreeDSecureEnrollmentStatusValidator(OperatorEqual, "Y")

	if err != nil {
		t.Error("unexpected error while creating new 3D secure enrollment status validator for validation test:", err.Error())
		return
	}

	if ok := validator.Validate(transaction.Transaction{
		Amount:                       1,
		CurrencyCode:                 transaction.EUR,
		EntityId:                     "1",
		CustomerId:                   "1234",
		ThreeDSecureEnrollmentStatus: "Y",
	}); !ok {
		t.Error("expected validation to pass")
	}

	if ok := validator.Validate(transaction.Transaction{
		Amount:                       1,
		CurrencyCode:                 transaction.USD,
		EntityId:                     "1",
		CustomerId:                   "12345",
		ThreeDSecureEnrollmentStatus: "U",
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
		Amount:                       1,
		CurrencyCode:                 transaction.EUR,
		EntityId:                     "1",
		CustomerId:                   "1234",
		ThreeDSecureEnrollmentStatus: "Y",
	}); ok {
		t.Error("expected validation to fail")
	}

	if ok := validator.Validate(transaction.Transaction{
		Amount:                       1,
		CurrencyCode:                 transaction.USD,
		EntityId:                     "1",
		CustomerId:                   "12345",
		ThreeDSecureEnrollmentStatus: "U",
	}); !ok {
		t.Error("expected validation to pass")
	}
}
