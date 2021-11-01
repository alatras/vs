package rule

import (
	"testing"
	"validation-service/transaction"
)

func TestThreeDSecureSignatureVerificationValidatorSuccess(t *testing.T) {
	var err error

	// Should create a new customer id validator where validation equals value
	_, err = newThreeDSecureSignatureVerificationValidator(OperatorEqual, "Y")

	if err != nil {
		t.Error("unexpected error while creating new 3D secure signature validation validator with equal operator", err.Error())
	}

	// Should create a new customer id validator where validation does not equal value
	_, err = newThreeDSecureSignatureVerificationValidator(OperatorNotEqual, "Y")

	if err != nil {
		t.Error("unexpected error while creating new 3D secure signature validation validator with not-equal operator", err.Error())
	}
}

func TestNewThreeDSecureSignatureVerificationValidatorFailure(t *testing.T) {

	// Should return an error when factory receives an invalid operator
	_, err := newThreeDSecureSignatureVerificationValidator("foo", "Y")

	if err != InvalidOperatorError {
		t.Error("unexpected error while creating new 3D secure signature validation validator with invalid operator")
	}
}

func TestThreeDSecureSignatureVerificationValidator_Validate(t *testing.T) {
	var validator Validator
	var err error

	// Equal
	validator, err = newThreeDSecureSignatureVerificationValidator(OperatorEqual, "Y")

	if err != nil {
		t.Error("unexpected error while creating new 3D secure signature validation validator for validation test:", err.Error())
		return
	}

	if ok := validator.Validate(transaction.Transaction{
		Amount:                            1,
		CurrencyCode:                      transaction.EUR,
		EntityId:                          "Y",
		CustomerId:                        "1234",
		ThreeDSecureSignatureVerification: "Y",
	}); !ok {
		t.Error("expected validation to pass")
	}

	if ok := validator.Validate(transaction.Transaction{
		Amount:                            1,
		CurrencyCode:                      transaction.USD,
		EntityId:                          "Y",
		CustomerId:                        "12345",
		ThreeDSecureSignatureVerification: "N",
	}); ok {
		t.Error("expected validation to fail")
	}

	// Not equal
	validator, err = newThreeDSecureSignatureVerificationValidator(OperatorNotEqual, "Y")

	if err != nil {
		t.Error("unexpected error while creating new customer id validator:", err.Error())
		return
	}

	if ok := validator.Validate(transaction.Transaction{
		Amount:                            1,
		CurrencyCode:                      transaction.EUR,
		EntityId:                          "Y",
		CustomerId:                        "1234",
		ThreeDSecureSignatureVerification: "Y",
	}); ok {
		t.Error("expected validation to fail")
	}

	if ok := validator.Validate(transaction.Transaction{
		Amount:                            1,
		CurrencyCode:                      transaction.USD,
		EntityId:                          "Y",
		CustomerId:                        "12345",
		ThreeDSecureSignatureVerification: "N",
	}); !ok {
		t.Error("expected validation to pass")
	}
}
