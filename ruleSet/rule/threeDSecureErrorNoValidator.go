package rule

import (
	"strconv"
	"validation-service/ruleSet/compare"
	"validation-service/transaction"
)

type threeDSecureErrorNoValidator struct {
	comparator compare.Float64Comparator
}

func newThreeDSecureErrorNoValidator(operator Operator, value string) (*threeDSecureErrorNoValidator, error) {
	if !transaction.IsThreeDSecureErrorNo(value) {
		return nil, InvalidValueError
	}

	valueNumber, err := strconv.Atoi(value)
	if err != nil {
		return nil, InvalidValueError
	}

	var comparator compare.Float64Comparator

	switch operator {
	case OperatorEqual:
		comparator = compare.EqualFloat64(float64(valueNumber))
	case OperatorNotEqual:
		comparator = compare.NotEqualFloat64(float64(valueNumber))
	default:
		return nil, InvalidOperatorError
	}

	return &threeDSecureErrorNoValidator{comparator}, nil
}

func (v threeDSecureErrorNoValidator) Validate(transaction transaction.Transaction) bool {
	errNo, _ := strconv.Atoi(string(transaction.ThreeDSecureErrorNo))
	errNoFloat := float64(errNo)
	return v.comparator(errNoFloat)
}
