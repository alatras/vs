package rule

import (
	"testing"

	"validation-service/transaction"
)

func TestNewAmountValidator(t *testing.T) {
	var err error

	// Should create a new amount validator where validation is less than 10
	_, err = newAmountValidator(OperatorLess, "10")

	if err != nil {
		t.Error("unexpected error while creating new amount validator:", err.Error())
	}

	// Should create a new amount validator where validation is less than or equal to 10
	_, err = newAmountValidator(OperatorLessOrEqual, "10")

	if err != nil {
		t.Error("unexpected error while creating new amount validator:", err.Error())
	}

	// Should create a new amount validator where validation is equal to 10
	_, err = newAmountValidator(OperatorEqual, "10")

	if err != nil {
		t.Error("unexpected error while creating new amount validator:", err.Error())
	}

	// Should create a new amount validator where validation is not equal to 10
	_, err = newAmountValidator(OperatorNotEqual, "10")

	if err != nil {
		t.Error("unexpected error while creating new amount validator:", err.Error())
	}

	// Should create a new amount validator where validation is greater than or equal to 10
	_, err = newAmountValidator(OperatorGreaterOrEqual, "10")

	if err != nil {
		t.Error("unexpected error while creating new amount validator:", err.Error())
	}

	// Should create a new amount validator where validation is greater than 10
	_, err = newAmountValidator(OperatorGreater, "10")

	if err != nil {
		t.Error("unexpected error while creating new amount validator:", err.Error())
	}

	// Should return an error when factory receives an invalid operator
	_, err = newAmountValidator("foo", "10")

	if err != InvalidOperatorError {
		t.Error("expected error while creating new amount validator with invalid operator")
	}

	// Should return an error when factory receives an invalid value
	_, err = newAmountValidator(OperatorGreater, "foo")

	if err != InvalidValueError {
		t.Error("expected error while creating new amount validator with invalid operator")
	}
}

func TestAmountValidator_Validate(t *testing.T) {
	var validator Validator
	var err error

	// Less than
	validator, err = newAmountValidator(OperatorLess, "2")

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
	validator, err = newAmountValidator(OperatorLessOrEqual, "2")

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
	validator, err = newAmountValidator(OperatorEqual, "2")

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
	validator, err = newAmountValidator(OperatorGreaterOrEqual, "2")

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
	validator, err = newAmountValidator(OperatorGreater, "2")

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
	validator, err = newAmountValidator(OperatorEqual, "20")

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
