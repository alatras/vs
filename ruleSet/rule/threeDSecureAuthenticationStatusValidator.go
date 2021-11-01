package rule

import (
	"validation-service/ruleSet/compare"
	"validation-service/transaction"
)

type threeDSecureAuthenticationStatusValidator struct {
	comparator compare.StringComparator
}

func newThreeDSecureAuthenticationStatusValidator(operator Operator, value string) (*threeDSecureAuthenticationStatusValidator, error) {
	if !transaction.IsThreeDSecureAuthenticationStatus(value) {
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

	return &threeDSecureAuthenticationStatusValidator{comparator}, nil
}

func (v threeDSecureAuthenticationStatusValidator) Validate(transaction transaction.Transaction) bool {
	return v.comparator(string(transaction.ThreeDSecureAuthenticationStatus))
}
