package rule

import (
	"bitbucket.verifone.com/validation-service/transaction"
	"testing"
)

func TestNewAmountValidator(t *testing.T) {
	var err error

	// Should create a new amount validator where validation is less than 10
	_, err = newAmountValidator(less, "10")

	if err != nil {
		t.Error("unexpected error while creating new amount validator:", err.Error())
	}

	// Should create a new amount validator where validation is less than or equal to 10
	_, err = newAmountValidator(lessOrEqual, "10")

	if err != nil {
		t.Error("unexpected error while creating new amount validator:", err.Error())
	}

	// Should create a new amount validator where validation is equal to 10
	_, err = newAmountValidator(equal, "10")

	if err != nil {
		t.Error("unexpected error while creating new amount validator:", err.Error())
	}

	// Should create a new amount validator where validation is not equal to 10
	_, err = newAmountValidator(notEqual, "10")

	if err != nil {
		t.Error("unexpected error while creating new amount validator:", err.Error())
	}

	// Should create a new amount validator where validation is greater than or equal to 10
	_, err = newAmountValidator(greaterOrEqual, "10")

	if err != nil {
		t.Error("unexpected error while creating new amount validator:", err.Error())
	}

	// Should create a new amount validator where validation is greater than 10
	_, err = newAmountValidator(greater, "10")

	if err != nil {
		t.Error("unexpected error while creating new amount validator:", err.Error())
	}

	// Should return an error when factory receives an invalid operator
	_, err = newAmountValidator("foo", "10")

	if err != InvalidOperatorError {
		t.Error("expected error while creating new amount validator with invalid operator")
	}

	// Should return an error when factory receives an invalid value
	_, err = newAmountValidator(greater, "foo")

	if err != InvalidValueError {
		t.Error("expected error while creating new amount validator with invalid operator")
	}
}

func TestAmountValidator_Validate(t *testing.T) {
	var validator Validator
	var err error

	// Less than
	validator, err = newAmountValidator(less, "2")

	if err != nil {
		t.Error("unexpected error while creating new amount validator:", err.Error())
		return
	}

	if ok := validator.Validate(transaction.Transaction{
		Amount: 1,
	}); !ok {
		t.Error("expected validation to pass")
	}

	if ok := validator.Validate(transaction.Transaction{
		Amount: 2,
	}); ok {
		t.Error("expected validation to fail")
	}

	if ok := validator.Validate(transaction.Transaction{
		Amount: 3,
	}); ok {
		t.Error("expected validation to fail")
	}

	// Less than or equal
	validator, err = newAmountValidator(lessOrEqual, "2")

	if err != nil {
		t.Error("unexpected error while creating new amount validator:", err.Error())
		return
	}

	if ok := validator.Validate(transaction.Transaction{
		Amount: 1,
	}); !ok {
		t.Error("expected validation to pass")
	}

	if ok := validator.Validate(transaction.Transaction{
		Amount: 2,
	}); !ok {
		t.Error("expected validation to pass")
	}

	if ok := validator.Validate(transaction.Transaction{
		Amount: 3,
	}); ok {
		t.Error("expected validation to fail")
	}

	// Equal
	validator, err = newAmountValidator(equal, "2")

	if err != nil {
		t.Error("unexpected error while creating new amount validator:", err.Error())
		return
	}

	if ok := validator.Validate(transaction.Transaction{
		Amount: 1,
	}); ok {
		t.Error("expected validation to fail")
	}

	if ok := validator.Validate(transaction.Transaction{
		Amount: 2,
	}); !ok {
		t.Error("expected validation to pass")
	}

	if ok := validator.Validate(transaction.Transaction{
		Amount: 3,
	}); ok {
		t.Error("expected validation to fail")
	}

	// Greater or equal
	validator, err = newAmountValidator(greaterOrEqual, "2")

	if err != nil {
		t.Error("unexpected error while creating new amount validator:", err.Error())
		return
	}

	if ok := validator.Validate(transaction.Transaction{
		Amount: 1,
	}); ok {
		t.Error("expected validation to fail")
	}

	if ok := validator.Validate(transaction.Transaction{
		Amount: 2,
	}); !ok {
		t.Error("expected validation to pass")
	}

	if ok := validator.Validate(transaction.Transaction{
		Amount: 3,
	}); !ok {
		t.Error("expected validation to pass")
	}

	// Greater than
	validator, err = newAmountValidator(greater, "2")

	if err != nil {
		t.Error("unexpected error while creating new amount validator:", err.Error())
		return
	}

	if ok := validator.Validate(transaction.Transaction{
		Amount: 1,
	}); ok {
		t.Error("expected validation to fail")
	}

	if ok := validator.Validate(transaction.Transaction{
		Amount: 2,
	}); ok {
		t.Error("expected validation to fail")
	}

	if ok := validator.Validate(transaction.Transaction{
		Amount: 3,
	}); !ok {
		t.Error("expected validation to pass")
	}

	// Minor units
	validator, err = newAmountValidator(equal, "20")

	if err != nil {
		t.Error("unexpected error while creating new amount validator:", err.Error())
		return
	}

	if ok := validator.Validate(transaction.Transaction{
		Amount: 20,
	}); !ok {
		t.Error("expected validation to pass")
	}

	if ok := validator.Validate(transaction.Transaction{
		Amount:     200,
		MinorUnits: 1,
	}); !ok {
		t.Error("expected validation to pass")
	}

	if ok := validator.Validate(transaction.Transaction{
		Amount:     2000,
		MinorUnits: 2,
	}); !ok {
		t.Error("expected validation to pass")
	}

	if ok := validator.Validate(transaction.Transaction{
		Amount:     20000,
		MinorUnits: 3,
	}); !ok {
		t.Error("expected validation to pass")
	}
}
