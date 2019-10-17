package rule

import (
	"bitbucket.verifone.com/validation-service/ruleSet/compare"
	"bitbucket.verifone.com/validation-service/transaction"
)

type customerIPCountryCodeValidator struct {
	comparator compare.StringComparator
}

func newCustomerIPCountryCodeValidator(operator Operator, value string) (*customerIPCountryCodeValidator, error) {
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

	return &customerIPCountryCodeValidator{comparator}, nil
}

func (c customerIPCountryCodeValidator) Validate(transaction transaction.Transaction) bool {
	return c.comparator(transaction.CustomerIPCountry)
}
