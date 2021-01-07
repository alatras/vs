package rule

import (
	"validation-service/ruleSet/compare"
	"validation-service/transaction"
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
	case OperatorEqual:
		comparator = compare.EqualString(value)
	case OperatorNotEqual:
		comparator = compare.NotEqualString(value)
	default:
		return nil, InvalidOperatorError
	}

	return &customerIPValidator{comparator}, nil
}

func (c customerIPValidator) Validate(transaction transaction.Transaction) bool {
	return c.comparator(transaction.CustomerIP)
}
