package rule

import (
	"validation-service/ruleSet/compare"
	"validation-service/transaction"
)

type threeDSecureSignatureVerificationValidator struct {
	comparator compare.StringComparator
}

func newThreeDSecureSignatureVerificationValidator(operator Operator, value string) (*threeDSecureSignatureVerificationValidator, error) {
	if !transaction.IsThreeDSecureSignatureVerification(value) {
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

	return &threeDSecureSignatureVerificationValidator{comparator}, nil
}

func (v threeDSecureSignatureVerificationValidator) Validate(transaction transaction.Transaction) bool {
	return v.comparator(string(transaction.ThreeDSecureSignatureVerification))
}
