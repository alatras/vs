package rule

import (
	"fmt"
	"strconv"
	"validation-service/ruleSet/compare"
	"validation-service/transaction"
)

type threeDSecureErrorNoValidator struct {
	comparator compare.StringComparator
}

func newThreeDSecureErrorNoValidator(operator Operator, value string) (*threeDSecureErrorNoValidator, error) {
	valueNumber, err := strconv.Atoi(value)
	if err != nil {
		return nil, InvalidValueError
	}

	var comparator compare.StringComparator

	switch operator {
	case OperatorEqual:
		comparator = compare.EqualString(fmt.Sprint(valueNumber))
	case OperatorNotEqual:
		comparator = compare.NotEqualString(fmt.Sprint(valueNumber))
	default:
		return nil, InvalidOperatorError
	}

	return &threeDSecureErrorNoValidator{comparator}, nil
}

func (v threeDSecureErrorNoValidator) Validate(transaction transaction.Transaction) bool {
	return v.comparator(transaction.ThreeDSecureErrorNo)
}
