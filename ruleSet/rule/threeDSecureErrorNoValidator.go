package rule

import (
	"validation-service/ruleSet/compare"
	"validation-service/transaction"
)

type threeDSecureErrorNoValidator struct {
	comparator compare.StringComparator
}

func newThreeDSecureErrorNoValidator(operator Operator, value string) (*threeDSecureErrorNoValidator, error) {
	if !transaction.IsThreeDSecureErrorNo(value) {
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

	return &threeDSecureErrorNoValidator{comparator}, nil
}

func (v threeDSecureErrorNoValidator) Validate(transaction transaction.Transaction) bool {
	return v.comparator(string(transaction.ThreeDSecureErrorNo))
}
