package rule

import (
	"bitbucket.verifone.com/validation-service/ruleSet/compare"
	"bitbucket.verifone.com/validation-service/transaction"
)

type customerIPValidator struct {
	comparator compare.StringComparator
}

func newCustomerIPValidator(operator Operator, value string) (*customerIPValidator, error) {
	if !transaction.IsIPv4(value) {
		return nil, InvalidValueError
	}

	var comparator compare.StringComparator

	switch operator {
	case equal:
		comparator = compare.EqualString(value)
	case notEqual:
		comparator = compare.NotEqualString(value)
	default:
		return nil, InvalidOperatorError
	}

	return &customerIPValidator{comparator}, nil
}

func (c customerIPValidator) Validate(transaction transaction.Transaction) bool {
	return c.comparator(transaction.CustomerIP)
}