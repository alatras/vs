package rule

import (
	"bitbucket.verifone.com/validation-service/ruleSet/compare"
	"bitbucket.verifone.com/validation-service/transaction"
	"errors"
)

type countryCodeValidator struct {
	countryCodeComparator compare.StringComparator
}

func newCountryCodeValidator(operator Operator, value string) (*countryCodeValidator, error) {
	var countryCodeComparator compare.StringComparator

	switch operator {
	case equal:
		countryCodeComparator = compare.EqualString(value)
	case notEqual:
		countryCodeComparator = compare.NotEqualString(value)
	default:
		return nil, errors.New("invalid operator")
	}

	return &countryCodeValidator{countryCodeComparator}, nil
}

func (v countryCodeValidator) Validate(transaction transaction.Transaction) bool {
	return v.countryCodeComparator(transaction.CountryCode)
}
