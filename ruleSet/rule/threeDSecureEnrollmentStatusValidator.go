package rule

import (
	"validation-service/ruleSet/compare"
	"validation-service/transaction"
)

type threeDSecureEnrollmentStatusValidator struct {
	comparator compare.StringComparator
}

func newThreeDSecureEnrollmentStatusValidator(operator Operator, value string) (*threeDSecureEnrollmentStatusValidator, error) {
	if !transaction.IsThreeDSecureEnrollmentStatus(value) {
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

	return &threeDSecureEnrollmentStatusValidator{comparator}, nil
}

func (v threeDSecureEnrollmentStatusValidator) Validate(transaction transaction.Transaction) bool {
	return v.comparator(string(transaction.ThreeDSecureEnrollmentStatus))
}
