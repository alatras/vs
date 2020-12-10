package rule

import (
	"bitbucket.verifone.com/validation-service/transaction"
)

type dummyCardValidator struct {
}

func newCardValidator(operator Operator, value string) (*dummyCardValidator, error) {
	if operator != equal && operator != notEqual {
		return nil, InvalidOperatorError
	}

	return &dummyCardValidator{}, nil
}

func (c dummyCardValidator) Validate(transaction transaction.Transaction) bool {
	return false
}
