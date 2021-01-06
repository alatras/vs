package rule

import (
	"validation-service/transaction"
)

type dummyCardValidator struct {
}

func newCardValidator(operator Operator, value string) (*dummyCardValidator, error) {
	if operator != OperatorEqual && operator != OperatorNotEqual {
		return nil, InvalidOperatorError
	}

	return &dummyCardValidator{}, nil
}

func (c dummyCardValidator) Validate(transaction transaction.Transaction) bool {
	return false
}
