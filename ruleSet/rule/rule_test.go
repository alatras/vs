package rule

import (
	"testing"
)

func TestNewValidator(t *testing.T) {
	var err error

	// Should successfully create a trx amount is equal to 10 validator
	_, err = NewValidator(Metadata{
		Property: "amount",
		Operator: equal,
		Value:    "10",
	})

	if err != nil {
		t.Error("unexpected error while creating new amount validator:", err.Error())
	}

	// Should not successfully create a trx undefined field is equal to 10 validator
	_, err = NewValidator(Metadata{
		Property: "foo",
		Operator: equal,
		Value:    "10",
	})

	if err == nil || err.Error() != "invalid validator property" {
		t.Error("expected error while creating new undefined field validator")
	}
}
