package rule

import (
	"bitbucket.verifone.com/validation-service/transaction"
	"testing"
)

func TestNewAmountValidator(t *testing.T) {
	var err error

	// Should create a new amount validator where validation is less than 10
	_, err = newAmountValidator("<", "10")

	if err != nil {
		t.Error("unexpected error while creating new amount validator:", err.Error())
	}

	// Should create a new amount validator where validation is less than or equal to 10
	_, err = newAmountValidator("<=", "10")

	if err != nil {
		t.Error("unexpected error while creating new amount validator:", err.Error())
	}

	// Should create a new amount validator where validation is equal to 10
	_, err = newAmountValidator("==", "10")

	if err != nil {
		t.Error("unexpected error while creating new amount validator:", err.Error())
	}

	// Should create a new amount validator where validation is not equal to 10
	_, err = newAmountValidator("!=", "10")

	if err != nil {
		t.Error("unexpected error while creating new amount validator:", err.Error())
	}

	// Should create a new amount validator where validation is greater than or equal to 10
	_, err = newAmountValidator(">=", "10")

	if err != nil {
		t.Error("unexpected error while creating new amount validator:", err.Error())
	}

	// Should create a new amount validator where validation is greater than 10
	_, err = newAmountValidator(">", "10")

	if err != nil {
		t.Error("unexpected error while creating new amount validator:", err.Error())
	}

	// Should return an error when factory receives an invalid operator
	_, err = newAmountValidator("!", "10")

	if err == nil || err.Error() != "invalid operator" {
		t.Error("expected error while creating new amount validator with invalid operator")
	}

	// Should return an error when factory receives an invalid value
	_, err = newAmountValidator(">", "foo")

	if err == nil || err.Error() != "invalid value format" {
		t.Error("expected error while creating new amount validator with invalid operator")
	}
}

func TestAmountValidator_Validate(t *testing.T) {
	var validator Validator
	var err error

	// Less than
	validator, err = newAmountValidator("<", "2")

	if err != nil {
		t.Error("unexpected error while creating new amount validator:", err.Error())
		return
	}

	if ok := validator.Validate(transaction.Transaction{
		Amount:      1,
		CountryCode: "NL",
		EntityId:    "1",
	}); !ok {
		t.Error("expected validation to pass")
	}

	if ok := validator.Validate(transaction.Transaction{
		Amount:      2,
		CountryCode: "NL",
		EntityId:    "1",
	}); ok {
		t.Error("expected validation to fail")
	}

	if ok := validator.Validate(transaction.Transaction{
		Amount:      3,
		CountryCode: "NL",
		EntityId:    "1",
	}); ok {
		t.Error("expected validation to fail")
	}

	// Less than or equal
	validator, err = newAmountValidator("<=", "2")

	if err != nil {
		t.Error("unexpected error while creating new amount validator:", err.Error())
		return
	}

	if ok := validator.Validate(transaction.Transaction{
		Amount:      1,
		CountryCode: "NL",
		EntityId:    "1",
	}); !ok {
		t.Error("expected validation to pass")
	}

	if ok := validator.Validate(transaction.Transaction{
		Amount:      2,
		CountryCode: "NL",
		EntityId:    "1",
	}); !ok {
		t.Error("expected validation to pass")
	}

	if ok := validator.Validate(transaction.Transaction{
		Amount:      3,
		CountryCode: "NL",
		EntityId:    "1",
	}); ok {
		t.Error("expected validation to fail")
	}

	// Equal
	validator, err = newAmountValidator("==", "2")

	if err != nil {
		t.Error("unexpected error while creating new amount validator:", err.Error())
		return
	}

	if ok := validator.Validate(transaction.Transaction{
		Amount:      1,
		CountryCode: "NL",
		EntityId:    "1",
	}); ok {
		t.Error("expected validation to fail")
	}

	if ok := validator.Validate(transaction.Transaction{
		Amount:      2,
		CountryCode: "NL",
		EntityId:    "1",
	}); !ok {
		t.Error("expected validation to pass")
	}

	if ok := validator.Validate(transaction.Transaction{
		Amount:      3,
		CountryCode: "NL",
		EntityId:    "1",
	}); ok {
		t.Error("expected validation to fail")
	}

	// Greater or equal
	validator, err = newAmountValidator(">=", "2")

	if err != nil {
		t.Error("unexpected error while creating new amount validator:", err.Error())
		return
	}

	if ok := validator.Validate(transaction.Transaction{
		Amount:      1,
		CountryCode: "NL",
		EntityId:    "1",
	}); ok {
		t.Error("expected validation to fail")
	}

	if ok := validator.Validate(transaction.Transaction{
		Amount:      2,
		CountryCode: "NL",
		EntityId:    "1",
	}); !ok {
		t.Error("expected validation to pass")
	}

	if ok := validator.Validate(transaction.Transaction{
		Amount:      3,
		CountryCode: "NL",
		EntityId:    "1",
	}); !ok {
		t.Error("expected validation to pass")
	}

	// Greater than
	validator, err = newAmountValidator(">", "2")

	if err != nil {
		t.Error("unexpected error while creating new amount validator:", err.Error())
		return
	}

	if ok := validator.Validate(transaction.Transaction{
		Amount:      1,
		CountryCode: "NL",
		EntityId:    "1",
	}); ok {
		t.Error("expected validation to fail")
	}

	if ok := validator.Validate(transaction.Transaction{
		Amount:      2,
		CountryCode: "NL",
		EntityId:    "1",
	}); ok {
		t.Error("expected validation to fail")
	}

	if ok := validator.Validate(transaction.Transaction{
		Amount:      3,
		CountryCode: "NL",
		EntityId:    "1",
	}); !ok {
		t.Error("expected validation to pass")
	}
}
