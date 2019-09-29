package rule

import (
	"bitbucket.verifone.com/validation-service/ruleSet/compare"
	"bitbucket.verifone.com/validation-service/transaction"
)

type customerCountryCodeValidator struct {
	comparator compare.StringComparator
}

func newCustomerCountryCodeValidator(operator operator, value string) (*customerCountryCodeValidator, error) {
	if !transaction.IsCountryCodeIso31661Alpha2(value) {
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

	return &customerCountryCodeValidator{comparator}, nil
}

func (v customerCountryCodeValidator) Validate(transaction transaction.Transaction) bool {
	return v.comparator(string(transaction.CustomerCountryCode))
}
