package rule

import (
	"validation-service/ruleSet/compare"
	"validation-service/transaction"
)

type customerIdValidator struct {
	comparator compare.StringComparator
}

func newCustomerIdValidator(operator Operator, value string) (*customerIdValidator, error) {
	var comparator compare.StringComparator

	switch operator {
	case OperatorEqual:
		comparator = compare.EqualString(value)
	case OperatorNotEqual:
		comparator = compare.NotEqualString(value)
	default:
		return nil, InvalidOperatorError
	}

	return &customerIdValidator{comparator}, nil
}

func (c customerIdValidator) Validate(transaction transaction.Transaction) bool {
	return c.comparator(transaction.CustomerId)
}
