package rule

import (
	"bitbucket.verifone.com/validation-service/ruleSet/compare"
	"bitbucket.verifone.com/validation-service/transaction"
)

type issuerCountryCodeValidator struct {
	comparator compare.StringComparator
}

func newIssuerCountryCodeValidator(operator Operator, value string) (*issuerCountryCodeValidator, error) {
	if !transaction.IsCountryCodeIso31661Alpha3(value) {
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

	return &issuerCountryCodeValidator{comparator}, nil
}

func (i issuerCountryCodeValidator) Validate(transaction transaction.Transaction) bool {
	if transaction.IssuerCountryCode == "" {
		return false
	}

	return i.comparator(string(transaction.IssuerCountryCode))
}
