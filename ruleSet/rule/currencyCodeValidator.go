package rule

import (
	"bitbucket.verifone.com/validation-service/ruleSet/compare"
	"bitbucket.verifone.com/validation-service/transaction"
)

type currencyCodeValidator struct {
	comparator compare.StringComparator
}

func newCurrencyCodeValidator(operator operator, value string) (*currencyCodeValidator, error) {
	var comparator compare.StringComparator

	switch operator {
	case equal:
		comparator = compare.EqualString(value)
	case notEqual:
		comparator = compare.NotEqualString(value)
	default:
		return nil, InvalidOperatorError
	}

	return &currencyCodeValidator{comparator}, nil
}

func (v currencyCodeValidator) Validate(transaction transaction.Transaction) bool {
	return v.comparator(string(transaction.CurrencyCode))
}
