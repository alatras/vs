package rule

import (
	"bitbucket.verifone.com/validation-service/ruleSet/compare"
	"bitbucket.verifone.com/validation-service/transaction"
)

type cardValidator struct {
	comparator compare.StringComparator
}

func newCardValidator(operator Operator, value string) (*cardValidator, error) {
	var comparator compare.StringComparator

	switch operator {
	case equal:
		comparator = compare.EqualString(value)
	case notEqual:
		comparator = compare.NotEqualString(value)
	default:
		return nil, InvalidOperatorError
	}

	return &cardValidator{comparator}, nil
}

func (c cardValidator) Validate(transaction transaction.Transaction) bool {
	return c.comparator(transaction.Card)
}
