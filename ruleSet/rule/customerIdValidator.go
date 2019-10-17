package rule

import (
	"bitbucket.verifone.com/validation-service/ruleSet/compare"
	"bitbucket.verifone.com/validation-service/transaction"
)

type customerIdValidator struct {
	comparator compare.StringComparator
}

func newCustomerIdValidator(operator Operator, value string) (*customerIdValidator, error) {
	var comparator compare.StringComparator

	switch operator {
	case equal:
		comparator = compare.EqualString(value)
	case notEqual:
		comparator = compare.NotEqualString(value)
	default:
		return nil, InvalidOperatorError
	}

	return &customerIdValidator{comparator}, nil
}

func (c customerIdValidator) Validate(transaction transaction.Transaction) bool {
	return c.comparator(transaction.CustomerId)
}
